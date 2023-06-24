package keeper

import (
	"context"
	"fmt"
	"math"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) NewCollection(goCtx context.Context, msg *types.MsgNewCollection) (*types.MsgNewCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	NextCollectionId := k.GetNextCollectionId(ctx)
	k.IncrementNextCollectionId(ctx)

	collection := types.BadgeCollection{
		CollectionId:       NextCollectionId,
		CollectionMetadataTimeline: msg.CollectionMetadataTimeline,
		OffChainBalancesMetadataTimeline:   msg.OffChainBalancesMetadataTimeline,
		BadgeMetadataTimeline:      msg.BadgeMetadataTimeline,
		ManagerTimeline:            []*types.ManagerTimeline{
			{
				Manager: msg.Creator,
				Times: []*types.IdRange{
					{
						Start: sdk.NewUint(0),
						End:   sdk.NewUint(math.MaxUint64),
					},
				},
			},
		},
		Permissions:        msg.Permissions,
		ApprovedTransfersTimeline:  msg.ApprovedTransfersTimeline,
		CustomDataTimeline:         msg.CustomDataTimeline,
		ContractAddressTimeline:    msg.ContractAddressTimeline,
		StandardsTimeline:          msg.StandardsTimeline,
		NextBadgeId:        sdk.NewUint(1),
		ParentCollectionId: sdk.NewUint(0),
		BalancesType:       msg.BalancesType,
		IsArchivedTimeline: []*types.IsArchivedTimeline{
			{
				IsArchived: false,
				Times:      []*types.IdRange{
					{
						Start: sdk.NewUint(0),
						End:   sdk.NewUint(math.MaxUint64),
					},
			},
		},
		},
		UnmintedSupplys:    []*types.Balance{},
		TotalSupplys:         []*types.Balance{},
	}

	//Check badge metadata for isFrozen logic
	//TODO:
	// err := AssertIsFrozenLogicIsMaintained([]*types.BadgeMetadata{}, collection.BadgeMetadata)
	// if err != nil {
	// 	return nil, err
	// }

	// err = AssertIsFrozenLogicForApprovedTransfers([]*types.CollectionApprovedTransfer{}, collection.ApprovedTransfers)
	// if err != nil {
	// 	return nil, err
	// }

	if len(msg.BadgesToCreate) != 0 {
		err := *new(error)
		collection, err = k.CreateBadges(ctx, collection, msg.BadgesToCreate, msg.Transfers, msg.Creator)
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
			sdk.NewAttribute("collectionId", fmt.Sprint(NextCollectionId)),
		),
	)

	return &types.MsgNewCollectionResponse{
		CollectionId: NextCollectionId,
	}, nil
}
