// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721

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

// ERC721ContractABI is the input ABI used to generate the binding from.
const ERC721ContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"_name\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"_approved\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementsERC721\",\"outputs\":[{\"name\":\"_implementsERC721\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"_totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"_infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"_balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_approvedAddress\",\"type\":\"address\"},{\"name\":\"_metadata\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"_symbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numTokensTotal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getOwnerTokens\",\"outputs\":[{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC721Contract is an auto generated Go binding around an Ethereum contract.
type ERC721Contract struct {
	ERC721ContractCaller     // Read-only binding to the contract
	ERC721ContractTransactor // Write-only binding to the contract
	ERC721ContractFilterer   // Log filterer for contract events
}

// ERC721ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721ContractSession struct {
	Contract     *ERC721Contract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721ContractCallerSession struct {
	Contract *ERC721ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721ContractTransactorSession struct {
	Contract     *ERC721ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721ContractRaw struct {
	Contract *ERC721Contract // Generic contract binding to access the raw methods on
}

// ERC721ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721ContractCallerRaw struct {
	Contract *ERC721ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721ContractTransactorRaw struct {
	Contract *ERC721ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Contract creates a new instance of ERC721Contract, bound to a specific deployed contract.
func NewERC721Contract(address common.Address, backend bind.ContractBackend) (*ERC721Contract, error) {
	contract, err := bindERC721Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Contract{ERC721ContractCaller: ERC721ContractCaller{contract: contract}, ERC721ContractTransactor: ERC721ContractTransactor{contract: contract}, ERC721ContractFilterer: ERC721ContractFilterer{contract: contract}}, nil
}

// NewERC721ContractCaller creates a new read-only instance of ERC721Contract, bound to a specific deployed contract.
func NewERC721ContractCaller(address common.Address, caller bind.ContractCaller) (*ERC721ContractCaller, error) {
	contract, err := bindERC721Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractCaller{contract: contract}, nil
}

// NewERC721ContractTransactor creates a new write-only instance of ERC721Contract, bound to a specific deployed contract.
func NewERC721ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721ContractTransactor, error) {
	contract, err := bindERC721Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractTransactor{contract: contract}, nil
}

// NewERC721ContractFilterer creates a new log filterer instance of ERC721Contract, bound to a specific deployed contract.
func NewERC721ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721ContractFilterer, error) {
	contract, err := bindERC721Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractFilterer{contract: contract}, nil
}

// bindERC721Contract binds a generic wrapper to an already deployed contract.
func bindERC721Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Contract *ERC721ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Contract.Contract.ERC721ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Contract *ERC721ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Contract.Contract.ERC721ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Contract *ERC721ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Contract.Contract.ERC721ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Contract *ERC721ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Contract *ERC721ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Contract *ERC721ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Contract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_ERC721Contract *ERC721ContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_ERC721Contract *ERC721ContractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Contract.Contract.BalanceOf(&_ERC721Contract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 _balance)
func (_ERC721Contract *ERC721ContractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721Contract.Contract.BalanceOf(&_ERC721Contract.CallOpts, _owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _approved)
func (_ERC721Contract *ERC721ContractCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "getApproved", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _approved)
func (_ERC721Contract *ERC721ContractSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Contract.Contract.GetApproved(&_ERC721Contract.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _tokenId) view returns(address _approved)
func (_ERC721Contract *ERC721ContractCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Contract.Contract.GetApproved(&_ERC721Contract.CallOpts, _tokenId)
}

// GetOwnerTokens is a free data retrieval call binding the contract method 0xd63d4af0.
//
// Solidity: function getOwnerTokens(address _owner) view returns(uint256[] _tokenIds)
func (_ERC721Contract *ERC721ContractCaller) GetOwnerTokens(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "getOwnerTokens", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOwnerTokens is a free data retrieval call binding the contract method 0xd63d4af0.
//
// Solidity: function getOwnerTokens(address _owner) view returns(uint256[] _tokenIds)
func (_ERC721Contract *ERC721ContractSession) GetOwnerTokens(_owner common.Address) ([]*big.Int, error) {
	return _ERC721Contract.Contract.GetOwnerTokens(&_ERC721Contract.CallOpts, _owner)
}

// GetOwnerTokens is a free data retrieval call binding the contract method 0xd63d4af0.
//
// Solidity: function getOwnerTokens(address _owner) view returns(uint256[] _tokenIds)
func (_ERC721Contract *ERC721ContractCallerSession) GetOwnerTokens(_owner common.Address) ([]*big.Int, error) {
	return _ERC721Contract.Contract.GetOwnerTokens(&_ERC721Contract.CallOpts, _owner)
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() view returns(bool _implementsERC721)
func (_ERC721Contract *ERC721ContractCaller) ImplementsERC721(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "implementsERC721")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() view returns(bool _implementsERC721)
func (_ERC721Contract *ERC721ContractSession) ImplementsERC721() (bool, error) {
	return _ERC721Contract.Contract.ImplementsERC721(&_ERC721Contract.CallOpts)
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() view returns(bool _implementsERC721)
func (_ERC721Contract *ERC721ContractCallerSession) ImplementsERC721() (bool, error) {
	return _ERC721Contract.Contract.ImplementsERC721(&_ERC721Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string _name)
func (_ERC721Contract *ERC721ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string _name)
func (_ERC721Contract *ERC721ContractSession) Name() (string, error) {
	return _ERC721Contract.Contract.Name(&_ERC721Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string _name)
func (_ERC721Contract *ERC721ContractCallerSession) Name() (string, error) {
	return _ERC721Contract.Contract.Name(&_ERC721Contract.CallOpts)
}

// NumTokensTotal is a free data retrieval call binding the contract method 0xaf129dc2.
//
// Solidity: function numTokensTotal() view returns(uint256)
func (_ERC721Contract *ERC721ContractCaller) NumTokensTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "numTokensTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokensTotal is a free data retrieval call binding the contract method 0xaf129dc2.
//
// Solidity: function numTokensTotal() view returns(uint256)
func (_ERC721Contract *ERC721ContractSession) NumTokensTotal() (*big.Int, error) {
	return _ERC721Contract.Contract.NumTokensTotal(&_ERC721Contract.CallOpts)
}

// NumTokensTotal is a free data retrieval call binding the contract method 0xaf129dc2.
//
// Solidity: function numTokensTotal() view returns(uint256)
func (_ERC721Contract *ERC721ContractCallerSession) NumTokensTotal() (*big.Int, error) {
	return _ERC721Contract.Contract.NumTokensTotal(&_ERC721Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_ERC721Contract *ERC721ContractCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_ERC721Contract *ERC721ContractSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Contract.Contract.OwnerOf(&_ERC721Contract.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address _owner)
func (_ERC721Contract *ERC721ContractCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721Contract.Contract.OwnerOf(&_ERC721Contract.CallOpts, _tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string _symbol)
func (_ERC721Contract *ERC721ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string _symbol)
func (_ERC721Contract *ERC721ContractSession) Symbol() (string, error) {
	return _ERC721Contract.Contract.Symbol(&_ERC721Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string _symbol)
func (_ERC721Contract *ERC721ContractCallerSession) Symbol() (string, error) {
	return _ERC721Contract.Contract.Symbol(&_ERC721Contract.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 _tokenId) view returns(string _infoUrl)
func (_ERC721Contract *ERC721ContractCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "tokenMetadata", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 _tokenId) view returns(string _infoUrl)
func (_ERC721Contract *ERC721ContractSession) TokenMetadata(_tokenId *big.Int) (string, error) {
	return _ERC721Contract.Contract.TokenMetadata(&_ERC721Contract.CallOpts, _tokenId)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x6914db60.
//
// Solidity: function tokenMetadata(uint256 _tokenId) view returns(string _infoUrl)
func (_ERC721Contract *ERC721ContractCallerSession) TokenMetadata(_tokenId *big.Int) (string, error) {
	return _ERC721Contract.Contract.TokenMetadata(&_ERC721Contract.CallOpts, _tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_ERC721Contract *ERC721ContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_ERC721Contract *ERC721ContractSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Contract.Contract.TokenOfOwnerByIndex(&_ERC721Contract.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 _index) view returns(uint256 _tokenId)
func (_ERC721Contract *ERC721ContractCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721Contract.Contract.TokenOfOwnerByIndex(&_ERC721Contract.CallOpts, _owner, _index)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _totalSupply)
func (_ERC721Contract *ERC721ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Contract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _totalSupply)
func (_ERC721Contract *ERC721ContractSession) TotalSupply() (*big.Int, error) {
	return _ERC721Contract.Contract.TotalSupply(&_ERC721Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 _totalSupply)
func (_ERC721Contract *ERC721ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721Contract.Contract.TotalSupply(&_ERC721Contract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.Contract.Approve(&_ERC721Contract.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.Contract.Approve(&_ERC721Contract.TransactOpts, _to, _tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x84ba2642.
//
// Solidity: function mint(address _owner, uint256 _tokenId, address _approvedAddress, string _metadata) returns()
func (_ERC721Contract *ERC721ContractTransactor) Mint(opts *bind.TransactOpts, _owner common.Address, _tokenId *big.Int, _approvedAddress common.Address, _metadata string) (*types.Transaction, error) {
	return _ERC721Contract.contract.Transact(opts, "mint", _owner, _tokenId, _approvedAddress, _metadata)
}

// Mint is a paid mutator transaction binding the contract method 0x84ba2642.
//
// Solidity: function mint(address _owner, uint256 _tokenId, address _approvedAddress, string _metadata) returns()
func (_ERC721Contract *ERC721ContractSession) Mint(_owner common.Address, _tokenId *big.Int, _approvedAddress common.Address, _metadata string) (*types.Transaction, error) {
	return _ERC721Contract.Contract.Mint(&_ERC721Contract.TransactOpts, _owner, _tokenId, _approvedAddress, _metadata)
}

// Mint is a paid mutator transaction binding the contract method 0x84ba2642.
//
// Solidity: function mint(address _owner, uint256 _tokenId, address _approvedAddress, string _metadata) returns()
func (_ERC721Contract *ERC721ContractTransactorSession) Mint(_owner common.Address, _tokenId *big.Int, _approvedAddress common.Address, _metadata string) (*types.Transaction, error) {
	return _ERC721Contract.Contract.Mint(&_ERC721Contract.TransactOpts, _owner, _tokenId, _approvedAddress, _metadata)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.Contract.Transfer(&_ERC721Contract.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.Contract.Transfer(&_ERC721Contract.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.Contract.TransferFrom(&_ERC721Contract.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721Contract *ERC721ContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Contract.Contract.TransferFrom(&_ERC721Contract.TransactOpts, _from, _to, _tokenId)
}

// ERC721ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721Contract contract.
type ERC721ContractApprovalIterator struct {
	Event *ERC721ContractApproval // Event containing the contract specifics and raw log

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
func (it *ERC721ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractApproval)
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
		it.Event = new(ERC721ContractApproval)
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
func (it *ERC721ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractApproval represents a Approval event raised by the ERC721Contract contract.
type ERC721ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address) (*ERC721ContractApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Contract.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractApprovalIterator{contract: _ERC721Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721ContractApproval, _owner []common.Address, _approved []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}

	logs, sub, err := _ERC721Contract.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractApproval)
				if err := _ERC721Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _approved, uint256 _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) ParseApproval(log types.Log) (*ERC721ContractApproval, error) {
	event := new(ERC721ContractApproval)
	if err := _ERC721Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ContractMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ERC721Contract contract.
type ERC721ContractMintIterator struct {
	Event *ERC721ContractMint // Event containing the contract specifics and raw log

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
func (it *ERC721ContractMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractMint)
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
		it.Event = new(ERC721ContractMint)
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
func (it *ERC721ContractMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractMint represents a Mint event raised by the ERC721Contract contract.
type ERC721ContractMint struct {
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) FilterMint(opts *bind.FilterOpts, _to []common.Address, _tokenId []*big.Int) (*ERC721ContractMintIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Contract.contract.FilterLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractMintIterator{contract: _ERC721Contract.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ERC721ContractMint, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721Contract.contract.WatchLogs(opts, "Mint", _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractMint)
				if err := _ERC721Contract.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _to, uint256 indexed _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) ParseMint(log types.Log) (*ERC721ContractMint, error) {
	event := new(ERC721ContractMint)
	if err := _ERC721Contract.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721Contract contract.
type ERC721ContractTransferIterator struct {
	Event *ERC721ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721ContractTransfer)
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
		it.Event = new(ERC721ContractTransfer)
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
func (it *ERC721ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721ContractTransfer represents a Transfer event raised by the ERC721Contract contract.
type ERC721ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ERC721ContractTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Contract.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ERC721ContractTransferIterator{contract: _ERC721Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721ContractTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ERC721Contract.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721ContractTransfer)
				if err := _ERC721Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokenId)
func (_ERC721Contract *ERC721ContractFilterer) ParseTransfer(log types.Log) (*ERC721ContractTransfer, error) {
	event := new(ERC721ContractTransfer)
	if err := _ERC721Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
