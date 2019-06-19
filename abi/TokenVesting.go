// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenVestingABI is the input ABI used to generate the binding from.
const TokenVestingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cliff\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"released\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"releasable\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"cliffDuration\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenVestingBin is the compiled bytecode used for deploying new contracts.
const TokenVestingBin = `608060405234801561001057600080fd5b506040516080806109ac833981016040818152825160208401519184015160609094015160008054600160a060020a0319163317808255929593949192600160a060020a0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3600160a060020a038416151561009157600080fd5b8082111561009e57600080fd5b600081116100ab57600080fd5b426100c384836401000000006104b461011282021704565b116100cd57600080fd5b60018054600160a060020a031916600160a060020a0386161790556004819055610104838364010000000061011281026104b41704565b60025550506003555061012b565b60008282018381101561012457600080fd5b9392505050565b6108728061013a6000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630fb5a6b481146100b357806313d033c0146100da57806319165587146100ef57806338af3eed1461011f578063715018a61461015d5780638da5cb5b146101725780638f32d59b146101875780639852595c146101b0578063a3f8eace146101de578063be9a65551461020c578063f2fde38b14610221575b600080fd5b3480156100bf57600080fd5b506100c861024f565b60408051918252519081900360200190f35b3480156100e657600080fd5b506100c8610255565b3480156100fb57600080fd5b5061011d73ffffffffffffffffffffffffffffffffffffffff6004351661025b565b005b34801561012b57600080fd5b5061013461033f565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561016957600080fd5b5061011d61035b565b34801561017e57600080fd5b506101346103dd565b34801561019357600080fd5b5061019c6103f9565b604080519115158252519081900360200190f35b3480156101bc57600080fd5b506100c873ffffffffffffffffffffffffffffffffffffffff60043516610417565b3480156101ea57600080fd5b506100c873ffffffffffffffffffffffffffffffffffffffff6004351661043f565b34801561021857600080fd5b506100c8610450565b34801561022d57600080fd5b5061011d73ffffffffffffffffffffffffffffffffffffffff60043516610456565b60045490565b60025490565b600061026682610475565b90506000811161027557600080fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152600560205260409020546102ab908263ffffffff6104b416565b73ffffffffffffffffffffffffffffffffffffffff8084166000818152600560205260409020929092556001546102ea9291168363ffffffff6104d116565b6040805173ffffffffffffffffffffffffffffffffffffffff841681526020810183905281517fc7798891864187665ac6dd119286e44ec13f014527aeeb2b8eb3fd413df93179929181900390910190a15050565b60015473ffffffffffffffffffffffffffffffffffffffff1690565b6103636103f9565b151561036e57600080fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60005473ffffffffffffffffffffffffffffffffffffffff16331490565b73ffffffffffffffffffffffffffffffffffffffff1660009081526005602052604090205490565b600061044a82610475565b92915050565b60035490565b61045e6103f9565b151561046957600080fd5b610472816105b0565b50565b73ffffffffffffffffffffffffffffffffffffffff811660009081526005602052604081205461044a906104a88461065f565b9063ffffffff6107de16565b6000828201838110156104c657600080fd5b8091505b5092915050565b8273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561057457600080fd5b505af1158015610588573d6000803e3d6000fd5b505050506040513d602081101561059e57600080fd5b505115156105ab57600080fd5b505050565b73ffffffffffffffffffffffffffffffffffffffff811615156105d257600080fd5b6000805460405173ffffffffffffffffffffffffffffffffffffffff808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60008060008373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156106ff57600080fd5b505af1158015610713573d6000803e3d6000fd5b505050506040513d602081101561072957600080fd5b505173ffffffffffffffffffffffffffffffffffffffff851660009081526005602052604090205490925061076590839063ffffffff6104b416565b905060025442101561077a57600092506107d7565b60045460035461078f9163ffffffff6104b416565b421061079d578092506107d7565b6107d46004546107c86107bb600354426107de90919063ffffffff16565b849063ffffffff6107f516565b9063ffffffff61082316565b92505b5050919050565b600080838311156107ee57600080fd5b5050900390565b60008083151561080857600091506104ca565b5082820282848281151561081857fe5b04146104c657600080fd5b60008080831161083257600080fd5b828481151561083d57fe5b049493505050505600a165627a7a72305820095037c61597656ea2c7f72aeb0be83248b63fc1e654e1d5aaadcc86002ac2680029`

// DeployTokenVesting deploys a new Ethereum contract, binding an instance of TokenVesting to it.
func DeployTokenVesting(auth *bind.TransactOpts, backend bind.ContractBackend, beneficiary common.Address, start *big.Int, cliffDuration *big.Int, duration *big.Int) (common.Address, *types.Transaction, *TokenVesting, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenVestingBin), backend, beneficiary, start, cliffDuration, duration)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// TokenVesting is an auto generated Go binding around an Ethereum contract.
type TokenVesting struct {
	TokenVestingCaller     // Read-only binding to the contract
	TokenVestingTransactor // Write-only binding to the contract
	TokenVestingFilterer   // Log filterer for contract events
}

// TokenVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenVestingSession struct {
	Contract     *TokenVesting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenVestingCallerSession struct {
	Contract *TokenVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenVestingTransactorSession struct {
	Contract     *TokenVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenVestingRaw struct {
	Contract *TokenVesting // Generic contract binding to access the raw methods on
}

// TokenVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenVestingCallerRaw struct {
	Contract *TokenVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokenVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenVestingTransactorRaw struct {
	Contract *TokenVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenVesting creates a new instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVesting(address common.Address, backend bind.ContractBackend) (*TokenVesting, error) {
	contract, err := bindTokenVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenVesting{TokenVestingCaller: TokenVestingCaller{contract: contract}, TokenVestingTransactor: TokenVestingTransactor{contract: contract}, TokenVestingFilterer: TokenVestingFilterer{contract: contract}}, nil
}

// NewTokenVestingCaller creates a new read-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingCaller(address common.Address, caller bind.ContractCaller) (*TokenVestingCaller, error) {
	contract, err := bindTokenVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingCaller{contract: contract}, nil
}

// NewTokenVestingTransactor creates a new write-only instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenVestingTransactor, error) {
	contract, err := bindTokenVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenVestingTransactor{contract: contract}, nil
}

// NewTokenVestingFilterer creates a new log filterer instance of TokenVesting, bound to a specific deployed contract.
func NewTokenVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenVestingFilterer, error) {
	contract, err := bindTokenVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenVestingFilterer{contract: contract}, nil
}

// bindTokenVesting binds a generic wrapper to an already deployed contract.
func bindTokenVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.TokenVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.TokenVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenVesting *TokenVestingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenVesting *TokenVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenVesting *TokenVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenVesting.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingSession) Beneficiary() (common.Address, error) {
	return _TokenVesting.Contract.Beneficiary(&_TokenVesting.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Beneficiary() (common.Address, error) {
	return _TokenVesting.Contract.Beneficiary(&_TokenVesting.CallOpts)
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Cliff(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "cliff")
	return *ret0, err
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Cliff() (*big.Int, error) {
	return _TokenVesting.Contract.Cliff(&_TokenVesting.CallOpts)
}

// Cliff is a free data retrieval call binding the contract method 0x13d033c0.
//
// Solidity: function cliff() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Cliff() (*big.Int, error) {
	return _TokenVesting.Contract.Cliff(&_TokenVesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "duration")
	return *ret0, err
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Duration() (*big.Int, error) {
	return _TokenVesting.Contract.Duration(&_TokenVesting.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Duration() (*big.Int, error) {
	return _TokenVesting.Contract.Duration(&_TokenVesting.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_TokenVesting *TokenVestingCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_TokenVesting *TokenVestingSession) IsOwner() (bool, error) {
	return _TokenVesting.Contract.IsOwner(&_TokenVesting.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_TokenVesting *TokenVestingCallerSession) IsOwner() (bool, error) {
	return _TokenVesting.Contract.IsOwner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenVesting *TokenVestingCallerSession) Owner() (common.Address, error) {
	return _TokenVesting.Contract.Owner(&_TokenVesting.CallOpts)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address token) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Releasable(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "releasable", token)
	return *ret0, err
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address token) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Releasable(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Releasable(&_TokenVesting.CallOpts, token)
}

// Releasable is a free data retrieval call binding the contract method 0xa3f8eace.
//
// Solidity: function releasable(address token) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Releasable(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Releasable(&_TokenVesting.CallOpts, token)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address token) constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Released(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "released", token)
	return *ret0, err
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address token) constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Released(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Released(&_TokenVesting.CallOpts, token)
}

// Released is a free data retrieval call binding the contract method 0x9852595c.
//
// Solidity: function released(address token) constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Released(token common.Address) (*big.Int, error) {
	return _TokenVesting.Contract.Released(&_TokenVesting.CallOpts, token)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenVesting.contract.Call(opts, out, "start")
	return *ret0, err
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingSession) Start() (*big.Int, error) {
	return _TokenVesting.Contract.Start(&_TokenVesting.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() constant returns(uint256)
func (_TokenVesting *TokenVestingCallerSession) Start() (*big.Int, error) {
	return _TokenVesting.Contract.Start(&_TokenVesting.CallOpts)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address token) returns()
func (_TokenVesting *TokenVestingTransactor) Release(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "release", token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address token) returns()
func (_TokenVesting *TokenVestingSession) Release(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, token)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(address token) returns()
func (_TokenVesting *TokenVestingTransactorSession) Release(token common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.Release(&_TokenVesting.TransactOpts, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenVesting.Contract.RenounceOwnership(&_TokenVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenVesting *TokenVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenVesting.Contract.RenounceOwnership(&_TokenVesting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenVesting *TokenVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenVesting *TokenVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenVesting *TokenVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenVesting.Contract.TransferOwnership(&_TokenVesting.TransactOpts, newOwner)
}

// TokenVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenVesting contract.
type TokenVestingOwnershipTransferredIterator struct {
	Event *TokenVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingOwnershipTransferred)
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
		it.Event = new(TokenVestingOwnershipTransferred)
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
func (it *TokenVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TokenVesting contract.
type TokenVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenVesting *TokenVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenVestingOwnershipTransferredIterator{contract: _TokenVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenVesting *TokenVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingOwnershipTransferred)
				if err := _TokenVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenVestingTokensReleasedIterator is returned from FilterTokensReleased and is used to iterate over the raw logs and unpacked data for TokensReleased events raised by the TokenVesting contract.
type TokenVestingTokensReleasedIterator struct {
	Event *TokenVestingTokensReleased // Event containing the contract specifics and raw log

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
func (it *TokenVestingTokensReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenVestingTokensReleased)
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
		it.Event = new(TokenVestingTokensReleased)
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
func (it *TokenVestingTokensReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenVestingTokensReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenVestingTokensReleased represents a TokensReleased event raised by the TokenVesting contract.
type TokenVestingTokensReleased struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensReleased is a free log retrieval operation binding the contract event 0xc7798891864187665ac6dd119286e44ec13f014527aeeb2b8eb3fd413df93179.
//
// Solidity: event TokensReleased(address token, uint256 amount)
func (_TokenVesting *TokenVestingFilterer) FilterTokensReleased(opts *bind.FilterOpts) (*TokenVestingTokensReleasedIterator, error) {

	logs, sub, err := _TokenVesting.contract.FilterLogs(opts, "TokensReleased")
	if err != nil {
		return nil, err
	}
	return &TokenVestingTokensReleasedIterator{contract: _TokenVesting.contract, event: "TokensReleased", logs: logs, sub: sub}, nil
}

// WatchTokensReleased is a free log subscription operation binding the contract event 0xc7798891864187665ac6dd119286e44ec13f014527aeeb2b8eb3fd413df93179.
//
// Solidity: event TokensReleased(address token, uint256 amount)
func (_TokenVesting *TokenVestingFilterer) WatchTokensReleased(opts *bind.WatchOpts, sink chan<- *TokenVestingTokensReleased) (event.Subscription, error) {

	logs, sub, err := _TokenVesting.contract.WatchLogs(opts, "TokensReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenVestingTokensReleased)
				if err := _TokenVesting.contract.UnpackLog(event, "TokensReleased", log); err != nil {
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
