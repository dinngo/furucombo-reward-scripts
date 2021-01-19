// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package furucombo

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

// MerkleRedeemClaim is an auto generated low-level Go binding around an user-defined struct.
type MerkleRedeemClaim struct {
	Week        *big.Int
	Balance     *big.Int
	MerkleProof [][32]byte
}

// StakingContractABI is the input ABI used to generate the binding from.
const StakingContractABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"Approved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"week\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimWeek\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"week\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structMerkleRedeem.Claim[]\",\"name\":\"claims\",\"type\":\"tuple[]\"}],\"name\":\"claimWeeks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemable\",\"outputs\":[{\"internalType\":\"contractMerkleRedeem\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approval\",\"type\":\"bool\"}],\"name\":\"setApproval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstakeFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakingContract is an auto generated Go binding around an Ethereum contract.
type StakingContract struct {
	StakingContractCaller     // Read-only binding to the contract
	StakingContractTransactor // Write-only binding to the contract
	StakingContractFilterer   // Log filterer for contract events
}

// StakingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingContractSession struct {
	Contract     *StakingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingContractCallerSession struct {
	Contract *StakingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StakingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingContractTransactorSession struct {
	Contract     *StakingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingContractRaw struct {
	Contract *StakingContract // Generic contract binding to access the raw methods on
}

// StakingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingContractCallerRaw struct {
	Contract *StakingContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingContractTransactorRaw struct {
	Contract *StakingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingContract creates a new instance of StakingContract, bound to a specific deployed contract.
func NewStakingContract(address common.Address, backend bind.ContractBackend) (*StakingContract, error) {
	contract, err := bindStakingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingContract{StakingContractCaller: StakingContractCaller{contract: contract}, StakingContractTransactor: StakingContractTransactor{contract: contract}, StakingContractFilterer: StakingContractFilterer{contract: contract}}, nil
}

// NewStakingContractCaller creates a new read-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractCaller(address common.Address, caller bind.ContractCaller) (*StakingContractCaller, error) {
	contract, err := bindStakingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractCaller{contract: contract}, nil
}

// NewStakingContractTransactor creates a new write-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingContractTransactor, error) {
	contract, err := bindStakingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractTransactor{contract: contract}, nil
}

// NewStakingContractFilterer creates a new log filterer instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingContractFilterer, error) {
	contract, err := bindStakingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingContractFilterer{contract: contract}, nil
}

// bindStakingContract binds a generic wrapper to an already deployed contract.
func bindStakingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.StakingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_StakingContract *StakingContractCaller) BalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "balanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_StakingContract *StakingContractSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _StakingContract.Contract.BalanceOf(&_StakingContract.CallOpts, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address user) view returns(uint256)
func (_StakingContract *StakingContractCallerSession) BalanceOf(user common.Address) (*big.Int, error) {
	return _StakingContract.Contract.BalanceOf(&_StakingContract.CallOpts, user)
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(address user, address agent) view returns(bool)
func (_StakingContract *StakingContractCaller) IsApproved(opts *bind.CallOpts, user common.Address, agent common.Address) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "isApproved", user, agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(address user, address agent) view returns(bool)
func (_StakingContract *StakingContractSession) IsApproved(user common.Address, agent common.Address) (bool, error) {
	return _StakingContract.Contract.IsApproved(&_StakingContract.CallOpts, user, agent)
}

// IsApproved is a free data retrieval call binding the contract method 0xa389783e.
//
// Solidity: function isApproved(address user, address agent) view returns(bool)
func (_StakingContract *StakingContractCallerSession) IsApproved(user common.Address, agent common.Address) (bool, error) {
	return _StakingContract.Contract.IsApproved(&_StakingContract.CallOpts, user, agent)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingContract *StakingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingContract *StakingContractSession) Owner() (common.Address, error) {
	return _StakingContract.Contract.Owner(&_StakingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingContract *StakingContractCallerSession) Owner() (common.Address, error) {
	return _StakingContract.Contract.Owner(&_StakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractSession) Paused() (bool, error) {
	return _StakingContract.Contract.Paused(&_StakingContract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingContract *StakingContractCallerSession) Paused() (bool, error) {
	return _StakingContract.Contract.Paused(&_StakingContract.CallOpts)
}

// Redeemable is a free data retrieval call binding the contract method 0x2d7ecd11.
//
// Solidity: function redeemable() view returns(address)
func (_StakingContract *StakingContractCaller) Redeemable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "redeemable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Redeemable is a free data retrieval call binding the contract method 0x2d7ecd11.
//
// Solidity: function redeemable() view returns(address)
func (_StakingContract *StakingContractSession) Redeemable() (common.Address, error) {
	return _StakingContract.Contract.Redeemable(&_StakingContract.CallOpts)
}

// Redeemable is a free data retrieval call binding the contract method 0x2d7ecd11.
//
// Solidity: function redeemable() view returns(address)
func (_StakingContract *StakingContractCallerSession) Redeemable() (common.Address, error) {
	return _StakingContract.Contract.Redeemable(&_StakingContract.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakingContract *StakingContractCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakingContract *StakingContractSession) StakingToken() (common.Address, error) {
	return _StakingContract.Contract.StakingToken(&_StakingContract.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_StakingContract *StakingContractCallerSession) StakingToken() (common.Address, error) {
	return _StakingContract.Contract.StakingToken(&_StakingContract.CallOpts)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address user, uint256 week, uint256 balance, bytes32[] merkleProof) returns()
func (_StakingContract *StakingContractTransactor) ClaimWeek(opts *bind.TransactOpts, user common.Address, week *big.Int, balance *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "claimWeek", user, week, balance, merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address user, uint256 week, uint256 balance, bytes32[] merkleProof) returns()
func (_StakingContract *StakingContractSession) ClaimWeek(user common.Address, week *big.Int, balance *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ClaimWeek(&_StakingContract.TransactOpts, user, week, balance, merkleProof)
}

// ClaimWeek is a paid mutator transaction binding the contract method 0x58b4e4b4.
//
// Solidity: function claimWeek(address user, uint256 week, uint256 balance, bytes32[] merkleProof) returns()
func (_StakingContract *StakingContractTransactorSession) ClaimWeek(user common.Address, week *big.Int, balance *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _StakingContract.Contract.ClaimWeek(&_StakingContract.TransactOpts, user, week, balance, merkleProof)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xc804c39a.
//
// Solidity: function claimWeeks(address user, (uint256,uint256,bytes32[])[] claims) returns()
func (_StakingContract *StakingContractTransactor) ClaimWeeks(opts *bind.TransactOpts, user common.Address, claims []MerkleRedeemClaim) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "claimWeeks", user, claims)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xc804c39a.
//
// Solidity: function claimWeeks(address user, (uint256,uint256,bytes32[])[] claims) returns()
func (_StakingContract *StakingContractSession) ClaimWeeks(user common.Address, claims []MerkleRedeemClaim) (*types.Transaction, error) {
	return _StakingContract.Contract.ClaimWeeks(&_StakingContract.TransactOpts, user, claims)
}

// ClaimWeeks is a paid mutator transaction binding the contract method 0xc804c39a.
//
// Solidity: function claimWeeks(address user, (uint256,uint256,bytes32[])[] claims) returns()
func (_StakingContract *StakingContractTransactorSession) ClaimWeeks(user common.Address, claims []MerkleRedeemClaim) (*types.Transaction, error) {
	return _StakingContract.Contract.ClaimWeeks(&_StakingContract.TransactOpts, user, claims)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractSession) Pause() (*types.Transaction, error) {
	return _StakingContract.Contract.Pause(&_StakingContract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingContract *StakingContractTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingContract.Contract.Pause(&_StakingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingContract *StakingContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingContract *StakingContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingContract.Contract.RenounceOwnership(&_StakingContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingContract *StakingContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingContract.Contract.RenounceOwnership(&_StakingContract.TransactOpts)
}

// SetApproval is a paid mutator transaction binding the contract method 0xdb9b7170.
//
// Solidity: function setApproval(address agent, bool approval) returns()
func (_StakingContract *StakingContractTransactor) SetApproval(opts *bind.TransactOpts, agent common.Address, approval bool) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "setApproval", agent, approval)
}

// SetApproval is a paid mutator transaction binding the contract method 0xdb9b7170.
//
// Solidity: function setApproval(address agent, bool approval) returns()
func (_StakingContract *StakingContractSession) SetApproval(agent common.Address, approval bool) (*types.Transaction, error) {
	return _StakingContract.Contract.SetApproval(&_StakingContract.TransactOpts, agent, approval)
}

// SetApproval is a paid mutator transaction binding the contract method 0xdb9b7170.
//
// Solidity: function setApproval(address agent, bool approval) returns()
func (_StakingContract *StakingContractTransactorSession) SetApproval(agent common.Address, approval bool) (*types.Transaction, error) {
	return _StakingContract.Contract.SetApproval(&_StakingContract.TransactOpts, agent, approval)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakingContract *StakingContractTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakingContract *StakingContractSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Stake(&_StakingContract.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_StakingContract *StakingContractTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Stake(&_StakingContract.TransactOpts, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address onBehalfOf, uint256 amount) returns()
func (_StakingContract *StakingContractTransactor) StakeFor(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "stakeFor", onBehalfOf, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address onBehalfOf, uint256 amount) returns()
func (_StakingContract *StakingContractSession) StakeFor(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.StakeFor(&_StakingContract.TransactOpts, onBehalfOf, amount)
}

// StakeFor is a paid mutator transaction binding the contract method 0x2ee40908.
//
// Solidity: function stakeFor(address onBehalfOf, uint256 amount) returns()
func (_StakingContract *StakingContractTransactorSession) StakeFor(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.StakeFor(&_StakingContract.TransactOpts, onBehalfOf, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingContract *StakingContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingContract *StakingContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.TransferOwnership(&_StakingContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingContract *StakingContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingContract.Contract.TransferOwnership(&_StakingContract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractSession) Unpause() (*types.Transaction, error) {
	return _StakingContract.Contract.Unpause(&_StakingContract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingContract *StakingContractTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingContract.Contract.Unpause(&_StakingContract.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_StakingContract *StakingContractTransactor) Unstake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "unstake", amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_StakingContract *StakingContractSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Unstake(&_StakingContract.TransactOpts, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 amount) returns()
func (_StakingContract *StakingContractTransactorSession) Unstake(amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Unstake(&_StakingContract.TransactOpts, amount)
}

// UnstakeFor is a paid mutator transaction binding the contract method 0x36ef088c.
//
// Solidity: function unstakeFor(address onBehalfOf, uint256 amount) returns()
func (_StakingContract *StakingContractTransactor) UnstakeFor(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "unstakeFor", onBehalfOf, amount)
}

// UnstakeFor is a paid mutator transaction binding the contract method 0x36ef088c.
//
// Solidity: function unstakeFor(address onBehalfOf, uint256 amount) returns()
func (_StakingContract *StakingContractSession) UnstakeFor(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.UnstakeFor(&_StakingContract.TransactOpts, onBehalfOf, amount)
}

// UnstakeFor is a paid mutator transaction binding the contract method 0x36ef088c.
//
// Solidity: function unstakeFor(address onBehalfOf, uint256 amount) returns()
func (_StakingContract *StakingContractTransactorSession) UnstakeFor(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.UnstakeFor(&_StakingContract.TransactOpts, onBehalfOf, amount)
}

// StakingContractApprovedIterator is returned from FilterApproved and is used to iterate over the raw logs and unpacked data for Approved events raised by the StakingContract contract.
type StakingContractApprovedIterator struct {
	Event *StakingContractApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingContractApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingContractApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingContractApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractApproved represents a Approved event raised by the StakingContract contract.
type StakingContractApproved struct {
	User     common.Address
	Agent    common.Address
	Approval bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproved is a free log retrieval operation binding the contract event 0x1299be4ab334acdc5747f03ba508209acf422601c7ee3040dc8ff385b0aca989.
//
// Solidity: event Approved(address indexed user, address indexed agent, bool approval)
func (_StakingContract *StakingContractFilterer) FilterApproved(opts *bind.FilterOpts, user []common.Address, agent []common.Address) (*StakingContractApprovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Approved", userRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractApprovedIterator{contract: _StakingContract.contract, event: "Approved", logs: logs, sub: sub}, nil
}

// WatchApproved is a free log subscription operation binding the contract event 0x1299be4ab334acdc5747f03ba508209acf422601c7ee3040dc8ff385b0aca989.
//
// Solidity: event Approved(address indexed user, address indexed agent, bool approval)
func (_StakingContract *StakingContractFilterer) WatchApproved(opts *bind.WatchOpts, sink chan<- *StakingContractApproved, user []common.Address, agent []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Approved", userRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractApproved)
				if err := _StakingContract.contract.UnpackLog(event, "Approved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproved is a log parse operation binding the contract event 0x1299be4ab334acdc5747f03ba508209acf422601c7ee3040dc8ff385b0aca989.
//
// Solidity: event Approved(address indexed user, address indexed agent, bool approval)
func (_StakingContract *StakingContractFilterer) ParseApproved(log types.Log) (*StakingContractApproved, error) {
	event := new(StakingContractApproved)
	if err := _StakingContract.contract.UnpackLog(event, "Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingContract contract.
type StakingContractOwnershipTransferredIterator struct {
	Event *StakingContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractOwnershipTransferred represents a OwnershipTransferred event raised by the StakingContract contract.
type StakingContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingContract *StakingContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractOwnershipTransferredIterator{contract: _StakingContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingContract *StakingContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractOwnershipTransferred)
				if err := _StakingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingContract *StakingContractFilterer) ParseOwnershipTransferred(log types.Log) (*StakingContractOwnershipTransferred, error) {
	event := new(StakingContractOwnershipTransferred)
	if err := _StakingContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingContract contract.
type StakingContractPausedIterator struct {
	Event *StakingContractPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingContractPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingContractPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractPaused represents a Paused event raised by the StakingContract contract.
type StakingContractPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingContractPausedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingContractPausedIterator{contract: _StakingContract.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingContractPaused) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractPaused)
				if err := _StakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingContract *StakingContractFilterer) ParsePaused(log types.Log) (*StakingContractPaused, error) {
	event := new(StakingContractPaused)
	if err := _StakingContract.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingContract contract.
type StakingContractStakedIterator struct {
	Event *StakingContractStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingContractStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingContractStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingContractStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractStaked represents a Staked event raised by the StakingContract contract.
type StakingContractStaked struct {
	Sender     common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed sender, address indexed onBehalfOf, uint256 amount)
func (_StakingContract *StakingContractFilterer) FilterStaked(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address) (*StakingContractStakedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Staked", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractStakedIterator{contract: _StakingContract.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed sender, address indexed onBehalfOf, uint256 amount)
func (_StakingContract *StakingContractFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingContractStaked, sender []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Staked", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractStaked)
				if err := _StakingContract.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed sender, address indexed onBehalfOf, uint256 amount)
func (_StakingContract *StakingContractFilterer) ParseStaked(log types.Log) (*StakingContractStaked, error) {
	event := new(StakingContractStaked)
	if err := _StakingContract.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingContract contract.
type StakingContractUnpausedIterator struct {
	Event *StakingContractUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingContractUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingContractUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingContractUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractUnpaused represents a Unpaused event raised by the StakingContract contract.
type StakingContractUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingContractUnpausedIterator, error) {

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingContractUnpausedIterator{contract: _StakingContract.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingContractUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractUnpaused)
				if err := _StakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingContract *StakingContractFilterer) ParseUnpaused(log types.Log) (*StakingContractUnpaused, error) {
	event := new(StakingContractUnpaused)
	if err := _StakingContract.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingContractUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the StakingContract contract.
type StakingContractUnstakedIterator struct {
	Event *StakingContractUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingContractUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingContractUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingContractUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractUnstaked represents a Unstaked event raised by the StakingContract contract.
type StakingContractUnstaked struct {
	Sender     common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed sender, address indexed onBehalfOf, uint256 amount)
func (_StakingContract *StakingContractFilterer) FilterUnstaked(opts *bind.FilterOpts, sender []common.Address, onBehalfOf []common.Address) (*StakingContractUnstakedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Unstaked", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractUnstakedIterator{contract: _StakingContract.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed sender, address indexed onBehalfOf, uint256 amount)
func (_StakingContract *StakingContractFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingContractUnstaked, sender []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Unstaked", senderRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractUnstaked)
				if err := _StakingContract.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed sender, address indexed onBehalfOf, uint256 amount)
func (_StakingContract *StakingContractFilterer) ParseUnstaked(log types.Log) (*StakingContractUnstaked, error) {
	event := new(StakingContractUnstaked)
	if err := _StakingContract.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
