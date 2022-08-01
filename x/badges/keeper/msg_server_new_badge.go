package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

func (k msgServer) NewBadge(goCtx context.Context, msg *types.MsgNewBadge) (*types.MsgNewBadgeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Creator will already be registered, so we can do this and panic if it fails
	creator_account_num := k.Keeper.MustGetAccountNumberForAddressString(ctx, msg.Creator)

	//Get next badge ID and increment
	next_badge_id := k.GetNextAssetId(ctx)
	k.IncrementNextAssetId(ctx)

	//Create and store the badge
	badge := types.BitBadge{
		Id:                              next_badge_id,
		Uri:                             msg.Uri,
		Manager:                         creator_account_num,
		PermissionFlags:                 msg.Permissions,
		SubassetUriFormat:               msg.SubassetUris,
		SubassetsTotalSupply:            []*types.Subasset{},
		NextSubassetId:                  0,
		FrozenOrUnfrozenAddressesDigest: msg.FreezeAddressesDigest,
	}

	if err := k.SetBadgeInStore(ctx, badge); err != nil {
		return nil, err
	}
	
	return &types.MsgNewBadgeResponse{
		Id:      next_badge_id,
	}, nil
}
