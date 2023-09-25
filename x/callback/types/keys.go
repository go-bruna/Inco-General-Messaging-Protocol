package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName defines the module name
	ModuleName = "callback"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_callback"

	// Version defines the current version the IBC module supports
	Version = "pio-ibc-example-v1"

	// PortID is the default port id that module binds to
	PortID = "callback"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TestDataKey      = "TestData-value-"
	TestDataCountKey = "TestData-count-"
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}
