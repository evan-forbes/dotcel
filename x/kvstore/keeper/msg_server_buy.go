package keeper

import (
	"context"

	"github.com/celestiaorg/dotcel/x/kvstore/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyResponse{}, nil
}
