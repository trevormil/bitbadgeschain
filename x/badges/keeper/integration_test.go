package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/bitbadges/bitbadgeschain/x/badges/keeper"
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	bitbadgesapp "github.com/bitbadges/bitbadgeschain/app"

	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"
)

const (
	//Note these are alphanumerically sorted (needed for approvals test)
	alice   = "cosmos1e0w5t53nrq7p66fye6c8p0ynyhf6y24l4yuxd7"
	bob     = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d3"
	charlie = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8g"
)

// var DefaultConsensusParams = &abci.ConsensusParams{
// 	Block: &abci.BlockParams{
// 		MaxBytes: 200000,
// 		MaxGas:   2000000,
// 	},
// 	Evidence: &tmproto.EvidenceParams{
// 		MaxAgeNumBlocks: 302400,
// 		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
// 		MaxBytes:        10000,
// 	},
// 	Validator: &tmproto.ValidatorParams{
// 		PubKeyTypes: []string{
// 			tmtypes.ABCIPubKeyTypeEd25519,
// 		},
// 	},
// }

type TestSuite struct {
	suite.Suite

	app         *bitbadgesapp.App
	ctx         sdk.Context
	queryClient types.QueryClient
	msgServer   types.MsgServer
}

// Bunch of weird config stuff to setup the app. Inherited most from Cosmos SDK tutorials and existing Cosmos SDK modules.
func (suite *TestSuite) SetupTest() {
	app := bitbadgesapp.Setup(
		false,
	)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.ctx = ctx
	suite.msgServer = keeper.NewMsgServerImpl(app.BadgesKeeper)
	suite.queryClient = queryClient

	bob_acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, sdk.MustAccAddressFromBech32(bob))
	alice_acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, sdk.MustAccAddressFromBech32(alice))
	charlie_acc := suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, sdk.MustAccAddressFromBech32(charlie))

	suite.app.AccountKeeper.SetAccount(suite.ctx, bob_acc)
	suite.app.AccountKeeper.SetAccount(suite.ctx, alice_acc)
	suite.app.AccountKeeper.SetAccount(suite.ctx, charlie_acc)

	//initialize bob with 1000 coins

	suite.ctx = suite.ctx.WithBlockTime(time.Now())

	// for i := uint64(0); i < 1000; i++ {
	// 	suite.app.AccountKeeper.SetAccount(suite.ctx, suite.app.AccountKeeper.NewAccountWithAddress(suite.ctx, sdk.AccAddress([]byte{byte(i)})))
	// }

	banktestutil.FundAccount(suite.app.BankKeeper, suite.ctx, sdk.MustAccAddressFromBech32(bob), sdk.NewCoins(sdk.NewInt64Coin("badge", 1000)))
}

func TestBadgesKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
