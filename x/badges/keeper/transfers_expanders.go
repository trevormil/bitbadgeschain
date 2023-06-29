package keeper

import (
	"math"

	sdkmath "cosmossdk.io/math"
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
)

func ExpandCollectionApprovedTransfers(approvedTransfers []*types.CollectionApprovedTransfer) []*types.CollectionApprovedTransfer {
	newCurrApprovedTransfers := []*types.CollectionApprovedTransfer{}
	for _, approvedTransfer := range approvedTransfers {
		for _, allowedCombination := range approvedTransfer.AllowedCombinations {
			badgeIds := approvedTransfer.BadgeIds
			if allowedCombination.InvertBadgeIds {
				badgeIds = types.InvertIdRanges(badgeIds, sdkmath.NewUint(math.MaxUint64))
			}

			times := approvedTransfer.TransferTimes
			if allowedCombination.InvertTransferTimes {
				times = types.InvertIdRanges(times, sdkmath.NewUint(math.MaxUint64))
			}

			toMappingId := approvedTransfer.ToMappingId
			if allowedCombination.InvertTo {
				toMappingId = "!" + toMappingId
			}

			fromMappingId := approvedTransfer.FromMappingId
			if allowedCombination.InvertFrom {
				fromMappingId = "!" + fromMappingId
			}

			initiatedByMappingId := approvedTransfer.InitiatedByMappingId
			if allowedCombination.InvertInitiatedBy {
				initiatedByMappingId = "!" + initiatedByMappingId
			}

			newCurrApprovedTransfers = append(newCurrApprovedTransfers, &types.CollectionApprovedTransfer{
				ToMappingId: toMappingId,
				FromMappingId: fromMappingId,
				InitiatedByMappingId: initiatedByMappingId,
				TransferTimes: times,
				BadgeIds: badgeIds,
				AllowedCombinations: []*types.IsCollectionTransferAllowed{
					{
						IsAllowed: allowedCombination.IsAllowed,
					},
				},
				OverallApprovals: approvedTransfer.OverallApprovals,
				PerAddressApprovals: approvedTransfer.PerAddressApprovals,
				IncrementIdsBy: approvedTransfer.IncrementIdsBy,
				IncrementTimesBy: approvedTransfer.IncrementTimesBy,
				RequireFromEqualsInitiatedBy: approvedTransfer.RequireFromEqualsInitiatedBy,
				RequireFromDoesNotEqualInitiatedBy: approvedTransfer.RequireFromDoesNotEqualInitiatedBy,
				RequireToEqualsInitiatedBy: approvedTransfer.RequireToEqualsInitiatedBy,
				RequireToDoesNotEqualInitiatedBy: approvedTransfer.RequireToDoesNotEqualInitiatedBy,
				OverridesFromApprovedOutgoingTransfers: approvedTransfer.OverridesFromApprovedOutgoingTransfers,
				OverridesToApprovedIncomingTransfers: approvedTransfer.OverridesToApprovedIncomingTransfers,
				CustomData: approvedTransfer.CustomData,
				Uri: approvedTransfer.Uri,
				TrackerId: approvedTransfer.TrackerId,
				Challenges: approvedTransfer.Challenges,
			})
		}
	}

	return newCurrApprovedTransfers
}

func AppendDefaultForIncoming(currApprovedTransfers []*types.UserApprovedIncomingTransfer, userAddress string) []*types.UserApprovedIncomingTransfer {
	currApprovedTransfers = append(currApprovedTransfers, &types.UserApprovedIncomingTransfer{
		FromMappingId: "All", //everyone
		InitiatedByMappingId: userAddress,
		TransferTimes: []*types.IdRange{
			{
				Start: sdkmath.NewUint(0),
				End: sdkmath.NewUint(uint64(math.MaxUint64)),
			},
		},
		BadgeIds: []*types.IdRange{
			{
				Start: sdkmath.NewUint(1),
				End: sdkmath.NewUint(math.MaxUint64),
			},
		},
		AllowedCombinations: []*types.IsUserIncomingTransferAllowed{
			{
				IsAllowed: true,
			},
		},
	})

	return currApprovedTransfers
}

func AppendDefaultForOutgoing(currApprovedTransfers []*types.UserApprovedOutgoingTransfer, userAddress string) []*types.UserApprovedOutgoingTransfer {
	currApprovedTransfers = append(currApprovedTransfers, &types.UserApprovedOutgoingTransfer{
		ToMappingId: "All", //everyone
		InitiatedByMappingId: userAddress,
		TransferTimes: []*types.IdRange{
			{
				Start: sdkmath.NewUint(0),
				End: sdkmath.NewUint(uint64(math.MaxUint64)),
			},
		},
		BadgeIds: []*types.IdRange{
			{
				Start: sdkmath.NewUint(1),
				End: sdkmath.NewUint(math.MaxUint64),
			},
		},
		AllowedCombinations: []*types.IsUserOutgoingTransferAllowed{
			{
				IsAllowed: true,
			},
		},
	})

	return currApprovedTransfers
}

