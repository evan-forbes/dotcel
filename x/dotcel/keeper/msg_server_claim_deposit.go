package keeper

import (
	"context"

	"github.com/celestiaorg/dotcel/x/dotcel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimDeposit(goCtx context.Context, msg *types.MsgClaimDeposit) (*types.MsgClaimDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimDepositResponse{}, nil
}
