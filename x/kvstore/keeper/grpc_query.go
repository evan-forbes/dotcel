package keeper

import (
	"github.com/celestiaorg/dotcel/x/kvstore/types"
)

var _ types.QueryServer = Keeper{}
