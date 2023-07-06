package keeper_test

import (
	sdkmath "cosmossdk.io/math"
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *TestSuite) TestUpdateManager() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetCollectionsToCreate()

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating badge")

	err = UpdateManager(suite, wctx, &types.MsgUpdateManager{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		ManagerTimeline: []*types.ManagerTimeline{
			{
				Manager: alice,
				TimelineTimes: GetFullIdRanges(),
			},
		},
	})
	suite.Require().Nil(err, "Error transferring manager")

	collection, _ := GetCollection(suite, wctx, sdkmath.NewUint(1))
	suite.Require().Equal(alice, types.GetCurrentManager(suite.ctx, collection))
}