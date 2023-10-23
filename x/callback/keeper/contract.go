// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package keeper

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"value\",\"type\":\"uint32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decryptStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decryptView\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reveal\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"encryptedValue\",\"type\":\"bytes\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"testValue\",\"outputs\":[{\"internalType\":\"euint32\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DecryptView is a free data retrieval call binding the contract method 0x1dc0afe0.
//
// Solidity: function decryptView() view returns(uint32)
func (_Contract *ContractCaller) DecryptView(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "decryptView")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// DecryptView is a free data retrieval call binding the contract method 0x1dc0afe0.
//
// Solidity: function decryptView() view returns(uint32)
func (_Contract *ContractSession) DecryptView() (uint32, error) {
	return _Contract.Contract.DecryptView(&_Contract.CallOpts)
}

// DecryptView is a free data retrieval call binding the contract method 0x1dc0afe0.
//
// Solidity: function decryptView() view returns(uint32)
func (_Contract *ContractCallerSession) DecryptView() (uint32, error) {
	return _Contract.Contract.DecryptView(&_Contract.CallOpts)
}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() view returns(uint32)
func (_Contract *ContractCaller) Reveal(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "reveal")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() view returns(uint32)
func (_Contract *ContractSession) Reveal() (uint32, error) {
	return _Contract.Contract.Reveal(&_Contract.CallOpts)
}

// Reveal is a free data retrieval call binding the contract method 0xa475b5dd.
//
// Solidity: function reveal() view returns(uint32)
func (_Contract *ContractCallerSession) Reveal() (uint32, error) {
	return _Contract.Contract.Reveal(&_Contract.CallOpts)
}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_Contract *ContractCaller) TestValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "testValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_Contract *ContractSession) TestValue() (*big.Int, error) {
	return _Contract.Contract.TestValue(&_Contract.CallOpts)
}

// TestValue is a free data retrieval call binding the contract method 0x8af5de72.
//
// Solidity: function testValue() view returns(uint256)
func (_Contract *ContractCallerSession) TestValue() (*big.Int, error) {
	return _Contract.Contract.TestValue(&_Contract.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xcbe3a072.
//
// Solidity: function add(uint32 value) returns()
func (_Contract *ContractTransactor) Add(opts *bind.TransactOpts, value uint32) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "add", value)
}

// Add is a paid mutator transaction binding the contract method 0xcbe3a072.
//
// Solidity: function add(uint32 value) returns()
func (_Contract *ContractSession) Add(value uint32) (*types.Transaction, error) {
	return _Contract.Contract.Add(&_Contract.TransactOpts, value)
}

// Add is a paid mutator transaction binding the contract method 0xcbe3a072.
//
// Solidity: function add(uint32 value) returns()
func (_Contract *ContractTransactorSession) Add(value uint32) (*types.Transaction, error) {
	return _Contract.Contract.Add(&_Contract.TransactOpts, value)
}

// DecryptStore is a paid mutator transaction binding the contract method 0x2f59d16b.
//
// Solidity: function decryptStore() returns()
func (_Contract *ContractTransactor) DecryptStore(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "decryptStore")
}

// DecryptStore is a paid mutator transaction binding the contract method 0x2f59d16b.
//
// Solidity: function decryptStore() returns()
func (_Contract *ContractSession) DecryptStore() (*types.Transaction, error) {
	return _Contract.Contract.DecryptStore(&_Contract.TransactOpts)
}

// DecryptStore is a paid mutator transaction binding the contract method 0x2f59d16b.
//
// Solidity: function decryptStore() returns()
func (_Contract *ContractTransactorSession) DecryptStore() (*types.Transaction, error) {
	return _Contract.Contract.DecryptStore(&_Contract.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0xb374012b.
//
// Solidity: function store(bytes encryptedValue) returns()
func (_Contract *ContractTransactor) Store(opts *bind.TransactOpts, encryptedValue []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "store", encryptedValue)
}

// Store is a paid mutator transaction binding the contract method 0xb374012b.
//
// Solidity: function store(bytes encryptedValue) returns()
func (_Contract *ContractSession) Store(encryptedValue []byte) (*types.Transaction, error) {
	return _Contract.Contract.Store(&_Contract.TransactOpts, encryptedValue)
}

// Store is a paid mutator transaction binding the contract method 0xb374012b.
//
// Solidity: function store(bytes encryptedValue) returns()
func (_Contract *ContractTransactorSession) Store(encryptedValue []byte) (*types.Transaction, error) {
	return _Contract.Contract.Store(&_Contract.TransactOpts, encryptedValue)
}
