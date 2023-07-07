package keeper

import (
	"context"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteCollection(goCtx context.Context, msg *types.MsgDeleteCollection) (*types.MsgDeleteCollectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	collection, found := k.GetCollectionFromStore(ctx, msg.CollectionId)
	if !found {
		return nil, ErrCollectionNotExists
	}

	err := k.UniversalValidate(ctx, collection, UniversalValidationParams{
		Creator:       msg.Creator,
		MustBeManager: true,
	})
	if err != nil {
		return nil, err
	}

	//Check deleted permission is valid for current time
	err = k.CheckActionPermission(ctx, collection.Permissions.CanDeleteCollection)
	if err != nil {
		return nil, err
	}

	k.DeleteCollectionFromStore(ctx, collection.CollectionId)
	if err != nil {
		return nil, err
	}

	//TODO: should we prune all balances and challenge stores here too?

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, "badges"),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	)

	return &types.MsgDeleteCollectionResponse{}, nil
}
