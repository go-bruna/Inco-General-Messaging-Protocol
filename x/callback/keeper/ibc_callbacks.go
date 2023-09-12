package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/evmos/evmos/v9/x/callback/types"
)

func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
) error {
	return nil
}

// OnRecvPacket performs an IBC receive callback. Once a user receives an IBC
// transfer from a counterparty chain and the transfer is successful, the
// claimable amount for the receiver's claims record `ActionIBCTransfer` is
// claimed and transferred to the receivers address.
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack exported.Acknowledgement,
) exported.Acknowledgement {
	response := &types.WhoAmIResponse{
		Account: string(ctx.BlockHeader().ProposerAddress),
		BlockInfo: &types.BlockInfo{
			Height:  uint64(ctx.BlockHeight()),
			Time:    uint64(ctx.BlockTime().Unix()),
			ChainId: ctx.ChainID(),
		},
	}

	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(response),
	)
	// return the original success acknowledgement
	return ack
}
