// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2

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

// PairContractABI is the input ABI used to generate the binding from.
const PairContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// PairContract is an auto generated Go binding around an Ethereum contract.
type PairContract struct {
	PairContractCaller     // Read-only binding to the contract
	PairContractTransactor // Write-only binding to the contract
	PairContractFilterer   // Log filterer for contract events
}

// PairContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type PairContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PairContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PairContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PairContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PairContractSession struct {
	Contract     *PairContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PairContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PairContractCallerSession struct {
	Contract *PairContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PairContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PairContractTransactorSession struct {
	Contract     *PairContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PairContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type PairContractRaw struct {
	Contract *PairContract // Generic contract binding to access the raw methods on
}

// PairContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PairContractCallerRaw struct {
	Contract *PairContractCaller // Generic read-only contract binding to access the raw methods on
}

// PairContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PairContractTransactorRaw struct {
	Contract *PairContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPairContract creates a new instance of PairContract, bound to a specific deployed contract.
func NewPairContract(address common.Address, backend bind.ContractBackend) (*PairContract, error) {
	contract, err := bindPairContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PairContract{PairContractCaller: PairContractCaller{contract: contract}, PairContractTransactor: PairContractTransactor{contract: contract}, PairContractFilterer: PairContractFilterer{contract: contract}}, nil
}

// NewPairContractCaller creates a new read-only instance of PairContract, bound to a specific deployed contract.
func NewPairContractCaller(address common.Address, caller bind.ContractCaller) (*PairContractCaller, error) {
	contract, err := bindPairContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PairContractCaller{contract: contract}, nil
}

// NewPairContractTransactor creates a new write-only instance of PairContract, bound to a specific deployed contract.
func NewPairContractTransactor(address common.Address, transactor bind.ContractTransactor) (*PairContractTransactor, error) {
	contract, err := bindPairContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PairContractTransactor{contract: contract}, nil
}

// NewPairContractFilterer creates a new log filterer instance of PairContract, bound to a specific deployed contract.
func NewPairContractFilterer(address common.Address, filterer bind.ContractFilterer) (*PairContractFilterer, error) {
	contract, err := bindPairContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PairContractFilterer{contract: contract}, nil
}

// bindPairContract binds a generic wrapper to an already deployed contract.
func bindPairContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PairContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairContract *PairContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairContract.Contract.PairContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairContract *PairContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairContract.Contract.PairContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairContract *PairContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairContract.Contract.PairContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PairContract *PairContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PairContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PairContract *PairContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PairContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PairContract *PairContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PairContract.Contract.contract.Transact(opts, method, params...)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairContract *PairContractCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PairContract.contract.Call(opts, out, "token0")
	return *ret0, err
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairContract *PairContractSession) Token0() (common.Address, error) {
	return _PairContract.Contract.Token0(&_PairContract.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_PairContract *PairContractCallerSession) Token0() (common.Address, error) {
	return _PairContract.Contract.Token0(&_PairContract.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairContract *PairContractCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PairContract.contract.Call(opts, out, "token1")
	return *ret0, err
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairContract *PairContractSession) Token1() (common.Address, error) {
	return _PairContract.Contract.Token1(&_PairContract.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_PairContract *PairContractCallerSession) Token1() (common.Address, error) {
	return _PairContract.Contract.Token1(&_PairContract.CallOpts)
}

// PairContractSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the PairContract contract.
type PairContractSwapIterator struct {
	Event *PairContractSwap // Event containing the contract specifics and raw log

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
func (it *PairContractSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PairContractSwap)
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
		it.Event = new(PairContractSwap)
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
func (it *PairContractSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PairContractSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PairContractSwap represents a Swap event raised by the PairContract contract.
type PairContractSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_PairContract *PairContractFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PairContractSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PairContract.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PairContractSwapIterator{contract: _PairContract.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_PairContract *PairContractFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PairContractSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PairContract.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PairContractSwap)
				if err := _PairContract.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_PairContract *PairContractFilterer) ParseSwap(log types.Log) (*PairContractSwap, error) {
	event := new(PairContractSwap)
	if err := _PairContract.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}
