package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/celestiaorg/dotcel/testutil/keeper"
	"github.com/celestiaorg/dotcel/x/cquery/keeper"
	"github.com/celestiaorg/dotcel/x/cquery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CqueryKeeper(t, nil)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
