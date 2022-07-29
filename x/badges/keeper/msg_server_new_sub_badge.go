package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

func (k msgServer) NewSubBadge(goCtx context.Context, msg *types.MsgNewSubBadge) (*types.MsgNewSubBadgeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Creator will already be registered, so we can do this and panic if it fails
	creator_account_num := k.Keeper.MustGetAccountNumberForAddressString(ctx, msg.Creator)

	if msg.Supply == 0 {
		return nil, ErrSupplyEqualsZero
	}

	badge, found := k.GetBadgeFromStore(ctx, msg.Id)
	if !found {
		return nil, ErrBadgeNotExists
	}

	if badge.Manager != creator_account_num {
		return nil, ErrSenderIsNotManager
	}

	//Check permissions (can_create)
	permission_flags := GetPermissions(badge.PermissionFlags)
	if !permission_flags.can_create {
		return nil, sdkerrors.Wrapf(ErrInvalidPermissions, "%s", "Manager can't create more subbadges")
	}

	//Once here, we are safe to mint
	//By default, we assume non fungible subbadge (i.e. supply == 1) so we don't store if supply == 1
	subasset_id := badge.NextSubassetId
	if msg.Supply != 1 {
		badge.SubassetsTotalSupply = append(badge.SubassetsTotalSupply, &types.Subasset{
			Id:     subasset_id,
			Supply: msg.Supply,
		})
	}
	badge.NextSubassetId += 1

	//Mint the total supplt of subbadge to the manager
	manager_balance_id := GetBalanceKey(creator_account_num, msg.Id, subasset_id)
	if err := k.AddToBadgeBalance(ctx, manager_balance_id, msg.Supply); err != nil {
		return nil, err
	}
	
	if err := k.UpdateBadgeInStore(ctx, badge); err != nil {
		return nil, err
	}

	return &types.MsgNewSubBadgeResponse{
		SubassetId: subasset_id,
	}, nil
}
