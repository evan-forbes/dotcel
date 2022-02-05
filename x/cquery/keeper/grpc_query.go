package keeper

import (
	"github.com/celestiaorg/dotcel/x/cquery/types"
)

var _ types.QueryServer = Keeper{}
