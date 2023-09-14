package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/callback/types"
)

func (k msgServer) CallEvmAdd(goCtx context.Context, msg *types.MsgCallEvmAdd) (*types.MsgCallEvmAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract := types.GetContractAddress(msg.ContractAddress)
	abi, err := ContractMetaData.GetAbi()
	if err != nil {
		return &types.MsgCallEvmAddResponse{}, nil
	}

	// Call add function
	resp, err := k.CallEVM(ctx, *abi, types.ModuleAddress, contract, true, msg.FuncName, msg.Arg)
	if err != nil {
		return nil, err
	}

	result, err := strconv.Atoi(string(resp.Ret))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeEncryptAdd,
				sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
				sdk.NewAttribute(types.EventTypeEncryptAddResult, fmt.Sprintf("%d", result)),
			),
		},
	)

	return &types.MsgCallEvmAddResponse{Result: (uint64)(result)}, nil
}
