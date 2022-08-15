package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

// Handles a transfer from one address to another. If it can be a forceful transfer, it will forcefully transfer the balances and approvals. If it is a pending transfer, it will add it to the pending transfers.
func HandleTransfer(ctx sdk.Context, badge types.BitBadge, subbadgeRange *types.IdRange, fromUserBalanceInfo types.UserBalanceInfo, toUserBalanceInfo types.UserBalanceInfo, amount uint64, from uint64, to uint64, approvedBy uint64, expirationTime uint64, cantCancelBeforeTime uint64) (types.UserBalanceInfo, types.UserBalanceInfo, error) {
	permissions := types.GetPermissions(badge.Permissions)
	err := *new(error)
	sendingToReservedAddress := false //TODO: implement this; Check if to Address is reserved; if so, we automatically forceful transfer

	canForcefulTransfer := sendingToReservedAddress || permissions.ForcefulTransfers || badge.Manager == to
	if canForcefulTransfer {
		fromUserBalanceInfo, toUserBalanceInfo, err = ForcefulTransfer(ctx, badge, subbadgeRange, fromUserBalanceInfo, toUserBalanceInfo, amount, from, to, approvedBy, expirationTime)
	} else {
		fromUserBalanceInfo, toUserBalanceInfo, err = PendingTransfer(ctx, badge, subbadgeRange, fromUserBalanceInfo, toUserBalanceInfo, amount, from, to, approvedBy, expirationTime, cantCancelBeforeTime)
	}

	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	return fromUserBalanceInfo, toUserBalanceInfo, nil
}

//Forceful transfers will transfer the balances and deduct from approvals directly without adding it to pending.
func ForcefulTransfer(ctx sdk.Context, badge types.BitBadge, subbadgeRange *types.IdRange, fromUserBalanceInfo types.UserBalanceInfo, toUserBalanceInfo types.UserBalanceInfo, amount uint64, from uint64, to uint64, approvedBy uint64, expirationTime uint64) (types.UserBalanceInfo, types.UserBalanceInfo, error) {
	// 1. Check if the from address is frozen
	// 2. Remove approvals if approvedBy != from
	// 3. Deduct from "From" balance
	// 4. Add to "To" balance
	err := AssertAccountNotFrozen(ctx, badge, from)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	fromUserBalanceInfo, err = DeductApprovals(ctx, fromUserBalanceInfo, badge, badge.Id, subbadgeRange, from, to, approvedBy, amount)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}
	
	fromUserBalanceInfo, err = SubtractBalancesForIdRanges(ctx, fromUserBalanceInfo, []*types.IdRange{subbadgeRange}, amount)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	toUserBalanceInfo, err = AddBalancesForIdRanges(ctx, toUserBalanceInfo, []*types.IdRange{subbadgeRange}, amount)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	return fromUserBalanceInfo, toUserBalanceInfo, nil
}

// Removes balances and approvals from "From" address and puts them in escrow. Adds the pending transfer object to both parties' pending array.
func PendingTransfer(ctx sdk.Context, badge types.BitBadge, subbadgeRange *types.IdRange, fromUserBalanceInfo types.UserBalanceInfo, toUserBalanceInfo types.UserBalanceInfo, amount uint64, from uint64, to uint64, approvedBy uint64, expirationTime uint64, cantCancelBeforeTime uint64) (types.UserBalanceInfo, types.UserBalanceInfo, error) {
	// 1. Check if the from address is frozen
	// 2. Remove approvals if approvedBy != from
	// 3. Deduct from "From" balance
	// 4. Append pending transfers to both parties
	// 5. If the pending tranfer is eventually accepted, we simply add the balance to "To". If it is removed, we revert the balance and approvals.
	err := AssertAccountNotFrozen(ctx, badge, from)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	fromUserBalanceInfo, err = DeductApprovals(ctx, fromUserBalanceInfo, badge, badge.Id, subbadgeRange, from, to, approvedBy, amount)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	fromUserBalanceInfo, err = SubtractBalancesForIdRanges(ctx, fromUserBalanceInfo, []*types.IdRange{subbadgeRange}, amount)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	fromUserBalanceInfo, toUserBalanceInfo, err = AppendPendingTransferForBothParties(ctx, fromUserBalanceInfo, toUserBalanceInfo, subbadgeRange, to, from, amount, approvedBy, true, expirationTime, cantCancelBeforeTime)
	if err != nil {
		return types.UserBalanceInfo{}, types.UserBalanceInfo{}, err
	}

	return fromUserBalanceInfo, toUserBalanceInfo, nil
}

// Deduct approvals from requester if requester != from
func DeductApprovals(ctx sdk.Context, userBalanceInfo types.UserBalanceInfo, badge types.BitBadge, badgeId uint64, rangeToDeduct *types.IdRange, from uint64, to uint64, requester uint64, amount uint64) (types.UserBalanceInfo, error) {
	newUserBalanceInfo := userBalanceInfo

	if from != requester {
		err := *new(error)
		newUserBalanceInfo, err = RemoveBalanceFromApproval(ctx, newUserBalanceInfo, amount, requester, rangeToDeduct)
		if err != nil {
			return userBalanceInfo, err
		}
	}

	return newUserBalanceInfo, nil
}

// Revert escrowed balances and approvals if a pending transfer is rejected / cancelled. 
func RevertEscrowedBalancesAndApprovals(ctx sdk.Context, userBalanceInfo types.UserBalanceInfo, rangeToRevert *types.IdRange, from uint64, approvedBy uint64, amount uint64) (types.UserBalanceInfo, error) {
	err := *new(error)
	userBalanceInfo, err = AddBalancesForIdRanges(ctx, userBalanceInfo, []*types.IdRange{rangeToRevert}, amount)
	if err != nil {
		return types.UserBalanceInfo{}, err
	}

	//If it was sent via an approval, we need to add the approval back
	if approvedBy != from {
		userBalanceInfo, err = AddBalanceToApproval(ctx, userBalanceInfo, amount, approvedBy, rangeToRevert)
		if err != nil {
			return types.UserBalanceInfo{}, err
		}
	}

	return userBalanceInfo, nil
}

// Checks if account is frozen or not.
func IsAccountFrozen(badge types.BitBadge, permissions types.Permissions, address uint64) bool {
	frozenByDefault := permissions.FrozenByDefault

	accountIsFrozen := false
	if frozenByDefault {
		_, found := SearchIdRangesForId(address, badge.FreezeRanges)
		if !found {
			accountIsFrozen = true
		}
	} else {
		_, found := SearchIdRangesForId(address, badge.FreezeRanges)
		if found {
			accountIsFrozen = true
		}
	}
	return accountIsFrozen
}

// Returns an error if account is Frozen
func AssertAccountNotFrozen(ctx sdk.Context, badge types.BitBadge, from uint64) error {
	permissions := types.GetPermissions(badge.Permissions)

	accountIsFrozen := IsAccountFrozen(badge, permissions, from)
	if accountIsFrozen {
		return ErrAddressFrozen
	}

	return nil
}
