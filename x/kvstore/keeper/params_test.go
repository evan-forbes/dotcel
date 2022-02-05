package keeper_test

import (
	"testing"

	testkeeper "github.com/celestiaorg/dotcel/testutil/keeper"
	"github.com/celestiaorg/dotcel/x/kvstore/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.KvstoreKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
