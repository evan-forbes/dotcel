package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/celestiaorg/dotcel/testutil/keeper"
	"github.com/celestiaorg/dotcel/testutil/nullify"
	"github.com/celestiaorg/dotcel/x/kvstore/keeper"
	"github.com/celestiaorg/dotcel/x/kvstore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDomain(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Domain {
	items := make([]types.Domain, n)
	for i := range items {
		items[i].Name = strconv.Itoa(i)

		keeper.SetDomain(ctx, items[i])
	}
	return items
}

func TestDomainGet(t *testing.T) {
	keeper, ctx := keepertest.KvstoreKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDomain(ctx,
			item.Name,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDomainRemove(t *testing.T) {
	keeper, ctx := keepertest.KvstoreKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDomain(ctx,
			item.Name,
		)
		_, found := keeper.GetDomain(ctx,
			item.Name,
		)
		require.False(t, found)
	}
}

func TestDomainGetAll(t *testing.T) {
	keeper, ctx := keepertest.KvstoreKeeper(t)
	items := createNDomain(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDomain(ctx)),
	)
}
