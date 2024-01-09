package badges

import (
	"github.com/bitbadges/bitbadgeschain/x/badges/keeper"
	"github.com/bitbadges/bitbadgeschain/x/badges/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.

// NOTE: We assume that all badges are validly formed here
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined; default 0
	if genState.NextCollectionId.Equal(sdkmath.NewUint(0)) {
		genState.NextCollectionId = sdkmath.NewUint(1)
	}

	k.SetNextCollectionId(ctx, genState.NextCollectionId)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}

	for _, badge := range genState.Collections {
		if err := k.SetCollectionInStore(ctx, badge); err != nil {
			panic(err)
		}
	}

	for idx, balance := range genState.Balances {
		if err := k.SetUserBalanceInStore(ctx, genState.BalanceStoreKeys[idx], balance); err != nil {
			panic(err)
		}
	}

	for idx, numUsed := range genState.ChallengeTrackers {
		if err := k.SetChallengeTrackerInStore(ctx, genState.ChallengeTrackerStoreKeys[idx], numUsed); err != nil {
			panic(err)
		}
	}

	for _, addressList := range genState.AddressLists {
		if err := k.SetAddressListInStore(ctx, *addressList); err != nil {
			panic(err)
		}
	}

	for idx, approvalTracker := range genState.ApprovalTrackers {
		if err := k.SetApprovalTrackerInStoreViaKey(ctx, genState.ApprovalTrackerStoreKeys[idx], *approvalTracker); err != nil {
			panic(err)
		}
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.NextCollectionId = k.GetNextCollectionId(ctx)

	genesis.Collections = k.GetCollectionsFromStore(ctx)
	addresses := []string{}
	balanceIds := []sdkmath.Uint{}
	genesis.Balances, addresses, balanceIds = k.GetUserBalancesFromStore(ctx)

	for i, address := range addresses {
		genesis.BalanceStoreKeys = append(genesis.BalanceStoreKeys, keeper.ConstructBalanceKey(address, balanceIds[i]))
	}

	genesis.ChallengeTrackers, genesis.ChallengeTrackerStoreKeys = k.GetChallengeTrackersFromStore(ctx)

	genesis.AddressLists = k.GetAddressListsFromStore(ctx)

	genesis.ApprovalTrackers, genesis.ApprovalTrackerStoreKeys = k.GetApprovalTrackersFromStore(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
