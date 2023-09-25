package keeper

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"

	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
	"github.com/evmos/evmos/v9/x/callback/types"
)

type (
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		paramstore    paramtypes.Subspace
		ics4Wrapper   porttypes.ICS4Wrapper
		scopedKeeper  capabilitykeeper.ScopedKeeper
		accountKeeper types.AccountKeeper

		channelKeeper types.ChannelKeeper
		portKeeper    types.PortKeeper
		evmKeeper     types.EVMKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	evmKeeper types.EVMKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		paramstore:    ps,
		accountKeeper: ak,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		evmKeeper:     evmKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetICS4Wrapper sets the ICS4 wrapper to the keeper.
// It panics if already set
func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	if k.ics4Wrapper != nil {
		panic("ICS4 wrapper already set")
	}

	k.ics4Wrapper = ics4Wrapper
}

// ----------------------------------------------------------------------------
// IBC Keeper Logic
// ----------------------------------------------------------------------------

// ChanCloseInit defines a wrapper function for the channel Keeper's function.
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IBC callbacks and transfer handlers
// SendPacket implements the ICS4Wrapper interface from the transfer module.
// It calls the underlying SendPacket function directly to move down the middleware stack.
func (k Keeper) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet exported.PacketI) error {
	return k.ics4Wrapper.SendPacket(ctx, channelCap, packet)
}

// WriteAcknowledgement implements the ICS4Wrapper interface from the transfer module.
// It calls the underlying WriteAcknowledgement function directly to move down the middleware stack.
func (k Keeper) WriteAcknowledgement(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet exported.PacketI, ack exported.Acknowledgement) error {
	return k.ics4Wrapper.WriteAcknowledgement(ctx, channelCap, packet, ack)
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the IBC app module to claim a capability that core IBC
// passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// IsBound checks if the IBC app module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the port Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// InitGenesis initializes the ibc-module state and binds to PortID.
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, types.PortID) {
		// and claims the returned capability
		err := k.BindPort(ctx, types.PortID)
		if err != nil {
			panic(fmt.Sprintf("could not claim port capability: %v", err))
		}
	}
}

func (k Keeper) CallEvmAdd(ctx sdk.Context, creator string, contractAddr string, funcName string, arg string) (*types.MsgCallEvmAddResponse, error) {
	contract := types.GetContractAddress(contractAddr)
	abi, err := ContractMetaData.GetAbi()
	if err != nil {
		return &types.MsgCallEvmAddResponse{}, err
	}

	arg64, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		return &types.MsgCallEvmAddResponse{}, err
	}

	arg32 := uint32(arg64)
	// Call add function
	resp, err := k.CallEVM(ctx, *abi, types.ModuleAddress, contract, true, funcName, arg32)
	if err != nil || len(resp.Ret) == 0 {
		return &types.MsgCallEvmAddResponse{}, err
	}

	result, err := strconv.Atoi(string(resp.Ret))
	if err != nil {
		return &types.MsgCallEvmAddResponse{}, err
	}

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				types.EventTypeEncryptAdd,
				sdk.NewAttribute(sdk.AttributeKeySender, creator),
				sdk.NewAttribute(types.EventTypeEncryptAddResult, fmt.Sprintf("%d", result)),
			),
		},
	)

	return &types.MsgCallEvmAddResponse{Result: (uint64)(result)}, nil
}
