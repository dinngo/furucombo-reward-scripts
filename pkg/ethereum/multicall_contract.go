// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Target   common.Address
	CallData []byte
}

// MulticallContractABI is the input ABI used to generate the binding from.
const MulticallContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"name\":\"coinbase\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MulticallContract is an auto generated Go binding around an Ethereum contract.
type MulticallContract struct {
	MulticallContractCaller     // Read-only binding to the contract
	MulticallContractTransactor // Write-only binding to the contract
	MulticallContractFilterer   // Log filterer for contract events
}

// MulticallContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallContractSession struct {
	Contract     *MulticallContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MulticallContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallContractCallerSession struct {
	Contract *MulticallContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MulticallContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallContractTransactorSession struct {
	Contract     *MulticallContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MulticallContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallContractRaw struct {
	Contract *MulticallContract // Generic contract binding to access the raw methods on
}

// MulticallContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallContractCallerRaw struct {
	Contract *MulticallContractCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallContractTransactorRaw struct {
	Contract *MulticallContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallContract creates a new instance of MulticallContract, bound to a specific deployed contract.
func NewMulticallContract(address common.Address, backend bind.ContractBackend) (*MulticallContract, error) {
	contract, err := bindMulticallContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulticallContract{MulticallContractCaller: MulticallContractCaller{contract: contract}, MulticallContractTransactor: MulticallContractTransactor{contract: contract}, MulticallContractFilterer: MulticallContractFilterer{contract: contract}}, nil
}

// NewMulticallContractCaller creates a new read-only instance of MulticallContract, bound to a specific deployed contract.
func NewMulticallContractCaller(address common.Address, caller bind.ContractCaller) (*MulticallContractCaller, error) {
	contract, err := bindMulticallContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallContractCaller{contract: contract}, nil
}

// NewMulticallContractTransactor creates a new write-only instance of MulticallContract, bound to a specific deployed contract.
func NewMulticallContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallContractTransactor, error) {
	contract, err := bindMulticallContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallContractTransactor{contract: contract}, nil
}

// NewMulticallContractFilterer creates a new log filterer instance of MulticallContract, bound to a specific deployed contract.
func NewMulticallContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallContractFilterer, error) {
	contract, err := bindMulticallContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallContractFilterer{contract: contract}, nil
}

// bindMulticallContract binds a generic wrapper to an already deployed contract.
func bindMulticallContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MulticallContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallContract *MulticallContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallContract.Contract.MulticallContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallContract *MulticallContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallContract.Contract.MulticallContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallContract *MulticallContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallContract.Contract.MulticallContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallContract *MulticallContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallContract *MulticallContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallContract *MulticallContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallContract.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallContract *MulticallContractCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallContract *MulticallContractSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallContract.Contract.GetBlockHash(&_MulticallContract.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallContract *MulticallContractCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallContract.Contract.GetBlockHash(&_MulticallContract.CallOpts, blockNumber)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallContract *MulticallContractCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallContract *MulticallContractSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallContract.Contract.GetCurrentBlockCoinbase(&_MulticallContract.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallContract *MulticallContractCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallContract.Contract.GetCurrentBlockCoinbase(&_MulticallContract.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallContract *MulticallContractCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallContract *MulticallContractSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallContract.Contract.GetCurrentBlockDifficulty(&_MulticallContract.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallContract *MulticallContractCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallContract.Contract.GetCurrentBlockDifficulty(&_MulticallContract.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallContract *MulticallContractCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallContract *MulticallContractSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallContract.Contract.GetCurrentBlockGasLimit(&_MulticallContract.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallContract *MulticallContractCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallContract.Contract.GetCurrentBlockGasLimit(&_MulticallContract.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallContract *MulticallContractCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallContract *MulticallContractSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallContract.Contract.GetCurrentBlockTimestamp(&_MulticallContract.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallContract *MulticallContractCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallContract.Contract.GetCurrentBlockTimestamp(&_MulticallContract.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallContract *MulticallContractCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallContract *MulticallContractSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallContract.Contract.GetEthBalance(&_MulticallContract.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallContract *MulticallContractCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallContract.Contract.GetEthBalance(&_MulticallContract.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallContract *MulticallContractCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MulticallContract.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallContract *MulticallContractSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallContract.Contract.GetLastBlockHash(&_MulticallContract.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallContract *MulticallContractCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallContract.Contract.GetLastBlockHash(&_MulticallContract.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallContract *MulticallContractTransactor) Aggregate(opts *bind.TransactOpts, calls []Struct0) (*types.Transaction, error) {
	return _MulticallContract.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallContract *MulticallContractSession) Aggregate(calls []Struct0) (*types.Transaction, error) {
	return _MulticallContract.Contract.Aggregate(&_MulticallContract.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallContract *MulticallContractTransactorSession) Aggregate(calls []Struct0) (*types.Transaction, error) {
	return _MulticallContract.Contract.Aggregate(&_MulticallContract.TransactOpts, calls)
}
