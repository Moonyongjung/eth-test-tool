// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rpc

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
)

// EthTestToolMetaData contains all meta data concerning the EthTestTool contract.
var EthTestToolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100d9565b60405180910390f35b610073600480360381019061006e919061009d565b61007e565b005b60008054905090565b8060008190555050565b60008135905061009781610103565b92915050565b6000602082840312156100b3576100b26100fe565b5b60006100c184828501610088565b91505092915050565b6100d3816100f4565b82525050565b60006020820190506100ee60008301846100ca565b92915050565b6000819050919050565b600080fd5b61010c816100f4565b811461011757600080fd5b5056fea26469706673582212209a159a4f3847890f10bfb87871a61eba91c5dbf5ee3cf6398207e292eee22a1664736f6c63430008070033",
}

// EthTestToolABI is the input ABI used to generate the binding from.
// Deprecated: Use EthTestToolMetaData.ABI instead.
var EthTestToolABI = EthTestToolMetaData.ABI

// EthTestToolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthTestToolMetaData.Bin instead.
var EthTestToolBin = EthTestToolMetaData.Bin

// DeployEthTestTool deploys a new Ethereum contract, binding an instance of EthTestTool to it.
func DeployEthTestTool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthTestTool, error) {
	parsed, err := EthTestToolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthTestToolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthTestTool{EthTestToolCaller: EthTestToolCaller{contract: contract}, EthTestToolTransactor: EthTestToolTransactor{contract: contract}, EthTestToolFilterer: EthTestToolFilterer{contract: contract}}, nil
}

// EthTestTool is an auto generated Go binding around an Ethereum contract.
type EthTestTool struct {
	EthTestToolCaller     // Read-only binding to the contract
	EthTestToolTransactor // Write-only binding to the contract
	EthTestToolFilterer   // Log filterer for contract events
}

// EthTestToolCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthTestToolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTestToolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthTestToolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTestToolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthTestToolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTestToolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthTestToolSession struct {
	Contract     *EthTestTool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthTestToolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthTestToolCallerSession struct {
	Contract *EthTestToolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EthTestToolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthTestToolTransactorSession struct {
	Contract     *EthTestToolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EthTestToolRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthTestToolRaw struct {
	Contract *EthTestTool // Generic contract binding to access the raw methods on
}

// EthTestToolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthTestToolCallerRaw struct {
	Contract *EthTestToolCaller // Generic read-only contract binding to access the raw methods on
}

// EthTestToolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthTestToolTransactorRaw struct {
	Contract *EthTestToolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthTestTool creates a new instance of EthTestTool, bound to a specific deployed contract.
func NewEthTestTool(address common.Address, backend bind.ContractBackend) (*EthTestTool, error) {
	contract, err := bindEthTestTool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthTestTool{EthTestToolCaller: EthTestToolCaller{contract: contract}, EthTestToolTransactor: EthTestToolTransactor{contract: contract}, EthTestToolFilterer: EthTestToolFilterer{contract: contract}}, nil
}

// NewEthTestToolCaller creates a new read-only instance of EthTestTool, bound to a specific deployed contract.
func NewEthTestToolCaller(address common.Address, caller bind.ContractCaller) (*EthTestToolCaller, error) {
	contract, err := bindEthTestTool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthTestToolCaller{contract: contract}, nil
}

// NewEthTestToolTransactor creates a new write-only instance of EthTestTool, bound to a specific deployed contract.
func NewEthTestToolTransactor(address common.Address, transactor bind.ContractTransactor) (*EthTestToolTransactor, error) {
	contract, err := bindEthTestTool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthTestToolTransactor{contract: contract}, nil
}

// NewEthTestToolFilterer creates a new log filterer instance of EthTestTool, bound to a specific deployed contract.
func NewEthTestToolFilterer(address common.Address, filterer bind.ContractFilterer) (*EthTestToolFilterer, error) {
	contract, err := bindEthTestTool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthTestToolFilterer{contract: contract}, nil
}

// bindEthTestTool binds a generic wrapper to an already deployed contract.
func bindEthTestTool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthTestToolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthTestTool *EthTestToolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthTestTool.Contract.EthTestToolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthTestTool *EthTestToolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthTestTool.Contract.EthTestToolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthTestTool *EthTestToolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthTestTool.Contract.EthTestToolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthTestTool *EthTestToolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthTestTool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthTestTool *EthTestToolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthTestTool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthTestTool *EthTestToolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthTestTool.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_EthTestTool *EthTestToolCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthTestTool.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_EthTestTool *EthTestToolSession) Retrieve() (*big.Int, error) {
	return _EthTestTool.Contract.Retrieve(&_EthTestTool.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_EthTestTool *EthTestToolCallerSession) Retrieve() (*big.Int, error) {
	return _EthTestTool.Contract.Retrieve(&_EthTestTool.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_EthTestTool *EthTestToolTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _EthTestTool.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_EthTestTool *EthTestToolSession) Store(num *big.Int) (*types.Transaction, error) {
	return _EthTestTool.Contract.Store(&_EthTestTool.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_EthTestTool *EthTestToolTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _EthTestTool.Contract.Store(&_EthTestTool.TransactOpts, num)
}
