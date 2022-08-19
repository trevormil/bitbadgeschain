package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/trevormil/bitbadgeschain/x/badges/keeper"
	"github.com/trevormil/bitbadgeschain/x/badges/types"
)

func SimulateMsgTransferBadge(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		randomAccounts := []uint64{}
		for i := 0; i < r.Intn(10); i++ {
			randomAccounts = append(randomAccounts, r.Uint64())
		}

		randomAmounts := []uint64{}
		for i := 0; i < r.Intn(10); i++ {
			randomAmounts = append(randomAmounts, r.Uint64())
		}

		msg := &types.MsgTransferBadge{
			Creator:              simAccount.Address.String(),
			From:                 r.Uint64(),
			ToAddresses:          randomAccounts,
			Amounts:              randomAmounts,
			BadgeId:              r.Uint64(),
			ExpirationTime:       r.Uint64(),
			CantCancelBeforeTime: r.Uint64(),
			SubbadgeRanges: []*types.IdRange{
				{
					Start: r.Uint64(),
					End:   r.Uint64(),
				},
				{
					Start: r.Uint64(),
					End:   r.Uint64(),
				},
				{
					Start: r.Uint64(),
					End:   r.Uint64(),
				},
			},
		}

		return simtypes.NewOperationMsg(msg, true, "", types.ModuleCdc), nil, nil
	}
}
