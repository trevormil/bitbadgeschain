package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *TestSuite) TestAccountsAreCreatedAndStoredCorrectly() {
	sample := "cosmos1f6k8dr0hzafe78uhnmcg7e9qm07ejue8mmd8xq"
	sampleAccNumber := suite.app.BadgesKeeper.GetOrCreateAccountNumberForAccAddressBech32(suite.ctx, sdk.MustAccAddressFromBech32(sample))
	suite.Equal(sampleAccNumber, uint64(10))
}
