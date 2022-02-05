package keeper

import (
	"context"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

func (k Keeper) DepositByHash(ctx context.Context, hash []byte) (*coretypes.ResultTx, error) {
	return k.client.Tx(ctx, hash, false)
}
