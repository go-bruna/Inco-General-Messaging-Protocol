package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/callback/types"
)

func (k msgServer) CallEvmAdd(goCtx context.Context, msg *types.MsgCallEvmAdd) (*types.MsgCallEvmAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCallEvmAddResponse{}, nil
}
