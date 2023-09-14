package types

import (
	"github.com/ethereum/go-ethereum/common"
)

// GetContractAddress casts the hex string address of the ERC20 to common.Address
func GetContractAddress(contractAddr string) common.Address {
	return common.HexToAddress(contractAddr)
}
