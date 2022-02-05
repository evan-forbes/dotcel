package kvstore

import (
	"github.com/celestiaorg/dotcel/x/kvstore/keeper"
	"github.com/celestiaorg/dotcel/x/kvstore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the domain
	for _, elem := range genState.DomainList {
		k.SetDomain(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DomainList = k.GetAllDomain(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
