package types_test

import (
	math "math"

	sdkmath "cosmossdk.io/math"
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
)

func GetValidUserPermissions() *types.UserPermissions {
	return &types.UserPermissions{
		CanUpdateApprovedOutgoingTransfers: []*types.UserApprovedTransferPermission{
			{
				DefaultValues: &types.UserApprovedOutgoingTransferDefaultValues{
					PermittedTimes: []*types.IdRange{ { Start: sdkmath.NewUint(1), End: sdkmath.NewUint(2) } },
					ForbiddenTimes: []*types.IdRange{ { Start: sdkmath.NewUint(6), End: sdkmath.NewUint(8) } },
				},
				Combinations: []*types.UserApprovedOutgoingTransferCombination{{}},
			},
		},
	}
}

func GetValidCollectionMetadataTimeline() []*types.CollectionMetadataTimeline {
	return []*types.CollectionMetadataTimeline{
		{
			CollectionMetadata: &types.CollectionMetadata{
				Uri: "https://example.com/{id}",
			},
			Times: []*types.IdRange{
				{
					Start: sdkmath.NewUint(1),
					End:  sdkmath.NewUint(math.MaxUint64),
				},
			},
		},
	}
}

func GetValidBadgeMetadataTimeline() []*types.BadgeMetadataTimeline {
	return []*types.BadgeMetadataTimeline{
		{
			BadgeMetadata: []*types.BadgeMetadata{
				{
					Uri: "https://example.com/{id}",
					BadgeIds: []*types.IdRange{
						{
							Start: sdkmath.NewUint(1),

							End:   sdkmath.NewUint(math.MaxUint64),
						},
					},
				},
			},
		},
	}
}

