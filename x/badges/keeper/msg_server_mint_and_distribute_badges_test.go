package keeper_test

import (
	"math"

	sdkmath "cosmossdk.io/math"
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *TestSuite) TestNewBadges() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating badge: %s")

	err = MintAndDistributeBadges(suite, wctx, &types.MsgMintAndDistributeBadges{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		BadgesToCreate: []*types.Balance{
			{
				Amount: sdkmath.NewUint(1),
				BadgeIds: GetOneIdRange(),
				Times: GetFullIdRanges(),
			},
		},
		Transfers: []*types.Transfer{
			{
				From: bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneIdRange(),
						Times: GetOneIdRange(),
					},
				},
			},
		},
	})
	suite.Require().Nil(err, "Error creating badge: %s")
}

func (suite *TestSuite) TestNewBadgesNotManager() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating badge: %s")

	err = MintAndDistributeBadges(suite, wctx, &types.MsgMintAndDistributeBadges{
		Creator:      alice,
		CollectionId: sdkmath.NewUint(1),
		BadgesToCreate: []*types.Balance{
			{
				Amount: sdkmath.NewUint(1),
				BadgeIds: GetOneIdRange(),
				Times: GetFullIdRanges(),
			},
		},
		Transfers: []*types.Transfer{
			{
				From: bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneIdRange(),
						Times: GetOneIdRange(),
					},
				},
			},
		},
	})
	suite.Require().Error(err, "Error creating badge: %s")
}

func (suite *TestSuite) TestNewBadgeBadgeNotExists() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating badge: %s")

	err = MintAndDistributeBadges(suite, wctx, &types.MsgMintAndDistributeBadges{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		BadgesToCreate: []*types.Balance{
			{
				Amount: sdkmath.NewUint(1),
				BadgeIds: GetOneIdRange(),
				Times: GetFullIdRanges(),
			},
		},
		Transfers: []*types.Transfer{
			{
				From: bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: []*types.IdRange{
							{
								Start: sdkmath.NewUint(2),
								End: sdkmath.NewUint(math.MaxUint64).Add(sdkmath.NewUint(1)),
							},
						},
						Times: GetOneIdRange(),
					},
				},
			},
		},
	})
	suite.Require().Error(err, "Error creating badge: %s")
}