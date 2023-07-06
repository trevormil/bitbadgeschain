package keeper

import (
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CastUserApprovedIncomingTransferPermissionToUniversalPermission(ctx sdk.Context, managerAddress string, permissions []*types.UserApprovedIncomingTransferPermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, permission := range permissions {
		castedCombinations := []*types.UniversalCombination{}
		for _, combination := range permission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				BadgeIdsOptions: combination.BadgeIdsOptions,
				PermittedTimesOptions: combination.PermittedTimesOptions,
				ForbiddenTimesOptions: combination.ForbiddenTimesOptions,
				TimelineTimesOptions: combination.TimelineTimesOptions,
				TransferTimesOptions: combination.TransferTimesOptions,
				FromMappingOptions: combination.FromMappingOptions,
				InitiatedByMappingOptions: combination.InitiatedByMappingOptions,
			})
		}

		fromMapping, err := k.GetAddressMapping(ctx, permission.DefaultValues.FromMappingId, managerAddress)
		if err != nil {
			return nil, err
		}

		initiatedByMapping, err := k.GetAddressMapping(ctx, permission.DefaultValues.InitiatedByMappingId, managerAddress)
		if err != nil {
			return nil, err
		}

		

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				BadgeIds: permission.DefaultValues.BadgeIds,
				TimelineTimes: permission.DefaultValues.TimelineTimes,
				TransferTimes: permission.DefaultValues.TransferTimes,
				FromMapping: fromMapping,
				InitiatedByMapping: initiatedByMapping,
				UsesBadgeIds: true,
				UsesTimelineTimes: true,
				UsesTransferTimes: true,
				UsesToMapping: true,
				UsesInitiatedByMapping: true,
				PermittedTimes: permission.DefaultValues.PermittedTimes,
				ForbiddenTimes: permission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}

func (k Keeper) CastUserApprovedOutgoingTransferPermissionToUniversalPermission(ctx sdk.Context, managerAddress string, permissions []*types.UserApprovedOutgoingTransferPermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, permission := range permissions {
		castedCombinations := []*types.UniversalCombination{}
		for _, combination := range permission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				BadgeIdsOptions: combination.BadgeIdsOptions,
				PermittedTimesOptions: combination.PermittedTimesOptions,
				ForbiddenTimesOptions: combination.ForbiddenTimesOptions,
				TimelineTimesOptions: combination.TimelineTimesOptions,
				TransferTimesOptions: combination.TransferTimesOptions,
				ToMappingOptions: combination.ToMappingOptions,
				InitiatedByMappingOptions: combination.InitiatedByMappingOptions,
			})
		}

		initiatedByMapping, err := k.GetAddressMapping(ctx, permission.DefaultValues.InitiatedByMappingId, managerAddress)
		if err != nil {
			return nil, err
		}

		toMapping, err := k.GetAddressMapping(ctx, permission.DefaultValues.ToMappingId, managerAddress)
		if err != nil {
			return nil, err
		}
		

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				BadgeIds: permission.DefaultValues.BadgeIds,
				TimelineTimes: permission.DefaultValues.TimelineTimes,
				TransferTimes: permission.DefaultValues.TransferTimes,
				ToMapping: toMapping,
				InitiatedByMapping: initiatedByMapping,
				UsesBadgeIds: true,
				UsesTimelineTimes: true,
				UsesTransferTimes: true,
				UsesToMapping: true,
				UsesInitiatedByMapping: true,
				PermittedTimes: permission.DefaultValues.PermittedTimes,
				ForbiddenTimes: permission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}


func (k Keeper) CastActionPermissionToUniversalPermission(actionPermission []*types.ActionPermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, actionPermission := range actionPermission {
		castedCombinations := []*types.UniversalCombination{}
		for _, actionCombination := range actionPermission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				PermittedTimesOptions: actionCombination.PermittedTimesOptions,
				ForbiddenTimesOptions: actionCombination.ForbiddenTimesOptions,
			})
		}

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				PermittedTimes: actionPermission.DefaultValues.PermittedTimes,
				ForbiddenTimes: actionPermission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}

func (k Keeper) CastCollectionApprovedTransferPermissionToUniversalPermission(ctx sdk.Context, managerAddress string, collectionUpdatePermission []*types.CollectionApprovedTransferPermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, collectionUpdatePermission := range collectionUpdatePermission {
		castedCombinations := []*types.UniversalCombination{}
		for _, collectionUpdateCombination := range collectionUpdatePermission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				PermittedTimesOptions: collectionUpdateCombination.PermittedTimesOptions,
				ForbiddenTimesOptions: collectionUpdateCombination.ForbiddenTimesOptions,
				TimelineTimesOptions: collectionUpdateCombination.TimelineTimesOptions,
				TransferTimesOptions: collectionUpdateCombination.TransferTimesOptions,
				ToMappingOptions: collectionUpdateCombination.ToMappingOptions,
				FromMappingOptions: collectionUpdateCombination.FromMappingOptions,
				InitiatedByMappingOptions: collectionUpdateCombination.InitiatedByMappingOptions,
				BadgeIdsOptions: collectionUpdateCombination.BadgeIdsOptions,
			})
		}

		fromMapping, err := k.GetAddressMapping(ctx, collectionUpdatePermission.DefaultValues.FromMappingId, managerAddress)
		if err != nil {
			return nil, err
		}

		initiatedByMapping, err := k.GetAddressMapping(ctx, collectionUpdatePermission.DefaultValues.InitiatedByMappingId, managerAddress)
		if err != nil {
			return nil, err
		}

		toMapping, err := k.GetAddressMapping(ctx, collectionUpdatePermission.DefaultValues.ToMappingId, managerAddress)
		if err != nil {
			return nil, err
		}

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				TimelineTimes: collectionUpdatePermission.DefaultValues.TimelineTimes,
				TransferTimes: collectionUpdatePermission.DefaultValues.TransferTimes,
				ToMapping: toMapping,
				FromMapping: fromMapping,
				InitiatedByMapping: initiatedByMapping,
				BadgeIds: collectionUpdatePermission.DefaultValues.BadgeIds,
				UsesBadgeIds: true,
				UsesTimelineTimes: true,
				UsesTransferTimes: true,
				UsesToMapping: true,
				UsesFromMapping: true,
				UsesInitiatedByMapping: true,
				PermittedTimes: collectionUpdatePermission.DefaultValues.PermittedTimes,
				ForbiddenTimes: collectionUpdatePermission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}


func (k Keeper) CastTimedUpdateWithBadgeIdsPermissionToUniversalPermission(timedUpdateWithBadgeIdsPermission []*types.TimedUpdateWithBadgeIdsPermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, timedUpdateWithBadgeIdsPermission := range timedUpdateWithBadgeIdsPermission {
		castedCombinations := []*types.UniversalCombination{}
		for _, timedUpdateWithBadgeIdsCombination := range timedUpdateWithBadgeIdsPermission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				BadgeIdsOptions: timedUpdateWithBadgeIdsCombination.BadgeIdsOptions,
				PermittedTimesOptions: timedUpdateWithBadgeIdsCombination.PermittedTimesOptions,
				ForbiddenTimesOptions: timedUpdateWithBadgeIdsCombination.ForbiddenTimesOptions,
				TimelineTimesOptions: timedUpdateWithBadgeIdsCombination.TimelineTimesOptions,
			})
		}

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				TimelineTimes: timedUpdateWithBadgeIdsPermission.DefaultValues.TimelineTimes,
				BadgeIds: timedUpdateWithBadgeIdsPermission.DefaultValues.BadgeIds,
				UsesTimelineTimes: true,
				UsesBadgeIds: true,
				PermittedTimes: timedUpdateWithBadgeIdsPermission.DefaultValues.PermittedTimes,
				ForbiddenTimes: timedUpdateWithBadgeIdsPermission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}

func (k Keeper) CastTimedUpdatePermissionToUniversalPermission(timedUpdatePermission []*types.TimedUpdatePermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, timedUpdatePermission := range timedUpdatePermission {
		castedCombinations := []*types.UniversalCombination{}
		for _, timedUpdateCombination := range timedUpdatePermission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				PermittedTimesOptions: timedUpdateCombination.PermittedTimesOptions,
				ForbiddenTimesOptions: timedUpdateCombination.ForbiddenTimesOptions,
				TimelineTimesOptions: timedUpdateCombination.TimelineTimesOptions,
			})
		}

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				TimelineTimes: timedUpdatePermission.DefaultValues.TimelineTimes,
				UsesTimelineTimes: true,
				PermittedTimes: timedUpdatePermission.DefaultValues.PermittedTimes,
				ForbiddenTimes: timedUpdatePermission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}


func (k Keeper) CastActionWithBadgeIdsAndTimesPermissionToUniversalPermission(ActionWithBadgeIdsAndTimesPermission []*types.ActionWithBadgeIdsAndTimesPermission) ([]*types.UniversalPermission, error) {
	castedPermissions := []*types.UniversalPermission{}
	for _, ActionWithBadgeIdsAndTimesPermission := range ActionWithBadgeIdsAndTimesPermission {
		castedCombinations := []*types.UniversalCombination{}
		for _, ActionWithBadgeIdsAndTimesCombination := range ActionWithBadgeIdsAndTimesPermission.Combinations {
			castedCombinations = append(castedCombinations, &types.UniversalCombination{
				BadgeIdsOptions: ActionWithBadgeIdsAndTimesCombination.BadgeIdsOptions,
				TransferTimesOptions: ActionWithBadgeIdsAndTimesCombination.TransferTimesOptions,
				PermittedTimesOptions: ActionWithBadgeIdsAndTimesCombination.PermittedTimesOptions,
				ForbiddenTimesOptions: ActionWithBadgeIdsAndTimesCombination.ForbiddenTimesOptions,
			})
		}

		castedPermissions = append(castedPermissions, &types.UniversalPermission{
			DefaultValues: &types.UniversalDefaultValues{
				BadgeIds: ActionWithBadgeIdsAndTimesPermission.DefaultValues.BadgeIds,
				TransferTimes: ActionWithBadgeIdsAndTimesPermission.DefaultValues.TransferTimes,
				UsesBadgeIds: true,
				UsesTransferTimes: true,
				PermittedTimes: ActionWithBadgeIdsAndTimesPermission.DefaultValues.PermittedTimes,
				ForbiddenTimes: ActionWithBadgeIdsAndTimesPermission.DefaultValues.ForbiddenTimes,
			},
			Combinations: castedCombinations,
		})
	}
	return castedPermissions, nil
}