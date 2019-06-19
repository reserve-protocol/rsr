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
const SlowWalletBin = `608060405234801561001057600080fd5b50604051602080610fec8339810180604052602081101561003057600080fd5b5051600080546001600160a01b039092166001600160a01b03199283161790556001805490911633179055610f828061006a6000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80637b166f6011610076578063bf60f8ef1161005b578063bf60f8ef146102af578063d0a08d6e146102b7578063fc0c546a146102f6576100a3565b80637b166f60146101ec5780638da5cb5b1461027e576100a3565b8063013cf08b146100a8578063315a69fa1461018957806344c7c867146101ca5780636a42b8f8146101e4575b600080fd5b6100c5600480360360208110156100be57600080fd5b50356102fe565b604051808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018581526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561014a578181015183820152602001610132565b50505050905090810190601f1680156101775780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6101c86004803603606081101561019f57600080fd5b5080359073ffffffffffffffffffffffffffffffffffffffff60208201351690604001356103ed565b005b6101d26105c0565b60408051918252519081900360200190f35b6101d26105c6565b6101c86004803603606081101561020257600080fd5b73ffffffffffffffffffffffffffffffffffffffff8235169160208101359181019060608101604082013564010000000081111561023f57600080fd5b82018360208201111561025157600080fd5b8035906020019184600183028401116401000000008311171561027357600080fd5b5090925090506105cd565b6102866108b5565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101c86108d1565b6101c8600480360360608110156102cd57600080fd5b5080359073ffffffffffffffffffffffffffffffffffffffff6020820135169060400135610987565b610286610c8e565b600260208181526000928352604092839020805460018083015483860154600385018054895161010095821615959095027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff011697909704601f810187900487028401870190985287835273ffffffffffffffffffffffffffffffffffffffff90931696909592949192918301828280156103da5780601f106103af576101008083540402835291602001916103da565b820191906000526020600020905b8154815290600101906020018083116103bd57829003601f168201915b5050506004909301549192505060ff1685565b60015473ffffffffffffffffffffffffffffffffffffffff16331461047357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d757374206265206f776e657200000000000000000000000000000000000000604482015290519081900360640190fd5b61047e838383610caa565b6000838152600260208181526040928390206004810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584518881529283018690526060948301858152600390920180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9281161561010002929092019091169390930493820184905273ffffffffffffffffffffffffffffffffffffffff8616937fa99797fde63ee29177c0bcd12725053c1fcecc39957a12d910dad426ddefafcf938893879391929091906080830190849080156105ab5780601f10610580576101008083540402835291602001916105ab565b820191906000526020600020905b81548152906001019060200180831161058e57829003601f168201915b505094505050505060405180910390a2505050565b60035481565b6224ea0081565b60015473ffffffffffffffffffffffffffffffffffffffff16331461065357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d757374206265206f776e657200000000000000000000000000000000000000604482015290519081900360640190fd5b426224ea008101908110156106c957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f64656c6179206f766572666c6f77656400000000000000000000000000000000604482015290519081900360640190fd5b6040518060a001604052808673ffffffffffffffffffffffffffffffffffffffff16815260200185815260200182815260200184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525093855250505060209182018190526003805482526002808452604092839020855181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617815585850151600182015592850151908301556060840151805192936107b89392850192910190610e99565b5060809182015160049190910180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016911515919091179055600380546001810190915560408051828152602081018890529081018490526060810183815292810185905273ffffffffffffffffffffffffffffffffffffffff8816927fc21d9f71ad43be26dd663dc2ae89d8a5f2e2e627520624eeff74bf1004d80b3d929188918691899189919060a08201848480828437600083820152604051601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169092018290039850909650505050505050a25050505050565b60015473ffffffffffffffffffffffffffffffffffffffff1681565b60015473ffffffffffffffffffffffffffffffffffffffff16331461095757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d757374206265206f776e657200000000000000000000000000000000000000604482015290519081900360640190fd5b600060038190556040517ff4a9bad09d9916720a4c78733bb1b637bccdc7b56ae0f78e6bb63c99934b49b99190a1565b60015473ffffffffffffffffffffffffffffffffffffffff163314610a0d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600d60248201527f6d757374206265206f776e657200000000000000000000000000000000000000604482015290519081900360640190fd5b610a18838383610caa565b600083815260026020819052604090912001544211610a9857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f746f6f206561726c790000000000000000000000000000000000000000000000604482015290519081900360640190fd5b6000838152600260208181526040928390206004810180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915584518881529283018690526060948301858152600390920180547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9281161561010002929092019091169390930493820184905273ffffffffffffffffffffffffffffffffffffffff8616937f0725491da501611b9ba3843fe633719becb32a870a63c7008fab061512c8341793889387939192909190608083019084908015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b820191906000526020600020905b815481529060010190602001808311610ba857829003601f168201915b505094505050505060405180910390a260008054604080517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018690529151919092169263a9059cbb92604480820193602093909283900390910190829087803b158015610c5457600080fd5b505af1158015610c68573d6000803e3d6000fd5b505050506040513d6020811015610c7e57600080fd5b5051610c8957600080fd5b505050565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6003548310610d04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610f356022913960400191505060405180910390fd5b60008381526002602052604090206004015460ff1615610d8557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f7472616e7366657220616c726561647920636c6f736564000000000000000000604482015290519081900360640190fd5b60008381526002602052604090205473ffffffffffffffffffffffffffffffffffffffff838116911614610e1a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f64657374696e6174696f6e206d69736d61746368656400000000000000000000604482015290519081900360640190fd5b6000838152600260205260409020600101548114610c8957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f76616c7565206d69736d61746368656400000000000000000000000000000000604482015290519081900360640190fd5b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610eda57805160ff1916838001178555610f07565b82800160010185558215610f07579182015b82811115610f07578251825591602001919060010190610eec565b50610f13929150610f17565b5090565b610f3191905b80821115610f135760008155600101610f1d565b9056fe696e64657820746f6f20686967682c206f72207472616e7366657220766f69646564a165627a7a72305820239554896be14c0da4b3f70a031c5711107a1a48560a17af39ac77b864f529060029`

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
