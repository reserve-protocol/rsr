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

// SlowWalletABI is the input ABI used to generate the binding from.
const SlowWalletABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\"},{\"name\":\"notes\",\"type\":\"string\"},{\"name\":\"closed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"voidAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"delayUntil\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"TransferProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"TransferConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"notes\",\"type\":\"string\"}],\"name\":\"TransferCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"AllTransfersCancelled\",\"type\":\"event\"}]"

// SlowWalletBin is the compiled bytecode used for deploying new contracts.
const SlowWalletBin = `608060405234801561001057600080fd5b50604051602080610cad8339810180604052602081101561003057600080fd5b5051600080546001600160a01b039092166001600160a01b03199283161790556001805490911633179055610c438061006a6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80637b166f60116100665780637b166f60146101b55780638da5cb5b1461023a578063bf60f8ef1461025e578063d0a08d6e14610266578063fc0c546a1461029857610093565b8063013cf08b14610098578063315a69fa1461015f57806344c7c867146101935780636a42b8f8146101ad575b600080fd5b6100b5600480360360208110156100ae57600080fd5b50356102a0565b60405180866001600160a01b03166001600160a01b031681526020018581526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b83811015610120578181015183820152602001610108565b50505050905090810190601f16801561014d5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6101916004803603606081101561017557600080fd5b508035906001600160a01b036020820135169060400135610364565b005b61019b6104bd565b60408051918252519081900360200190f35b61019b6104c3565b610191600480360360608110156101cb57600080fd5b6001600160a01b03823516916020810135918101906060810160408201356401000000008111156101fb57600080fd5b82018360208201111561020d57600080fd5b8035906020019184600183028401116401000000008311171561022f57600080fd5b5090925090506104ca565b6102426106ef565b604080516001600160a01b039092168252519081900360200190f35b6101916106fe565b6101916004803603606081101561027c57600080fd5b508035906001600160a01b036020820135169060400135610783565b6102426109c5565b60026020818152600092835260409283902080546001808301548386015460038501805489516101009582161595909502600019011697909704601f81018790048702840187019098528783526001600160a01b0390931696909592949192918301828280156103515780601f1061032657610100808354040283529160200191610351565b820191906000526020600020905b81548152906001019060200180831161033457829003601f168201915b5050506004909301549192505060ff1685565b6001546001600160a01b031633146103b95760408051600160e51b62461bcd02815260206004820152600d6024820152600160991b6c36bab9ba1031329037bbb732b902604482015290519081900360640190fd5b6103c48383836109d4565b60008381526002602081815260409283902060048101805460ff191660019081179091558451888152928301869052606094830185815260039092018054600019928116156101000292909201909116939093049382018490526001600160a01b038616937fa99797fde63ee29177c0bcd12725053c1fcecc39957a12d910dad426ddefafcf938893879391929091906080830190849080156104a85780601f1061047d576101008083540402835291602001916104a8565b820191906000526020600020905b81548152906001019060200180831161048b57829003601f168201915b505094505050505060405180910390a2505050565b60035481565b6224ea0081565b6001546001600160a01b0316331461051f5760408051600160e51b62461bcd02815260206004820152600d6024820152600160991b6c36bab9ba1031329037bbb732b902604482015290519081900360640190fd5b426224ea0081019081101561057e5760408051600160e51b62461bcd02815260206004820152601060248201527f64656c6179206f766572666c6f77656400000000000000000000000000000000604482015290519081900360640190fd5b6040518060a00160405280866001600160a01b0316815260200185815260200182815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250505060209182018190526003805482526002808452604092839020855181546001600160a01b0319166001600160a01b03909116178155858501516001820155928501519083015560608401518051929361063b9392850192910190610b5a565b50608091820151600491909101805460ff191691151591909117905560038054600181019091556040805182815260208101889052908101849052606081018381529281018590526001600160a01b038816927fc21d9f71ad43be26dd663dc2ae89d8a5f2e2e627520624eeff74bf1004d80b3d929188918691899189919060a08201848480828437600083820152604051601f909101601f19169092018290039850909650505050505050a25050505050565b6001546001600160a01b031681565b6001546001600160a01b031633146107535760408051600160e51b62461bcd02815260206004820152600d6024820152600160991b6c36bab9ba1031329037bbb732b902604482015290519081900360640190fd5b600060038190556040517ff4a9bad09d9916720a4c78733bb1b637bccdc7b56ae0f78e6bb63c99934b49b99190a1565b6001546001600160a01b031633146107d85760408051600160e51b62461bcd02815260206004820152600d6024820152600160991b6c36bab9ba1031329037bbb732b902604482015290519081900360640190fd5b6107e38383836109d4565b60008381526002602081905260409091200154421161083b5760408051600160e51b62461bcd0281526020600482015260096024820152600160b81b68746f6f206561726c7902604482015290519081900360640190fd5b60008381526002602081815260409283902060048101805460ff191660019081179091558451888152928301869052606094830185815260039092018054600019928116156101000292909201909116939093049382018490526001600160a01b038616937f0725491da501611b9ba3843fe633719becb32a870a63c7008fab061512c834179388938793919290919060808301908490801561091f5780601f106108f45761010080835404028352916020019161091f565b820191906000526020600020905b81548152906001019060200180831161090257829003601f168201915b505094505050505060405180910390a26000805460408051600160e01b63a9059cbb0281526001600160a01b038681166004830152602482018690529151919092169263a9059cbb92604480820193602093909283900390910190829087803b15801561098b57600080fd5b505af115801561099f573d6000803e3d6000fd5b505050506040513d60208110156109b557600080fd5b50516109c057600080fd5b505050565b6000546001600160a01b031681565b6003548310610a1757604051600160e51b62461bcd028152600401808060200182810382526022815260200180610bf66022913960400191505060405180910390fd5b60008381526002602052604090206004015460ff1615610a815760408051600160e51b62461bcd02815260206004820152601760248201527f7472616e7366657220616c726561647920636c6f736564000000000000000000604482015290519081900360640190fd5b6000838152600260205260409020546001600160a01b03838116911614610af25760408051600160e51b62461bcd02815260206004820152601660248201527f64657374696e6174696f6e206d69736d61746368656400000000000000000000604482015290519081900360640190fd5b60008381526002602052604090206001015481146109c05760408051600160e51b62461bcd02815260206004820152601060248201527f76616c7565206d69736d61746368656400000000000000000000000000000000604482015290519081900360640190fd5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b9b57805160ff1916838001178555610bc8565b82800160010185558215610bc8579182015b82811115610bc8578251825591602001919060010190610bad565b50610bd4929150610bd8565b5090565b610bf291905b80821115610bd45760008155600101610bde565b9056fe696e64657820746f6f20686967682c206f72207472616e7366657220766f69646564a165627a7a72305820aca838febd883243b2027fc57681642ada8ef1c6086936762e926a7cfb5bd55d0029`

// DeploySlowWallet deploys a new Ethereum contract, binding an instance of SlowWallet to it.
func DeploySlowWallet(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *SlowWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(SlowWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SlowWalletBin), backend, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlowWallet{SlowWalletCaller: SlowWalletCaller{contract: contract}, SlowWalletTransactor: SlowWalletTransactor{contract: contract}, SlowWalletFilterer: SlowWalletFilterer{contract: contract}}, nil
}

// SlowWallet is an auto generated Go binding around an Ethereum contract.
type SlowWallet struct {
	SlowWalletCaller     // Read-only binding to the contract
	SlowWalletTransactor // Write-only binding to the contract
	SlowWalletFilterer   // Log filterer for contract events
}

// SlowWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlowWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlowWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlowWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlowWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlowWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlowWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlowWalletSession struct {
	Contract     *SlowWallet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlowWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlowWalletCallerSession struct {
	Contract *SlowWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SlowWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlowWalletTransactorSession struct {
	Contract     *SlowWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SlowWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlowWalletRaw struct {
	Contract *SlowWallet // Generic contract binding to access the raw methods on
}

// SlowWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlowWalletCallerRaw struct {
	Contract *SlowWalletCaller // Generic read-only contract binding to access the raw methods on
}

// SlowWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlowWalletTransactorRaw struct {
	Contract *SlowWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlowWallet creates a new instance of SlowWallet, bound to a specific deployed contract.
func NewSlowWallet(address common.Address, backend bind.ContractBackend) (*SlowWallet, error) {
	contract, err := bindSlowWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlowWallet{SlowWalletCaller: SlowWalletCaller{contract: contract}, SlowWalletTransactor: SlowWalletTransactor{contract: contract}, SlowWalletFilterer: SlowWalletFilterer{contract: contract}}, nil
}

// NewSlowWalletCaller creates a new read-only instance of SlowWallet, bound to a specific deployed contract.
func NewSlowWalletCaller(address common.Address, caller bind.ContractCaller) (*SlowWalletCaller, error) {
	contract, err := bindSlowWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlowWalletCaller{contract: contract}, nil
}

// NewSlowWalletTransactor creates a new write-only instance of SlowWallet, bound to a specific deployed contract.
func NewSlowWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*SlowWalletTransactor, error) {
	contract, err := bindSlowWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlowWalletTransactor{contract: contract}, nil
}

// NewSlowWalletFilterer creates a new log filterer instance of SlowWallet, bound to a specific deployed contract.
func NewSlowWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*SlowWalletFilterer, error) {
	contract, err := bindSlowWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlowWalletFilterer{contract: contract}, nil
}

// bindSlowWallet binds a generic wrapper to an already deployed contract.
func bindSlowWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SlowWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlowWallet *SlowWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SlowWallet.Contract.SlowWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlowWallet *SlowWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlowWallet.Contract.SlowWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlowWallet *SlowWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlowWallet.Contract.SlowWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlowWallet *SlowWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SlowWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlowWallet *SlowWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlowWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlowWallet *SlowWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlowWallet.Contract.contract.Transact(opts, method, params...)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() constant returns(uint256)
func (_SlowWallet *SlowWalletCaller) Delay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SlowWallet.contract.Call(opts, out, "delay")
	return *ret0, err
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() constant returns(uint256)
func (_SlowWallet *SlowWalletSession) Delay() (*big.Int, error) {
	return _SlowWallet.Contract.Delay(&_SlowWallet.CallOpts)
}

// Delay is a free data retrieval call binding the contract method 0x6a42b8f8.
//
// Solidity: function delay() constant returns(uint256)
func (_SlowWallet *SlowWalletCallerSession) Delay() (*big.Int, error) {
	return _SlowWallet.Contract.Delay(&_SlowWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SlowWallet *SlowWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SlowWallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SlowWallet *SlowWalletSession) Owner() (common.Address, error) {
	return _SlowWallet.Contract.Owner(&_SlowWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SlowWallet *SlowWalletCallerSession) Owner() (common.Address, error) {
	return _SlowWallet.Contract.Owner(&_SlowWallet.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(address destination, uint256 value, uint256 time, string notes, bool closed)
func (_SlowWallet *SlowWalletCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Time        *big.Int
	Notes       string
	Closed      bool
}, error) {
	ret := new(struct {
		Destination common.Address
		Value       *big.Int
		Time        *big.Int
		Notes       string
		Closed      bool
	})
	out := ret
	err := _SlowWallet.contract.Call(opts, out, "proposals", arg0)
	return *ret, err
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(address destination, uint256 value, uint256 time, string notes, bool closed)
func (_SlowWallet *SlowWalletSession) Proposals(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Time        *big.Int
	Notes       string
	Closed      bool
}, error) {
	return _SlowWallet.Contract.Proposals(&_SlowWallet.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) constant returns(address destination, uint256 value, uint256 time, string notes, bool closed)
func (_SlowWallet *SlowWalletCallerSession) Proposals(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Time        *big.Int
	Notes       string
	Closed      bool
}, error) {
	return _SlowWallet.Contract.Proposals(&_SlowWallet.CallOpts, arg0)
}

// ProposalsLength is a free data retrieval call binding the contract method 0x44c7c867.
//
// Solidity: function proposalsLength() constant returns(uint256)
func (_SlowWallet *SlowWalletCaller) ProposalsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SlowWallet.contract.Call(opts, out, "proposalsLength")
	return *ret0, err
}

// ProposalsLength is a free data retrieval call binding the contract method 0x44c7c867.
//
// Solidity: function proposalsLength() constant returns(uint256)
func (_SlowWallet *SlowWalletSession) ProposalsLength() (*big.Int, error) {
	return _SlowWallet.Contract.ProposalsLength(&_SlowWallet.CallOpts)
}

// ProposalsLength is a free data retrieval call binding the contract method 0x44c7c867.
//
// Solidity: function proposalsLength() constant returns(uint256)
func (_SlowWallet *SlowWalletCallerSession) ProposalsLength() (*big.Int, error) {
	return _SlowWallet.Contract.ProposalsLength(&_SlowWallet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SlowWallet *SlowWalletCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SlowWallet.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SlowWallet *SlowWalletSession) Token() (common.Address, error) {
	return _SlowWallet.Contract.Token(&_SlowWallet.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SlowWallet *SlowWalletCallerSession) Token() (common.Address, error) {
	return _SlowWallet.Contract.Token(&_SlowWallet.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x315a69fa.
//
// Solidity: function cancel(uint256 index, address addr, uint256 value) returns()
func (_SlowWallet *SlowWalletTransactor) Cancel(opts *bind.TransactOpts, index *big.Int, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _SlowWallet.contract.Transact(opts, "cancel", index, addr, value)
}

// Cancel is a paid mutator transaction binding the contract method 0x315a69fa.
//
// Solidity: function cancel(uint256 index, address addr, uint256 value) returns()
func (_SlowWallet *SlowWalletSession) Cancel(index *big.Int, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _SlowWallet.Contract.Cancel(&_SlowWallet.TransactOpts, index, addr, value)
}

// Cancel is a paid mutator transaction binding the contract method 0x315a69fa.
//
// Solidity: function cancel(uint256 index, address addr, uint256 value) returns()
func (_SlowWallet *SlowWalletTransactorSession) Cancel(index *big.Int, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _SlowWallet.Contract.Cancel(&_SlowWallet.TransactOpts, index, addr, value)
}

// Confirm is a paid mutator transaction binding the contract method 0xd0a08d6e.
//
// Solidity: function confirm(uint256 index, address destination, uint256 value) returns()
func (_SlowWallet *SlowWalletTransactor) Confirm(opts *bind.TransactOpts, index *big.Int, destination common.Address, value *big.Int) (*types.Transaction, error) {
	return _SlowWallet.contract.Transact(opts, "confirm", index, destination, value)
}

// Confirm is a paid mutator transaction binding the contract method 0xd0a08d6e.
//
// Solidity: function confirm(uint256 index, address destination, uint256 value) returns()
func (_SlowWallet *SlowWalletSession) Confirm(index *big.Int, destination common.Address, value *big.Int) (*types.Transaction, error) {
	return _SlowWallet.Contract.Confirm(&_SlowWallet.TransactOpts, index, destination, value)
}

// Confirm is a paid mutator transaction binding the contract method 0xd0a08d6e.
//
// Solidity: function confirm(uint256 index, address destination, uint256 value) returns()
func (_SlowWallet *SlowWalletTransactorSession) Confirm(index *big.Int, destination common.Address, value *big.Int) (*types.Transaction, error) {
	return _SlowWallet.Contract.Confirm(&_SlowWallet.TransactOpts, index, destination, value)
}

// Propose is a paid mutator transaction binding the contract method 0x7b166f60.
//
// Solidity: function propose(address destination, uint256 value, string notes) returns()
func (_SlowWallet *SlowWalletTransactor) Propose(opts *bind.TransactOpts, destination common.Address, value *big.Int, notes string) (*types.Transaction, error) {
	return _SlowWallet.contract.Transact(opts, "propose", destination, value, notes)
}

// Propose is a paid mutator transaction binding the contract method 0x7b166f60.
//
// Solidity: function propose(address destination, uint256 value, string notes) returns()
func (_SlowWallet *SlowWalletSession) Propose(destination common.Address, value *big.Int, notes string) (*types.Transaction, error) {
	return _SlowWallet.Contract.Propose(&_SlowWallet.TransactOpts, destination, value, notes)
}

// Propose is a paid mutator transaction binding the contract method 0x7b166f60.
//
// Solidity: function propose(address destination, uint256 value, string notes) returns()
func (_SlowWallet *SlowWalletTransactorSession) Propose(destination common.Address, value *big.Int, notes string) (*types.Transaction, error) {
	return _SlowWallet.Contract.Propose(&_SlowWallet.TransactOpts, destination, value, notes)
}

// VoidAll is a paid mutator transaction binding the contract method 0xbf60f8ef.
//
// Solidity: function voidAll() returns()
func (_SlowWallet *SlowWalletTransactor) VoidAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlowWallet.contract.Transact(opts, "voidAll")
}

// VoidAll is a paid mutator transaction binding the contract method 0xbf60f8ef.
//
// Solidity: function voidAll() returns()
func (_SlowWallet *SlowWalletSession) VoidAll() (*types.Transaction, error) {
	return _SlowWallet.Contract.VoidAll(&_SlowWallet.TransactOpts)
}

// VoidAll is a paid mutator transaction binding the contract method 0xbf60f8ef.
//
// Solidity: function voidAll() returns()
func (_SlowWallet *SlowWalletTransactorSession) VoidAll() (*types.Transaction, error) {
	return _SlowWallet.Contract.VoidAll(&_SlowWallet.TransactOpts)
}

// SlowWalletAllTransfersCancelledIterator is returned from FilterAllTransfersCancelled and is used to iterate over the raw logs and unpacked data for AllTransfersCancelled events raised by the SlowWallet contract.
type SlowWalletAllTransfersCancelledIterator struct {
	Event *SlowWalletAllTransfersCancelled // Event containing the contract specifics and raw log

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
func (it *SlowWalletAllTransfersCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlowWalletAllTransfersCancelled)
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
		it.Event = new(SlowWalletAllTransfersCancelled)
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
func (it *SlowWalletAllTransfersCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlowWalletAllTransfersCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlowWalletAllTransfersCancelled represents a AllTransfersCancelled event raised by the SlowWallet contract.
type SlowWalletAllTransfersCancelled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllTransfersCancelled is a free log retrieval operation binding the contract event 0xf4a9bad09d9916720a4c78733bb1b637bccdc7b56ae0f78e6bb63c99934b49b9.
//
// Solidity: event AllTransfersCancelled()
func (_SlowWallet *SlowWalletFilterer) FilterAllTransfersCancelled(opts *bind.FilterOpts) (*SlowWalletAllTransfersCancelledIterator, error) {

	logs, sub, err := _SlowWallet.contract.FilterLogs(opts, "AllTransfersCancelled")
	if err != nil {
		return nil, err
	}
	return &SlowWalletAllTransfersCancelledIterator{contract: _SlowWallet.contract, event: "AllTransfersCancelled", logs: logs, sub: sub}, nil
}

// WatchAllTransfersCancelled is a free log subscription operation binding the contract event 0xf4a9bad09d9916720a4c78733bb1b637bccdc7b56ae0f78e6bb63c99934b49b9.
//
// Solidity: event AllTransfersCancelled()
func (_SlowWallet *SlowWalletFilterer) WatchAllTransfersCancelled(opts *bind.WatchOpts, sink chan<- *SlowWalletAllTransfersCancelled) (event.Subscription, error) {

	logs, sub, err := _SlowWallet.contract.WatchLogs(opts, "AllTransfersCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlowWalletAllTransfersCancelled)
				if err := _SlowWallet.contract.UnpackLog(event, "AllTransfersCancelled", log); err != nil {
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

// SlowWalletTransferCancelledIterator is returned from FilterTransferCancelled and is used to iterate over the raw logs and unpacked data for TransferCancelled events raised by the SlowWallet contract.
type SlowWalletTransferCancelledIterator struct {
	Event *SlowWalletTransferCancelled // Event containing the contract specifics and raw log

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
func (it *SlowWalletTransferCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlowWalletTransferCancelled)
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
		it.Event = new(SlowWalletTransferCancelled)
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
func (it *SlowWalletTransferCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlowWalletTransferCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlowWalletTransferCancelled represents a TransferCancelled event raised by the SlowWallet contract.
type SlowWalletTransferCancelled struct {
	Index       *big.Int
	Destination common.Address
	Value       *big.Int
	Notes       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferCancelled is a free log retrieval operation binding the contract event 0xa99797fde63ee29177c0bcd12725053c1fcecc39957a12d910dad426ddefafcf.
//
// Solidity: event TransferCancelled(uint256 index, address indexed destination, uint256 value, string notes)
func (_SlowWallet *SlowWalletFilterer) FilterTransferCancelled(opts *bind.FilterOpts, destination []common.Address) (*SlowWalletTransferCancelledIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SlowWallet.contract.FilterLogs(opts, "TransferCancelled", destinationRule)
	if err != nil {
		return nil, err
	}
	return &SlowWalletTransferCancelledIterator{contract: _SlowWallet.contract, event: "TransferCancelled", logs: logs, sub: sub}, nil
}

// WatchTransferCancelled is a free log subscription operation binding the contract event 0xa99797fde63ee29177c0bcd12725053c1fcecc39957a12d910dad426ddefafcf.
//
// Solidity: event TransferCancelled(uint256 index, address indexed destination, uint256 value, string notes)
func (_SlowWallet *SlowWalletFilterer) WatchTransferCancelled(opts *bind.WatchOpts, sink chan<- *SlowWalletTransferCancelled, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SlowWallet.contract.WatchLogs(opts, "TransferCancelled", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlowWalletTransferCancelled)
				if err := _SlowWallet.contract.UnpackLog(event, "TransferCancelled", log); err != nil {
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

// SlowWalletTransferConfirmedIterator is returned from FilterTransferConfirmed and is used to iterate over the raw logs and unpacked data for TransferConfirmed events raised by the SlowWallet contract.
type SlowWalletTransferConfirmedIterator struct {
	Event *SlowWalletTransferConfirmed // Event containing the contract specifics and raw log

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
func (it *SlowWalletTransferConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlowWalletTransferConfirmed)
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
		it.Event = new(SlowWalletTransferConfirmed)
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
func (it *SlowWalletTransferConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlowWalletTransferConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlowWalletTransferConfirmed represents a TransferConfirmed event raised by the SlowWallet contract.
type SlowWalletTransferConfirmed struct {
	Index       *big.Int
	Destination common.Address
	Value       *big.Int
	Notes       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferConfirmed is a free log retrieval operation binding the contract event 0x0725491da501611b9ba3843fe633719becb32a870a63c7008fab061512c83417.
//
// Solidity: event TransferConfirmed(uint256 index, address indexed destination, uint256 value, string notes)
func (_SlowWallet *SlowWalletFilterer) FilterTransferConfirmed(opts *bind.FilterOpts, destination []common.Address) (*SlowWalletTransferConfirmedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SlowWallet.contract.FilterLogs(opts, "TransferConfirmed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &SlowWalletTransferConfirmedIterator{contract: _SlowWallet.contract, event: "TransferConfirmed", logs: logs, sub: sub}, nil
}

// WatchTransferConfirmed is a free log subscription operation binding the contract event 0x0725491da501611b9ba3843fe633719becb32a870a63c7008fab061512c83417.
//
// Solidity: event TransferConfirmed(uint256 index, address indexed destination, uint256 value, string notes)
func (_SlowWallet *SlowWalletFilterer) WatchTransferConfirmed(opts *bind.WatchOpts, sink chan<- *SlowWalletTransferConfirmed, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SlowWallet.contract.WatchLogs(opts, "TransferConfirmed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlowWalletTransferConfirmed)
				if err := _SlowWallet.contract.UnpackLog(event, "TransferConfirmed", log); err != nil {
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

// SlowWalletTransferProposedIterator is returned from FilterTransferProposed and is used to iterate over the raw logs and unpacked data for TransferProposed events raised by the SlowWallet contract.
type SlowWalletTransferProposedIterator struct {
	Event *SlowWalletTransferProposed // Event containing the contract specifics and raw log

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
func (it *SlowWalletTransferProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlowWalletTransferProposed)
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
		it.Event = new(SlowWalletTransferProposed)
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
func (it *SlowWalletTransferProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlowWalletTransferProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlowWalletTransferProposed represents a TransferProposed event raised by the SlowWallet contract.
type SlowWalletTransferProposed struct {
	Index       *big.Int
	Destination common.Address
	Value       *big.Int
	DelayUntil  *big.Int
	Notes       string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferProposed is a free log retrieval operation binding the contract event 0xc21d9f71ad43be26dd663dc2ae89d8a5f2e2e627520624eeff74bf1004d80b3d.
//
// Solidity: event TransferProposed(uint256 index, address indexed destination, uint256 value, uint256 delayUntil, string notes)
func (_SlowWallet *SlowWalletFilterer) FilterTransferProposed(opts *bind.FilterOpts, destination []common.Address) (*SlowWalletTransferProposedIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SlowWallet.contract.FilterLogs(opts, "TransferProposed", destinationRule)
	if err != nil {
		return nil, err
	}
	return &SlowWalletTransferProposedIterator{contract: _SlowWallet.contract, event: "TransferProposed", logs: logs, sub: sub}, nil
}

// WatchTransferProposed is a free log subscription operation binding the contract event 0xc21d9f71ad43be26dd663dc2ae89d8a5f2e2e627520624eeff74bf1004d80b3d.
//
// Solidity: event TransferProposed(uint256 index, address indexed destination, uint256 value, uint256 delayUntil, string notes)
func (_SlowWallet *SlowWalletFilterer) WatchTransferProposed(opts *bind.WatchOpts, sink chan<- *SlowWalletTransferProposed, destination []common.Address) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}

	logs, sub, err := _SlowWallet.contract.WatchLogs(opts, "TransferProposed", destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlowWalletTransferProposed)
				if err := _SlowWallet.contract.UnpackLog(event, "TransferProposed", log); err != nil {
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
