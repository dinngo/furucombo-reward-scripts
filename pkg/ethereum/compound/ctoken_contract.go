// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compound

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

// CTokenContractABI is the input ABI used to generate the binding from.
const CTokenContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"}]"

// CTokenContract is an auto generated Go binding around an Ethereum contract.
type CTokenContract struct {
	CTokenContractCaller     // Read-only binding to the contract
	CTokenContractTransactor // Write-only binding to the contract
	CTokenContractFilterer   // Log filterer for contract events
}

// CTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type CTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CTokenContractSession struct {
	Contract     *CTokenContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CTokenContractCallerSession struct {
	Contract *CTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// CTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CTokenContractTransactorSession struct {
	Contract     *CTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type CTokenContractRaw struct {
	Contract *CTokenContract // Generic contract binding to access the raw methods on
}

// CTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CTokenContractCallerRaw struct {
	Contract *CTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// CTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CTokenContractTransactorRaw struct {
	Contract *CTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCTokenContract creates a new instance of CTokenContract, bound to a specific deployed contract.
func NewCTokenContract(address common.Address, backend bind.ContractBackend) (*CTokenContract, error) {
	contract, err := bindCTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CTokenContract{CTokenContractCaller: CTokenContractCaller{contract: contract}, CTokenContractTransactor: CTokenContractTransactor{contract: contract}, CTokenContractFilterer: CTokenContractFilterer{contract: contract}}, nil
}

// NewCTokenContractCaller creates a new read-only instance of CTokenContract, bound to a specific deployed contract.
func NewCTokenContractCaller(address common.Address, caller bind.ContractCaller) (*CTokenContractCaller, error) {
	contract, err := bindCTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CTokenContractCaller{contract: contract}, nil
}

// NewCTokenContractTransactor creates a new write-only instance of CTokenContract, bound to a specific deployed contract.
func NewCTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*CTokenContractTransactor, error) {
	contract, err := bindCTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CTokenContractTransactor{contract: contract}, nil
}

// NewCTokenContractFilterer creates a new log filterer instance of CTokenContract, bound to a specific deployed contract.
func NewCTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*CTokenContractFilterer, error) {
	contract, err := bindCTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CTokenContractFilterer{contract: contract}, nil
}

// bindCTokenContract binds a generic wrapper to an already deployed contract.
func bindCTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CTokenContract *CTokenContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CTokenContract.Contract.CTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CTokenContract *CTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTokenContract.Contract.CTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CTokenContract *CTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CTokenContract.Contract.CTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CTokenContract *CTokenContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CTokenContract *CTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CTokenContract *CTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CTokenContract.Contract.contract.Transact(opts, method, params...)
}

// CTokenContractBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the CTokenContract contract.
type CTokenContractBorrowIterator struct {
	Event *CTokenContractBorrow // Event containing the contract specifics and raw log

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
func (it *CTokenContractBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenContractBorrow)
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
		it.Event = new(CTokenContractBorrow)
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
func (it *CTokenContractBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenContractBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenContractBorrow represents a Borrow event raised by the CTokenContract contract.
type CTokenContractBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CTokenContract *CTokenContractFilterer) FilterBorrow(opts *bind.FilterOpts) (*CTokenContractBorrowIterator, error) {

	logs, sub, err := _CTokenContract.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &CTokenContractBorrowIterator{contract: _CTokenContract.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CTokenContract *CTokenContractFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *CTokenContractBorrow) (event.Subscription, error) {

	logs, sub, err := _CTokenContract.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenContractBorrow)
				if err := _CTokenContract.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CTokenContract *CTokenContractFilterer) ParseBorrow(log types.Log) (*CTokenContractBorrow, error) {
	event := new(CTokenContractBorrow)
	if err := _CTokenContract.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenContractMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CTokenContract contract.
type CTokenContractMintIterator struct {
	Event *CTokenContractMint // Event containing the contract specifics and raw log

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
func (it *CTokenContractMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenContractMint)
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
		it.Event = new(CTokenContractMint)
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
func (it *CTokenContractMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenContractMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenContractMint represents a Mint event raised by the CTokenContract contract.
type CTokenContractMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CTokenContract *CTokenContractFilterer) FilterMint(opts *bind.FilterOpts) (*CTokenContractMintIterator, error) {

	logs, sub, err := _CTokenContract.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &CTokenContractMintIterator{contract: _CTokenContract.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CTokenContract *CTokenContractFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CTokenContractMint) (event.Subscription, error) {

	logs, sub, err := _CTokenContract.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenContractMint)
				if err := _CTokenContract.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CTokenContract *CTokenContractFilterer) ParseMint(log types.Log) (*CTokenContractMint, error) {
	event := new(CTokenContractMint)
	if err := _CTokenContract.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenContractRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the CTokenContract contract.
type CTokenContractRepayBorrowIterator struct {
	Event *CTokenContractRepayBorrow // Event containing the contract specifics and raw log

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
func (it *CTokenContractRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenContractRepayBorrow)
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
		it.Event = new(CTokenContractRepayBorrow)
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
func (it *CTokenContractRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenContractRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenContractRepayBorrow represents a RepayBorrow event raised by the CTokenContract contract.
type CTokenContractRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CTokenContract *CTokenContractFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*CTokenContractRepayBorrowIterator, error) {

	logs, sub, err := _CTokenContract.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &CTokenContractRepayBorrowIterator{contract: _CTokenContract.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CTokenContract *CTokenContractFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *CTokenContractRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _CTokenContract.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenContractRepayBorrow)
				if err := _CTokenContract.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CTokenContract *CTokenContractFilterer) ParseRepayBorrow(log types.Log) (*CTokenContractRepayBorrow, error) {
	event := new(CTokenContractRepayBorrow)
	if err := _CTokenContract.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
