package keeper

import (
	"context"
	"math"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"github.com/bitbadges/bitbadgeschain/x/maps/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	badgetypes "github.com/bitbadges/bitbadgeschain/x/badges/types"
)

func (k msgServer) SetValue(goCtx context.Context, msg *types.MsgSetValue) (*types.MsgSetValueResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	mapId := msg.MapId
	key := msg.Key
	value := msg.Value

	//Check the overwrite options
	if msg.Options.UseMostRecentCollectionId {
		nextCollectionId := k.badgesKeeper.GetNextCollectionId(ctx)
		value = nextCollectionId.Sub(sdk.NewUint(1)).BigInt().String()
	}

	currMap, found := k.GetMapFromStore(ctx, mapId)
	if !found {
		return nil, sdkerrors.Wrap(ErrMapDoesNotExist, "Failed to get map from store")
	}

	// bool expectUint = 3;
	// bool expectBoolean = 4;
	// bool expectAddress = 5;
	// bool expectUri = 6;

	if currMap.ValueOptions.ExpectUint {
		newUint := sdkmath.NewUintFromString(value)
		if newUint.IsNil() {
			return nil, sdkerrors.Wrap(ErrInvalidValue, "Value must be a valid uint")
		}
	}

	if currMap.ValueOptions.ExpectBoolean {
		if value != "true" && value != "false" {
			return nil, sdkerrors.Wrap(ErrInvalidValue, "Value must be a valid boolean")
		}
	}

	if currMap.ValueOptions.ExpectAddress {
		_, err := sdk.AccAddressFromBech32(value)
		if err != nil {
			return nil, sdkerrors.Wrap(ErrInvalidValue, "Value must be a valid address")
		}
	}

	if currMap.ValueOptions.ExpectUri {
		err := badgetypes.ValidateURI(value)
		if err != nil {
			return nil, sdkerrors.Wrap(ErrInvalidValue, "Value must be a valid URI")
		}
	}

	currVal := k.GetMapValueFromStore(ctx, mapId, key)
	if currVal.Value == value {
		return nil, sdkerrors.Wrap(ErrValueAlreadySet, "Value cannot be the same as the current value")
	}

	collection := &badgetypes.BadgeCollection{}
	if !currMap.InheritManagerTimelineFrom.IsNil() && !currMap.InheritManagerTimelineFrom.IsZero() {
		collectionRes, err := k.badgesKeeper.GetCollection(ctx, &badgetypes.QueryGetCollectionRequest{CollectionId: currMap.InheritManagerTimelineFrom})
		if err != nil {
			return nil, sdkerrors.Wrap(ErrInvalidMapId, "Could not find collection in store")
		}

		collection = collectionRes.Collection
	}

	currManager := types.GetCurrentManagerForMap(ctx, currMap, collection)

	forceUpdate := false
	canUpdate := false
	if currMap.UpdateCriteria.ManagerOnly && currManager == msg.Creator {
		canUpdate = true
	}

	if !currMap.UpdateCriteria.CollectionId.IsNil() && currMap.UpdateCriteria.CollectionId.GT(sdkmath.NewUint(0)) {
		badgeId := sdkmath.NewUintFromString(key)
		balancesRes, err := k.badgesKeeper.GetBalance(ctx, &badgetypes.QueryGetBalanceRequest{
			CollectionId: currMap.UpdateCriteria.CollectionId,
			Address:      msg.Creator,
		})
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Failed to get balance")
		}

		currTime := sdkmath.NewUint(uint64(ctx.BlockTime().UnixMilli()))
		bals, err := badgetypes.GetBalancesForIds(ctx, []*badgetypes.UintRange{
			{Start: badgeId, End: badgeId},
		}, []*badgetypes.UintRange{
			{Start: currTime, End: currTime},
		}, balancesRes.Balance.Balances)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Failed to get balances for ids")
		}

		for _, bal := range bals {
			if bal.Amount.GTE(sdkmath.NewUint(1)) {
				canUpdate = true
				break
			}
		}
	}

	if currMap.UpdateCriteria.CreatorOnly && key == msg.Creator {
		canUpdate = true
	}

	if currMap.UpdateCriteria.FirstComeFirstServe {

		if currVal.Value == "" {
			canUpdate = true
		}

		if currVal.Value != "" && currVal.LastSetBy == msg.Creator {
			canUpdate = true
		}
	}

	if currManager == msg.Creator {
		canForceUpdate := k.CheckIfKeyIsEditable(ctx, currMap.Permissions.CanForceEdit, key)
		if canForceUpdate {
			canUpdate = true
			forceUpdate = true
		}
	}

	if currMap.ValueOptions.PermanentOnceSet {
		if currVal.Value != "" {
			return nil, sdkerrors.Wrap(ErrValueAlreadySet, "Value already set and cannot be updated")
		}
	}

	if currMap.ValueOptions.NoDuplicates {
		if k.GetMapDuplicateValueFromStore(ctx, mapId, value) {
			return nil, sdkerrors.Wrap(ErrDuplicateValue, "Duplicate values not allowed")
		}

		if currVal.Value != "" {
			k.DeleteMapDuplicateValueFromStore(ctx, mapId, currVal.Value)
		}

		err := k.SetMapDuplicateValueInStore(ctx, mapId, value)
		if err != nil {
			return nil, sdkerrors.Wrap(err, "Failed to set duplicate tracker in store")
		}
	}

	if !forceUpdate {
		canEdit := k.CheckIfKeyIsEditable(ctx, currMap.Permissions.CanEdit, key)
		if !canEdit {
			return nil, sdkerrors.Wrap(ErrMapNotEditable, "Map is not editable")
		}
	}

	if !canUpdate {
		return nil, sdkerrors.Wrap(ErrCannotUpdateMapValue, "Cannot update map value")
	}

	err := k.SetMapValueInStore(ctx, mapId, key, value, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Failed to update map value in store")
	}

	return &types.MsgSetValueResponse{}, nil
}

func (k Keeper) CheckIfKeyIsEditable(ctx sdk.Context, permissions []*types.IsEditablePermission, key string) (bool) {
	toCheck := []*badgetypes.UniversalPermissionDetails{
		{
			ApprovalIdList: &badgetypes.AddressList{
				Whitelist: true,
				Addresses: []string{key},
			},
			ToList: &badgetypes.AddressList{
				Whitelist: false,
				Addresses: []string{},
			},
			FromList: &badgetypes.AddressList{
				Whitelist: false,
				Addresses: []string{},
			},
			InitiatedByList: &badgetypes.AddressList{
				Whitelist: false,
				Addresses: []string{},
			},
			BadgeId: &badgetypes.UintRange{
				Start: sdkmath.NewUint(math.MaxUint64), End: sdkmath.NewUint(math.MaxUint64),
			},
			TimelineTime: &badgetypes.UintRange{
				Start: sdkmath.NewUint(math.MaxUint64), End: sdkmath.NewUint(math.MaxUint64),
			},
			TransferTime: &badgetypes.UintRange{
				Start: sdkmath.NewUint(math.MaxUint64), End: sdkmath.NewUint(math.MaxUint64),
			},
			OwnershipTime: &badgetypes.UintRange{
				Start: sdkmath.NewUint(math.MaxUint64), End: sdkmath.NewUint(math.MaxUint64),
			},
		},
	}

	err := k.badgesKeeper.CheckIfCollectionApprovalPermissionPermits(ctx, toCheck, types.CastIsEditablePermissions(permissions), "can edit")
	if err != nil {
		return false
	}

	return true
}
