package keeper

import (
	"context"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateUserApprovedTransfers(goCtx context.Context, msg *types.MsgUpdateUserApprovedTransfers) (*types.MsgUpdateUserApprovedTransfersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection, err := k.UniversalValidate(ctx, UniversalValidationParams{
		Creator:                              msg.Creator,
		CollectionId:                         msg.CollectionId,
	})
	if err != nil {
		return nil, err
	}

	if !IsStandardBalances(collection) {
		return nil, ErrWrongBalancesType
	}

	for _, addressMapping := range msg.AddressMappings {
		if err := k.CreateAddressMapping(ctx, addressMapping); err != nil {
			return nil, err
		}
	}

	balanceKey := ConstructBalanceKey(msg.Creator, collection.CollectionId)
	userBalance, found := k.GetUserBalanceFromStore(ctx, balanceKey)
	if !found {
		userBalance = &types.UserBalanceStore{
			Balances : []*types.Balance{},
			ApprovedOutgoingTransfersTimeline: collection.DefaultUserApprovedOutgoingTransfersTimeline,
			ApprovedIncomingTransfersTimeline: collection.DefaultUserApprovedIncomingTransfersTimeline,
			Permissions: &types.UserPermissions{
				CanUpdateApprovedOutgoingTransfers: []*types.UserApprovedOutgoingTransferPermission{},
				CanUpdateApprovedIncomingTransfers: []*types.UserApprovedIncomingTransferPermission{},
			},
		}
	}

	manager := types.GetCurrentManager(ctx, collection)
	
	if err := k.ValidateUserApprovedOutgoingTransfersUpdate(ctx, userBalance.ApprovedOutgoingTransfersTimeline, msg.ApprovedOutgoingTransfersTimeline, userBalance.Permissions.CanUpdateApprovedOutgoingTransfers, manager); err != nil {
		return nil, err
	}
	userBalance.ApprovedOutgoingTransfersTimeline = msg.ApprovedOutgoingTransfersTimeline
	

	if err := k.ValidateUserApprovedIncomingTransfersUpdate(ctx, userBalance.ApprovedIncomingTransfersTimeline, msg.ApprovedIncomingTransfersTimeline, userBalance.Permissions.CanUpdateApprovedIncomingTransfers, manager); err != nil {
		return nil, err
	}
	userBalance.ApprovedIncomingTransfersTimeline = msg.ApprovedIncomingTransfersTimeline


	err = k.SetUserBalanceInStore(ctx, balanceKey, userBalance)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "badges"),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	)

	return &types.MsgUpdateUserApprovedTransfersResponse{}, nil
}
