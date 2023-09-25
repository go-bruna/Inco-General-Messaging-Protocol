package callback

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/evmos/evmos/v9/x/callback/keeper"
	"github.com/evmos/evmos/v9/x/callback/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, accountKeeper authkeeper.AccountKeeper, genState types.GenesisState) {
	// Set all the testData
	for _, elem := range genState.TestDataList {
		k.SetTestData(ctx, elem)
	}

	// Set testData count
	k.SetTestDataCount(ctx, genState.TestDataCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	k.InitGenesis(ctx, genState)

	// ensure callback module account is set on genesis
	if acc := accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		// NOTE: shouldn't occur
		panic("the callback module account has not been set")
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.TestDataList = k.GetAllTestData(ctx)
	genesis.TestDataCount = k.GetTestDataCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
