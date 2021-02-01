// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavev2

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

// LendingPoolContractABI is the input ABI used to generate the binding from.
const LendingPoolContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Deposit\",\"type\":\"event\"}]"

// LendingPoolContract is an auto generated Go binding around an Ethereum contract.
type LendingPoolContract struct {
	LendingPoolContractCaller     // Read-only binding to the contract
	LendingPoolContractTransactor // Write-only binding to the contract
	LendingPoolContractFilterer   // Log filterer for contract events
}

// LendingPoolContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingPoolContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingPoolContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingPoolContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingPoolContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingPoolContractSession struct {
	Contract     *LendingPoolContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LendingPoolContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingPoolContractCallerSession struct {
	Contract *LendingPoolContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LendingPoolContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingPoolContractTransactorSession struct {
	Contract     *LendingPoolContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LendingPoolContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingPoolContractRaw struct {
	Contract *LendingPoolContract // Generic contract binding to access the raw methods on
}

// LendingPoolContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingPoolContractCallerRaw struct {
	Contract *LendingPoolContractCaller // Generic read-only contract binding to access the raw methods on
}

// LendingPoolContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingPoolContractTransactorRaw struct {
	Contract *LendingPoolContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingPoolContract creates a new instance of LendingPoolContract, bound to a specific deployed contract.
func NewLendingPoolContract(address common.Address, backend bind.ContractBackend) (*LendingPoolContract, error) {
	contract, err := bindLendingPoolContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendingPoolContract{LendingPoolContractCaller: LendingPoolContractCaller{contract: contract}, LendingPoolContractTransactor: LendingPoolContractTransactor{contract: contract}, LendingPoolContractFilterer: LendingPoolContractFilterer{contract: contract}}, nil
}

// NewLendingPoolContractCaller creates a new read-only instance of LendingPoolContract, bound to a specific deployed contract.
func NewLendingPoolContractCaller(address common.Address, caller bind.ContractCaller) (*LendingPoolContractCaller, error) {
	contract, err := bindLendingPoolContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolContractCaller{contract: contract}, nil
}

// NewLendingPoolContractTransactor creates a new write-only instance of LendingPoolContract, bound to a specific deployed contract.
func NewLendingPoolContractTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingPoolContractTransactor, error) {
	contract, err := bindLendingPoolContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingPoolContractTransactor{contract: contract}, nil
}

// NewLendingPoolContractFilterer creates a new log filterer instance of LendingPoolContract, bound to a specific deployed contract.
func NewLendingPoolContractFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingPoolContractFilterer, error) {
	contract, err := bindLendingPoolContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingPoolContractFilterer{contract: contract}, nil
}

// bindLendingPoolContract binds a generic wrapper to an already deployed contract.
func bindLendingPoolContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingPoolContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolContract *LendingPoolContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPoolContract.Contract.LendingPoolContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolContract *LendingPoolContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolContract.Contract.LendingPoolContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolContract *LendingPoolContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolContract.Contract.LendingPoolContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendingPoolContract *LendingPoolContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendingPoolContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendingPoolContract *LendingPoolContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendingPoolContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendingPoolContract *LendingPoolContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendingPoolContract.Contract.contract.Transact(opts, method, params...)
}

// LendingPoolContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LendingPoolContract contract.
type LendingPoolContractDepositIterator struct {
	Event *LendingPoolContractDeposit // Event containing the contract specifics and raw log

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
func (it *LendingPoolContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingPoolContractDeposit)
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
		it.Event = new(LendingPoolContractDeposit)
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
func (it *LendingPoolContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingPoolContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingPoolContractDeposit represents a Deposit event raised by the LendingPoolContract contract.
type LendingPoolContractDeposit struct {
	Reserve    common.Address
	User       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_LendingPoolContract *LendingPoolContractFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*LendingPoolContractDepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _LendingPoolContract.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &LendingPoolContractDepositIterator{contract: _LendingPoolContract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_LendingPoolContract *LendingPoolContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LendingPoolContractDeposit, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _LendingPoolContract.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingPoolContractDeposit)
				if err := _LendingPoolContract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_LendingPoolContract *LendingPoolContractFilterer) ParseDeposit(log types.Log) (*LendingPoolContractDeposit, error) {
	event := new(LendingPoolContractDeposit)
	if err := _LendingPoolContract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
