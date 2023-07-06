package tokenfactory

import (
	"bita/x/tokenfactory/keeper"
	"bita/x/tokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the denom
	for _, elem := range genState.DenomList {
		k.SetDenom(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DenomList = k.GetAllDenom(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
