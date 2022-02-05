package cquery_test

import (
	"testing"

	keepertest "github.com/celestiaorg/dotcel/testutil/keeper"
	"github.com/celestiaorg/dotcel/testutil/nullify"
	"github.com/celestiaorg/dotcel/x/cquery"
	"github.com/celestiaorg/dotcel/x/cquery/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CqueryKeeper(t)
	cquery.InitGenesis(ctx, *k, genesisState)
	got := cquery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
