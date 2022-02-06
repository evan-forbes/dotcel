package types

import (
	context "context"

	"github.com/tendermint/tendermint/rpc/client/http"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

type TxGetter func(ctx context.Context, hash []byte) (*coretypes.ResultTx, error)

type HttpTxGetter struct {
	Client *http.HTTP
}

func (txg *HttpTxGetter) Get(ctx context.Context, hash []byte) (*coretypes.ResultTx, error) {
	return txg.Client.Tx(ctx, hash, false)
}
