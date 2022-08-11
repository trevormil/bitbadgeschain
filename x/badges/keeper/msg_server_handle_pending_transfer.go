package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

/*
	Initial Pending Transfer From A -> Balances / approvals put into escrow
		B Accepts -> Simple add escrowed balance to B, remove from both pending
		B Rejects -> Remove from B Pending
			A has to trigger the escrow revert and remove from their balance
		A cancels -> Revert escrows and remove from both pending

	Initial Transfer Request From A
		B Accepts It (Forceful) -> Forceful transfer B to A -> Remove from both pending
		B Accepts It (Approve) -> Set pending transfer as approved in B's pending list
			A can then trigger a forceful transfer B to A -> Remove from both pending
		B Rejects -> Remove from B Pending
			A has to remove from their pending
		A cancels -> Remove from both pending
*/
func (k msgServer) HandlePendingTransfer(goCtx context.Context, msg *types.MsgHandlePendingTransfer) (*types.MsgHandlePendingTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	CreatorAccountNum, badge, err := k.UniversalValidate(ctx, UniversalValidationParams{
		Creator: msg.Creator,
		BadgeId: msg.BadgeId,
	})
	if err != nil {
		return nil, err
	}

	//Get creator balance information
	creatorBalanceKey := ConstructBalanceKey(CreatorAccountNum, msg.BadgeId)
	creatorBalanceInfo, found := k.GetUserBalanceFromStore(ctx, creatorBalanceKey)
	if !found {
		return nil, ErrUserBalanceNotExists
	}

	//We will store user balances in a cache and a didUpdate accountNum -> bool map to keep track of which ones are updated to avoid unnecessary writes.
	balanceInfoCache := make(map[uint64]types.UserBalanceInfo, 0)
	didUpdateBalanceInfo := make(map[uint64]bool, 0)

	updated := false
	for _, nonceRange := range msg.NonceRanges {
		//TODO: In the future, we can make this a binary search since this is all sorted by the nonces (append-only)
		for idx, CurrPendingTransfer := range creatorBalanceInfo.Pending {
			if CurrPendingTransfer.ThisPendingNonce <= nonceRange.End && CurrPendingTransfer.ThisPendingNonce >= nonceRange.Start {
				updated = true

				// Handle accepting after expiration time
				if msg.Accept && CurrPendingTransfer.ExpirationTime != 0 {
					expired := uint64(ctx.BlockTime().Unix()) > CurrPendingTransfer.ExpirationTime
					if expired {
						return nil, ErrPendingTransferExpired
					}
				}
				
				outgoingTransfer := CurrPendingTransfer.From == CreatorAccountNum
				
				//Normal transfer flags
				acceptOwnOutgoingTransfer := CurrPendingTransfer.Sent && msg.Accept && outgoingTransfer
				acceptIncomingTransfer := !CurrPendingTransfer.Sent && msg.Accept && !outgoingTransfer
				cancelOwnOutgoingTransfer := CurrPendingTransfer.Sent && !msg.Accept && outgoingTransfer
				rejectIncomingTransfer := !CurrPendingTransfer.Sent && !msg.Accept

				if acceptOwnOutgoingTransfer {
					return nil, ErrCantAcceptOwnTransferRequest
				}

				//Pending transfer flags
				finalizeOwnTransferRequestAfterApprovedByOtherParty := CurrPendingTransfer.Sent && msg.Accept && !outgoingTransfer
				acceptTransferRequestButMarkAsApproved := !msg.ForcefulAccept && !CurrPendingTransfer.Sent && msg.Accept && outgoingTransfer
				acceptTransferRequestForcefully := msg.ForcefulAccept && !CurrPendingTransfer.Sent && msg.Accept && outgoingTransfer
				rejectTransferRequest := !CurrPendingTransfer.Sent && !msg.Accept
				cancelOwnTransferRequest := CurrPendingTransfer.Sent && !msg.Accept && !outgoingTransfer

				//Other helper flags
				onlyUpdatingCreatorBalance := rejectIncomingTransfer || rejectTransferRequest || acceptTransferRequestButMarkAsApproved
				otherPartyMayHaveRemovedAlready := cancelOwnOutgoingTransfer || cancelOwnTransferRequest
				needToRemoveAtLeastOneFromPending := !acceptTransferRequestButMarkAsApproved

				//Get basic information about the other party and update cache if not already in cache
				otherPartyAccountNum := CurrPendingTransfer.From
				if outgoingTransfer {
					otherPartyAccountNum = CurrPendingTransfer.To
				}
				otherPartyBalanceKey := ConstructBalanceKey(otherPartyAccountNum, msg.BadgeId)
				otherPartyNonce := CurrPendingTransfer.OtherPendingNonce
				otherPartyBalanceInfo, ok := balanceInfoCache[otherPartyAccountNum]
				if !ok && !onlyUpdatingCreatorBalance {
					otherPartyBalanceInfo, found = k.GetUserBalanceFromStore(ctx, otherPartyBalanceKey)
					balanceInfoCache[otherPartyAccountNum] = otherPartyBalanceInfo
					if !found {
						return nil, ErrUserBalanceNotExists
					}
				}

				for i := CurrPendingTransfer.SubbadgeRange.Start; i <= CurrPendingTransfer.SubbadgeRange.End; i++ {
					if acceptIncomingTransfer {
						creatorBalanceInfo, err = AddBalanceForId(ctx, creatorBalanceInfo, i, CurrPendingTransfer.Amount)
					} else if acceptTransferRequestButMarkAsApproved {
						creatorBalanceInfo.Pending[idx].MarkedAsAccepted = true
					} else if cancelOwnOutgoingTransfer {
						creatorBalanceInfo, err = k.RevertEscrowedBalancesAndApprovals(ctx, creatorBalanceInfo, i,CurrPendingTransfer.From, CurrPendingTransfer.ApprovedBy, CurrPendingTransfer.Amount)
					} else if finalizeOwnTransferRequestAfterApprovedByOtherParty {
						//TODO: clean this up
						approved := false
						for _, pending_info := range otherPartyBalanceInfo.Pending {
							if pending_info.ThisPendingNonce == otherPartyNonce && pending_info.OtherPendingNonce == CurrPendingTransfer.ThisPendingNonce && pending_info.MarkedAsAccepted {
								approved = true
							}
						}

						if !approved {
							return nil, ErrNotApproved
						}
						otherPartyBalanceInfo, creatorBalanceInfo, err = k.ForcefulTransfer(ctx, badge, types.IdRange{Start: i, End: i}, otherPartyBalanceInfo, creatorBalanceInfo, CurrPendingTransfer.Amount, CurrPendingTransfer.From, CurrPendingTransfer.To, CurrPendingTransfer.From, CurrPendingTransfer.ExpirationTime)
					} else if acceptTransferRequestForcefully {
						creatorBalanceInfo, otherPartyBalanceInfo, err = k.ForcefulTransfer(ctx, badge, types.IdRange{Start: i, End: i}, creatorBalanceInfo, otherPartyBalanceInfo, CurrPendingTransfer.Amount, CurrPendingTransfer.From, CurrPendingTransfer.To, CreatorAccountNum, CurrPendingTransfer.ExpirationTime)
					}

					if err != nil {
						return nil, err
					}
				}
				
				if needToRemoveAtLeastOneFromPending {
					creatorBalanceInfo, err = k.RemovePending(ctx, creatorBalanceInfo, CurrPendingTransfer.ThisPendingNonce, otherPartyNonce)
					if err != nil {
						return nil, err
					}

					//Try to remove from the other party's pending, but in some situations it may have already been removed
					if !onlyUpdatingCreatorBalance {
						otherPartyBalanceInfo, err = k.RemovePending(ctx, otherPartyBalanceInfo, otherPartyNonce, CurrPendingTransfer.ThisPendingNonce)
						didUpdateBalanceInfo[otherPartyAccountNum] = true
						balanceInfoCache[otherPartyAccountNum] = otherPartyBalanceInfo
						if err != nil {
							if !(err == ErrPendingNotFound && otherPartyMayHaveRemovedAlready) {
								return nil, err
							}
						}
					}
				}				
			}
		}
	}

	if !updated {
		return nil, ErrNoPendingTransferFound
	}

	err = k.SetUserBalanceInStore(ctx, creatorBalanceKey, creatorBalanceInfo)
	if err != nil {
		return nil, err
	}

	//For all user balances that we did update, update the store
	for key := range didUpdateBalanceInfo {
		err = k.SetUserBalanceInStore(ctx, ConstructBalanceKey(key, msg.BadgeId), balanceInfoCache[key])
		if err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "badges"),
			sdk.NewAttribute(sdk.AttributeKeyAction, "HandledPendingTransfers"),
			sdk.NewAttribute("Creator", fmt.Sprint(CreatorAccountNum)),
			sdk.NewAttribute("Accepted", fmt.Sprint(msg.Accept)),
			sdk.NewAttribute("BadgeId", fmt.Sprint(msg.BadgeId)),
			sdk.NewAttribute("NonceRanges", fmt.Sprint(msg.NonceRanges)),
		),
	)

	return &types.MsgHandlePendingTransferResponse{}, nil

}
