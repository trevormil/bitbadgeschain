package keeper

import (
	"context"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NewCollection(goCtx context.Context, msg *types.MsgNewCollection) (*types.MsgNewCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	CreatorAccountNum := k.Keeper.MustGetAccountNumberForBech32AddressString(ctx, msg.Creator)

	accsToCheck := []uint64{}
	for _, transfer := range msg.Transfers {
		accsToCheck = append(accsToCheck, transfer.ToAddresses...)
	}

	rangesToValidate := []*types.IdRange{}
	for _, transfer := range msg.Transfers {
		for _, balance := range transfer.Balances {
			rangesToValidate = append(rangesToValidate, balance.BadgeIds...)
		}
	}

	NextCollectionId := k.GetNextCollectionId(ctx)
	k.IncrementNextCollectionId(ctx)

	collection := types.BadgeCollection{
		CollectionId:             NextCollectionId,
		CollectionUri:            msg.CollectionUri,
		BadgeUris:                msg.BadgeUris,
		Manager:                  CreatorAccountNum,
		Permissions:              msg.Permissions,
		DisallowedTransfers:      msg.DisallowedTransfers,
		ManagerApprovedTransfers: msg.ManagerApprovedTransfers,
		Bytes:                    msg.Bytes,
		Standard:                 msg.Standard,
		NextBadgeId: 			  1,
	}

	if len(msg.BadgeSupplys) != 0 {
		err := *new(error)
		collection, err = k.CreateBadges(ctx, collection, msg.BadgeSupplys, msg.Transfers, msg.Claims, msg.Creator)
		if err != nil {
			return nil, err
		}
	}

	if err := k.SetCollectionInStore(ctx, collection); err != nil {
		return nil, err
	}
	

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "badges"),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	)

	return &types.MsgNewCollectionResponse{
		CollectionId: NextCollectionId,
	}, nil
}
