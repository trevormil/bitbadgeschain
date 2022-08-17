package keeper_test

import (
	"github.com/trevormil/bitbadgeschain/x/badges/keeper"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

func (suite *TestSuite) TestSetApprovals() {
	userBalanceInfo := types.UserBalanceInfo{}
	subbadgeRanges := []types.IdRange{
		{
			Start: 1,
			End:   0,
		},
		{
			Start: 0,
			End:   0,
		},
		{
			Start: 35,
			End:   0,
		},
		{
			Start: 2,
			End:   34,
		},
		{
			Start: 35,
			End:   100,
		},
	}

	for _, subbadgeRange := range subbadgeRanges {
		userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 1000, aliceAccountNum, &subbadgeRange)
		userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 1000, charlieAccountNum, &subbadgeRange)
	}

	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].Balance, uint64(1000))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 0, End: 100}})


	suite.Require().Equal(userBalanceInfo.Approvals[1].Address, charlieAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[1].ApprovalAmounts[0].Balance, uint64(1000))
	suite.Require().Equal(userBalanceInfo.Approvals[1].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 0, End: 100}})

	subbadgeRangesToRemove := []types.IdRange{
		{
			Start: 1,
			End:   0,
		},
		{
			Start: 0,
			End:   0,
		},
		{
			Start: 35,
			End:   0,
		},
		{
			Start: 35,
			End:   100,
		},
	}

	for _, subbadgeRange := range subbadgeRangesToRemove {
		userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 0, aliceAccountNum, &subbadgeRange)
		userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 0, charlieAccountNum, &subbadgeRange)
	}
	
	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].Balance, uint64(1000))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 2, End: 34}})

	userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 0, aliceAccountNum, &types.IdRange{Start: 2, End: 34})
	userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 0, charlieAccountNum, &types.IdRange{Start: 2, End: 34})
	suite.Require().Equal(len(userBalanceInfo.Approvals), 0)
}

func (suite *TestSuite) TestRemoveApprovals() {
	userBalanceInfo := types.UserBalanceInfo{}
	subbadgeRanges := []types.IdRange{
		{
			Start: 1,
			End:   0,
		},
		{
			Start: 0,
			End:   0,
		},
		{
			Start: 35,
			End:   0,
		},
		{
			Start: 2,
			End:   34,
		},
		{
			Start: 35,
			End:   100,
		},
	}

	for _, subbadgeRange := range subbadgeRanges {
		userBalanceInfo, _ = keeper.SetApproval(suite.ctx, userBalanceInfo, 1000, aliceAccountNum, &subbadgeRange)
	}

	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].Balance, uint64(1000))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 0, End: 100}})

	subbadgeRangesToRemove := []types.IdRange{
		{
			Start: 1,
			End:   0,
		},
		{
			Start: 35,
			End:   0,
		},
		{
			Start: 35,
			End:   100,
		},
	}

	for _, subbadgeRange := range subbadgeRangesToRemove {
		userBalanceInfo, _ = keeper.RemoveBalanceFromApproval(suite.ctx, userBalanceInfo, 1, aliceAccountNum, &subbadgeRange)
	}
	
	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].Balance, uint64(998))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 35, End: 0}})


	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[1].Balance, uint64(999))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[1].IdRanges, []*types.IdRange{{Start: 1, End: 0}, {Start: 36, End: 100}})
}

func (suite *TestSuite) TestAddApprovals() {
	userBalanceInfo := types.UserBalanceInfo{}
	subbadgeRanges := []types.IdRange{
		{
			Start: 1,
			End:   0,
		},
		{
			Start: 0,
			End:   0,
		},
		{
			Start: 35,
			End:   0,
		},
		{
			Start: 2,
			End:   34,
		},
		{
			Start: 35,
			End:   100,
		},
	}

	err := *new(error)
	for _, subbadgeRange := range subbadgeRanges {
		userBalanceInfo, err = keeper.AddBalanceToApproval(suite.ctx, userBalanceInfo, 1000, aliceAccountNum, &subbadgeRange)
		suite.Require().Nil(err, "error adding balance to approval")
	}

	
	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].Balance, uint64(1000))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 0, End: 34}, {Start: 36, End: 100}})

	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[1].Balance, uint64(2000))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[1].IdRanges, []*types.IdRange{{Start: 35, End: 0}})

	subbadgeRangesToRemove := []types.IdRange{
		{
			Start: 1,
			End:   0,
		},
		{
			Start: 35,
			End:   0,
		},
		{
			Start: 35,
			End:   100,
		},
	}

	for _, subbadgeRange := range subbadgeRangesToRemove {
		userBalanceInfo, _ = keeper.AddBalanceToApproval(suite.ctx, userBalanceInfo, 1, aliceAccountNum, &subbadgeRange)
		suite.Require().Nil(err, "error adding balance to approval")
	}
	
	
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].Balance, uint64(1000))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[0].IdRanges, []*types.IdRange{{Start: 0, End: 0}, {Start: 2, End: 34}})

	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[1].Balance, uint64(1001))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[1].IdRanges, []*types.IdRange{{Start: 1, End: 0}, {Start: 36, End: 100}})

	suite.Require().Equal(userBalanceInfo.Approvals[0].Address, aliceAccountNum)
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[2].Balance, uint64(2002))
	suite.Require().Equal(userBalanceInfo.Approvals[0].ApprovalAmounts[2].IdRanges, []*types.IdRange{{Start: 35, End: 0}})
}