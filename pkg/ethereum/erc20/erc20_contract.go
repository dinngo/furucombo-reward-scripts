// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20ContractABI is the input ABI used to generate the binding from.
const ERC20ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC20Contract is an auto generated Go binding around an Ethereum contract.
type ERC20Contract struct {
	ERC20ContractCaller     // Read-only binding to the contract
	ERC20ContractTransactor // Write-only binding to the contract
	ERC20ContractFilterer   // Log filterer for contract events
}

// ERC20ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20ContractSession struct {
	Contract     *ERC20Contract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20ContractCallerSession struct {
	Contract *ERC20ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20ContractTransactorSession struct {
	Contract     *ERC20ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20ContractRaw struct {
	Contract *ERC20Contract // Generic contract binding to access the raw methods on
}

// ERC20ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20ContractCallerRaw struct {
	Contract *ERC20ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20ContractTransactorRaw struct {
	Contract *ERC20ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Contract creates a new instance of ERC20Contract, bound to a specific deployed contract.
func NewERC20Contract(address common.Address, backend bind.ContractBackend) (*ERC20Contract, error) {
	contract, err := bindERC20Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Contract{ERC20ContractCaller: ERC20ContractCaller{contract: contract}, ERC20ContractTransactor: ERC20ContractTransactor{contract: contract}, ERC20ContractFilterer: ERC20ContractFilterer{contract: contract}}, nil
}

// NewERC20ContractCaller creates a new read-only instance of ERC20Contract, bound to a specific deployed contract.
func NewERC20ContractCaller(address common.Address, caller bind.ContractCaller) (*ERC20ContractCaller, error) {
	contract, err := bindERC20Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ContractCaller{contract: contract}, nil
}

// NewERC20ContractTransactor creates a new write-only instance of ERC20Contract, bound to a specific deployed contract.
func NewERC20ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20ContractTransactor, error) {
	contract, err := bindERC20Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20ContractTransactor{contract: contract}, nil
}

// NewERC20ContractFilterer creates a new log filterer instance of ERC20Contract, bound to a specific deployed contract.
func NewERC20ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20ContractFilterer, error) {
	contract, err := bindERC20Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20ContractFilterer{contract: contract}, nil
}

// bindERC20Contract binds a generic wrapper to an already deployed contract.
func bindERC20Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Contract *ERC20ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Contract.Contract.ERC20ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Contract *ERC20ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Contract.Contract.ERC20ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Contract *ERC20ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Contract.Contract.ERC20ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Contract *ERC20ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Contract *ERC20ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Contract *ERC20ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Contract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ERC20Contract *ERC20ContractCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Contract.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ERC20Contract *ERC20ContractSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20Contract.Contract.BalanceOf(&_ERC20Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ERC20Contract *ERC20ContractCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ERC20Contract.Contract.BalanceOf(&_ERC20Contract.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Contract *ERC20ContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ERC20Contract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Contract *ERC20ContractSession) Decimals() (uint8, error) {
	return _ERC20Contract.Contract.Decimals(&_ERC20Contract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Contract *ERC20ContractCallerSession) Decimals() (uint8, error) {
	return _ERC20Contract.Contract.Decimals(&_ERC20Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_ERC20Contract *ERC20ContractCaller) Symbol(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ERC20Contract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_ERC20Contract *ERC20ContractSession) Symbol() ([]byte, error) {
	return _ERC20Contract.Contract.Symbol(&_ERC20Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes)
func (_ERC20Contract *ERC20ContractCallerSession) Symbol() ([]byte, error) {
	return _ERC20Contract.Contract.Symbol(&_ERC20Contract.CallOpts)
}
