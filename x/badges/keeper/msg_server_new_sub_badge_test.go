package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/trevormil/bitbadgeschain/x/badges/keeper"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

func (suite *TestSuite) TestNewSubBadges() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	badgesToCreate := []BadgesToCreate{
		{
			Badge: types.MsgNewBadge{
				Uri: validUri,
				Permissions: 62,
				FreezeAddressesDigest: "",
				SubassetUris: validUri,
			},
			Amount: 1,
			Creator: bob,
		},
	}

	CreateBadges(suite, wctx, badgesToCreate)
	badge := GetBadge(suite, wctx, 0)

	//Create subbadge 1 with supply > 1
	err := CreateSubBadge(suite, wctx, bob, 0, 10)
	suite.Require().Nil(err, "Error creating subbadge")
	badge = GetBadge(suite, wctx, 0)
	bobBalanceInfo := GetBadgeBalance(suite, wctx, 0, 0, firstAccountNumCreated)

	suite.Require().Equal(uint64(1), badge.NextSubassetId)
	suite.Require().Equal([]*types.Subasset{
		{
			Id: 0,
			Supply: 10,
		},
	}, badge.SubassetsTotalSupply)
	suite.Require().Equal(uint64(10), bobBalanceInfo.Balance)

	//Create subbadge 2 with supply == 1
	err = CreateSubBadge(suite, wctx, bob, 0, 1)
	suite.Require().Nil(err, "Error creating subbadge")
	
	badge = GetBadge(suite, wctx, 0)
	bobBalanceInfo = GetBadgeBalance(suite, wctx, 0, 1, firstAccountNumCreated)

	suite.Require().Equal(uint64(2), badge.NextSubassetId)
	suite.Require().Equal([]*types.Subasset{
		{
			Id: 0,
			Supply: 10,
		},
	}, badge.SubassetsTotalSupply)
	suite.Require().Equal(uint64(1), bobBalanceInfo.Balance)

	//Create subbadge 2 with supply == 10
	err = CreateSubBadge(suite, wctx, bob, 0, 10)
	suite.Require().Nil(err, "Error creating subbadge")
	badge = GetBadge(suite, wctx, 0)
	bobBalanceInfo = GetBadgeBalance(suite, wctx, 0, 2, firstAccountNumCreated)

	suite.Require().Equal(uint64(3), badge.NextSubassetId)
	suite.Require().Equal([]*types.Subasset{
		{
			Id: 0,
			Supply: 10,
		},
		{
			Id: 2,
			Supply: 10,
		},
	}, badge.SubassetsTotalSupply)
	suite.Require().Equal(uint64(10), bobBalanceInfo.Balance)
}


func (suite *TestSuite) TestNewSubBadgesNotManager() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	badgesToCreate := []BadgesToCreate{
		{
			Badge: types.MsgNewBadge{
				Uri: validUri,
				Permissions: 62,
				FreezeAddressesDigest: "",
				SubassetUris: validUri,
			},
			Amount: 1,
			Creator: bob,
		},
	}

	CreateBadges(suite, wctx, badgesToCreate)
	err := CreateSubBadge(suite, wctx, alice, 0, 10)
	suite.Require().EqualError(err, keeper.ErrSenderIsNotManager.Error())
}

func (suite *TestSuite) TestNewSubBadgeBadgeNotExists() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	err := CreateSubBadge(suite, wctx, alice, 0, 10)
	suite.Require().EqualError(err, keeper.ErrBadgeNotExists.Error())
}

func (suite *TestSuite) TestNewSubBadgeCreateIsLocked() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	badgesToCreate := []BadgesToCreate{
		{
			Badge: types.MsgNewBadge{
				Uri: validUri,
				Permissions: 0,
				FreezeAddressesDigest: "",
				SubassetUris: validUri,
			},
			Amount: 1,
			Creator: bob,
		},
	}

	CreateBadges(suite, wctx, badgesToCreate)
	err := CreateSubBadge(suite, wctx, bob, 0, 10)
	suite.Require().EqualError(err, keeper.ErrInvalidPermissions.Error())
}