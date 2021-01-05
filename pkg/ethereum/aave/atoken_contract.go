// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// ATokenContractABI is the input ABI used to generate the binding from.
const ATokenContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromIndex\",\"type\":\"uint256\"}],\"name\":\"MintOnDeposit\",\"type\":\"event\"}]"

// ATokenContract is an auto generated Go binding around an Ethereum contract.
type ATokenContract struct {
	ATokenContractCaller     // Read-only binding to the contract
	ATokenContractTransactor // Write-only binding to the contract
	ATokenContractFilterer   // Log filterer for contract events
}

// ATokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ATokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ATokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ATokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ATokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ATokenContractSession struct {
	Contract     *ATokenContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ATokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ATokenContractCallerSession struct {
	Contract *ATokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ATokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ATokenContractTransactorSession struct {
	Contract     *ATokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ATokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ATokenContractRaw struct {
	Contract *ATokenContract // Generic contract binding to access the raw methods on
}

// ATokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ATokenContractCallerRaw struct {
	Contract *ATokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// ATokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ATokenContractTransactorRaw struct {
	Contract *ATokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewATokenContract creates a new instance of ATokenContract, bound to a specific deployed contract.
func NewATokenContract(address common.Address, backend bind.ContractBackend) (*ATokenContract, error) {
	contract, err := bindATokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ATokenContract{ATokenContractCaller: ATokenContractCaller{contract: contract}, ATokenContractTransactor: ATokenContractTransactor{contract: contract}, ATokenContractFilterer: ATokenContractFilterer{contract: contract}}, nil
}

// NewATokenContractCaller creates a new read-only instance of ATokenContract, bound to a specific deployed contract.
func NewATokenContractCaller(address common.Address, caller bind.ContractCaller) (*ATokenContractCaller, error) {
	contract, err := bindATokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenContractCaller{contract: contract}, nil
}

// NewATokenContractTransactor creates a new write-only instance of ATokenContract, bound to a specific deployed contract.
func NewATokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ATokenContractTransactor, error) {
	contract, err := bindATokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ATokenContractTransactor{contract: contract}, nil
}

// NewATokenContractFilterer creates a new log filterer instance of ATokenContract, bound to a specific deployed contract.
func NewATokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ATokenContractFilterer, error) {
	contract, err := bindATokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ATokenContractFilterer{contract: contract}, nil
}

// bindATokenContract binds a generic wrapper to an already deployed contract.
func bindATokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ATokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ATokenContract *ATokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ATokenContract.Contract.ATokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ATokenContract *ATokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ATokenContract.Contract.ATokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ATokenContract *ATokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ATokenContract.Contract.ATokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ATokenContract *ATokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ATokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ATokenContract *ATokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ATokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ATokenContract *ATokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ATokenContract.Contract.contract.Transact(opts, method, params...)
}

// ATokenContractMintOnDepositIterator is returned from FilterMintOnDeposit and is used to iterate over the raw logs and unpacked data for MintOnDeposit events raised by the ATokenContract contract.
type ATokenContractMintOnDepositIterator struct {
	Event *ATokenContractMintOnDeposit // Event containing the contract specifics and raw log

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
func (it *ATokenContractMintOnDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ATokenContractMintOnDeposit)
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
		it.Event = new(ATokenContractMintOnDeposit)
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
func (it *ATokenContractMintOnDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ATokenContractMintOnDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ATokenContractMintOnDeposit represents a MintOnDeposit event raised by the ATokenContract contract.
type ATokenContractMintOnDeposit struct {
	From                common.Address
	Value               *big.Int
	FromBalanceIncrease *big.Int
	FromIndex           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMintOnDeposit is a free log retrieval operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenContract *ATokenContractFilterer) FilterMintOnDeposit(opts *bind.FilterOpts, _from []common.Address) (*ATokenContractMintOnDepositIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenContract.contract.FilterLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return &ATokenContractMintOnDepositIterator{contract: _ATokenContract.contract, event: "MintOnDeposit", logs: logs, sub: sub}, nil
}

// WatchMintOnDeposit is a free log subscription operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenContract *ATokenContractFilterer) WatchMintOnDeposit(opts *bind.WatchOpts, sink chan<- *ATokenContractMintOnDeposit, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _ATokenContract.contract.WatchLogs(opts, "MintOnDeposit", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ATokenContractMintOnDeposit)
				if err := _ATokenContract.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
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

// ParseMintOnDeposit is a log parse operation binding the contract event 0xbe7799898ca2d813ff902b487c1b434ab45b47edd8f38b77ad5e99aae8341b7a.
//
// Solidity: event MintOnDeposit(address indexed _from, uint256 _value, uint256 _fromBalanceIncrease, uint256 _fromIndex)
func (_ATokenContract *ATokenContractFilterer) ParseMintOnDeposit(log types.Log) (*ATokenContractMintOnDeposit, error) {
	event := new(ATokenContractMintOnDeposit)
	if err := _ATokenContract.contract.UnpackLog(event, "MintOnDeposit", log); err != nil {
		return nil, err
	}
	return event, nil
}
