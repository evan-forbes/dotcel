package kvstore_test

import (
	"testing"

	keepertest "github.com/celestiaorg/dotcel/testutil/keeper"
	"github.com/celestiaorg/dotcel/testutil/nullify"
	"github.com/celestiaorg/dotcel/x/kvstore"
	"github.com/celestiaorg/dotcel/x/kvstore/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DomainList: []types.Domain{
			{
				Name: "0",
			},
			{
				Name: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KvstoreKeeper(t)
	kvstore.InitGenesis(ctx, *k, genesisState)
	got := kvstore.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DomainList, got.DomainList)
	// this line is used by starport scaffolding # genesis/test/assert
}
