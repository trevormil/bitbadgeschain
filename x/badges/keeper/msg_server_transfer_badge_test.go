package keeper_test

import (
	sdkmath "cosmossdk.io/math"

	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *TestSuite) TestTransferBadgeForceful() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating collections")

	bobbalance, _ := GetUserBalance(suite, wctx, sdkmath.NewUint(1), bob)

	fetchedBalance, err := types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), bobbalance.Balances)
	suite.Require().Equal(sdkmath.NewUint(1), fetchedBalance[0].Amount)
	suite.Require().Nil(err)

	err = TransferBadge(suite, wctx, &types.MsgTransferBadge{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		Transfers: []*types.Transfer{
			{
				From:         bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneUintRange(),
						OwnershipTimes: GetFullUintRanges(),
					},
				},
			},
		},
	})
	suite.Require().Nil(err, "Error transferring badge")

	bobbalance, _ = GetUserBalance(suite, wctx, sdkmath.NewUint(1), bob)
	fetchedBalance, err = types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), bobbalance.Balances)
	AssertUintsEqual(suite, sdkmath.NewUint(0), fetchedBalance[0].Amount)
	suite.Require().Nil(err)

	alicebalance, _ := GetUserBalance(suite, wctx, sdkmath.NewUint(1), alice)
	fetchedBalance, err = types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), alicebalance.Balances)
	AssertUintsEqual(suite, sdkmath.NewUint(1), fetchedBalance[0].Amount)
	suite.Require().Nil(err)
}


func (suite *TestSuite) TestTransferBadgeHandleDuplicateIDs() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating collections")

	bobbalance, _ := GetUserBalance(suite, wctx, sdkmath.NewUint(1), bob)

	fetchedBalance, err := types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), bobbalance.Balances)
	suite.Require().Equal(sdkmath.NewUint(1), fetchedBalance[0].Amount)
	suite.Require().Nil(err)

	err = TransferBadge(suite, wctx, &types.MsgTransferBadge{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		Transfers: []*types.Transfer{
			{
				From:         bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: []*types.UintRange{
							GetOneUintRange()[0],
							GetOneUintRange()[0],
						},
						OwnershipTimes: GetFullUintRanges(),
					},
				},
			},
		},
	})
	suite.Require().Error(err, "Error transferring badge")

}

func (suite *TestSuite) TestTransferBadgeNotApprovedCollectionLevel() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating collections")

	bobbalance, _ := GetUserBalance(suite, wctx, sdkmath.NewUint(1), bob)

	fetchedBalance, err := types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), bobbalance.Balances)
	suite.Require().Equal(sdkmath.NewUint(1), fetchedBalance[0].Amount)
	suite.Require().Nil(err)

	err = UpdateCollectionApprovedTransfers(suite, wctx, &types.MsgUpdateCollectionApprovedTransfers{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		CollectionApprovedTransfersTimeline: []*types.CollectionApprovedTransferTimeline{
			{
				ApprovedTransfers: []*types.CollectionApprovedTransfer{},
				TimelineTimes: GetFullUintRanges(),
			},
		},
	})
	suite.Require().Nil(err, "Error updating approved transfers")

	err = TransferBadge(suite, wctx, &types.MsgTransferBadge{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		Transfers: []*types.Transfer{
			{
				From:         bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneUintRange(),
						OwnershipTimes: GetFullUintRanges(),
					},
				},
			},
		},
	})
	suite.Require().Error(err, "Error transferring badge")
}


func (suite *TestSuite) TestTransferBadgeNotApprovedIncoming() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating collections")

	bobbalance, _ := GetUserBalance(suite, wctx, sdkmath.NewUint(1), bob)

	fetchedBalance, err := types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), bobbalance.Balances)
	suite.Require().Equal(sdkmath.NewUint(1), fetchedBalance[0].Amount)
	suite.Require().Nil(err)

	err = UpdateUserApprovedTransfers(suite, wctx, &types.MsgUpdateUserApprovedTransfers{
		Creator:      alice,
		CollectionId: sdkmath.NewUint(1),
		ApprovedIncomingTransfersTimeline: []*types.UserApprovedIncomingTransferTimeline{
			{
				// ApprovedIncomingTransfers: []*types.CollectionApprovedTransfer{},
				TimelineTimes: GetFullUintRanges(),
			},
		},
	})
	suite.Require().Nil(err, "Error updating approved transfers")

	err = TransferBadge(suite, wctx, &types.MsgTransferBadge{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		Transfers: []*types.Transfer{
			{
				From:         bob,
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneUintRange(),
						OwnershipTimes: GetFullUintRanges(),
					},
				},
			},
		},
	})
	suite.Require().Error(err, "Error transferring badge")
}

func (suite *TestSuite) TestTransferBadgeFromMintAddress() {
	wctx := sdk.WrapSDKContext(suite.ctx)

	collectionsToCreate := GetTransferableCollectionToCreateAllMintedToCreator(bob)

	err := CreateCollections(suite, wctx, collectionsToCreate)
	suite.Require().Nil(err, "Error creating collections")

	bobbalance, _ := GetUserBalance(suite, wctx, sdkmath.NewUint(1), bob)

	fetchedBalance, err := types.GetBalancesForIds(GetOneUintRange(), GetOneUintRange(), bobbalance.Balances)
	suite.Require().Equal(sdkmath.NewUint(1), fetchedBalance[0].Amount)
	suite.Require().Nil(err)

	unmintedSupplys, err := GetUserBalance(suite, wctx, sdk.NewUint(1), "Mint")
	suite.Require().Nil(err, "Error getting user balance: %s")
	AssertBalancesEqual(suite, []*types.Balance{}, unmintedSupplys.Balances)

	err = MintAndDistributeBadges(suite, wctx, &types.MsgMintAndDistributeBadges{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		BadgesToCreate: []*types.Balance{
			{
				Amount: sdkmath.NewUint(1),
				BadgeIds: GetOneUintRange(),
				OwnershipTimes: GetFullUintRanges(),
			},
		},
	})
	suite.Require().Nil(err, "Error transferring badge")

	err = TransferBadge(suite, wctx, &types.MsgTransferBadge{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		Transfers: []*types.Transfer{
			{
				From:         "Mint",
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneUintRange(),
						OwnershipTimes: GetFullUintRanges(),
					},
				},
			},
		},
	})
	suite.Require().Nil(err, "Error transferring badge")

	unmintedSupplys, err = GetUserBalance(suite, wctx, sdk.NewUint(1), "Mint")
	suite.Require().Nil(err, "Error getting user balance: %s")
	AssertBalancesEqual(suite, []*types.Balance{}, unmintedSupplys.Balances)

	err = TransferBadge(suite, wctx, &types.MsgTransferBadge{
		Creator:      bob,
		CollectionId: sdkmath.NewUint(1),
		Transfers: []*types.Transfer{
			{
				From:         "Mint",
				ToAddresses: []string{alice},
				Balances: []*types.Balance{
					{
						Amount: sdkmath.NewUint(1),
						BadgeIds: GetOneUintRange(),
						OwnershipTimes: GetFullUintRanges(),
					},
				},
			},
		},
	})
	suite.Require().Error(err, "Error transferring badge")
}