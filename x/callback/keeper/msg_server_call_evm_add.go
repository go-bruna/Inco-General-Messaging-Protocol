package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/callback/types"
)

func (k msgServer) CallEvmAdd(goCtx context.Context, msg *types.MsgCallEvmAdd) (*types.MsgCallEvmAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return k.Keeper.CallEvmAdd(ctx, msg.Creator, msg.ContractAddress, msg.FuncName, msg.Arg)
}
