package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/evmos/v9/x/callback/types"
)

func (k msgServer) DeployContract(goCtx context.Context, msg *types.MsgDeployContract) (*types.MsgDeployContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := k.DeployIncoContract(ctx)

	if err != nil {
		return &types.MsgDeployContractResponse{Address: ""}, err
	}
	return &types.MsgDeployContractResponse{Address: address.String()}, nil
}
