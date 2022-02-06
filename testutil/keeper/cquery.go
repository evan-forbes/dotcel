package keeper

import (
	"context"
	"errors"
	"testing"

	"github.com/celestiaorg/dotcel/x/cquery/keeper"
	"github.com/celestiaorg/dotcel/x/cquery/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
	tmdb "github.com/tendermint/tm-db"
)

func CqueryKeeper(t testing.TB, mockedTxs map[string]*coretypes.ResultTx) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"CqueryParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		newMockTxGetter(mockedTxs),
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func newMockTxGetter(m map[string]*coretypes.ResultTx) types.TxGetter {
	var txg *mappedTxGetter
	if m == nil {
		txg = &mappedTxGetter{m: make(map[string]*coretypes.ResultTx)}
	} else {
		txg = &mappedTxGetter{m: m}
	}

	return txg.get
}

type mappedTxGetter struct {
	m map[string]*coretypes.ResultTx
}

func (txg *mappedTxGetter) get(_ context.Context, hash []byte) (*coretypes.ResultTx, error) {
	res, has := txg.m[tmbytes.HexBytes(hash).String()]
	if !has {
		return nil, errors.New("no transaction found using that hash")
	}
	return res, nil
}
