package keeper

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
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
	var data types.EncryptAddMessage
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		// NOTE: shouldn't happen as the packet has already
		// been decoded on ICS20 transfer logic
		err = errorsmod.Wrapf(errortypes.ErrInvalidType, "cannot unmarshal ICS-20 transfer packet data")
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	resp, err := k.CallEvmAdd(ctx, "osmosis smart contract", data.ContractAddress, data.FunctionName, data.Argument)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	arg64, err := strconv.ParseUint(data.Argument, 10, 64)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	count := k.GetTestDataCount(ctx)
	testData := types.TestData{
		Id:            count + 1,
		OriginValue:   arg64,
		ExecutedValue: resp.Result,
		Timestamp:     uint64(ctx.BlockTime().Unix()),
	}

	k.AppendTestData(ctx, testData)

	response := &types.EncryptAddIResponse{
		Value: (uint32)(resp.Result),
	}

	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(response),
	)
	// return the original success acknowledgement
	return ack
}
