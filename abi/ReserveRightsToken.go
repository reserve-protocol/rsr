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

// ReserveRightsTokenABI is the input ABI used to generate the binding from.
const ReserveRightsTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"consent\",\"type\":\"string\"}],\"name\":\"lockMyTokensForever\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserveTeamMemberOrEarlyInvestor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"previousContract\",\"type\":\"address\"},{\"name\":\"reservePrimaryWallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"lockedAccount\",\"type\":\"address\"}],\"name\":\"AccountLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ReserveRightsTokenBin is the compiled bytecode used for deploying new contracts.
const ReserveRightsTokenBin = `60c0604052600e60808190527f526573657276652052696768747300000000000000000000000000000000000060a090815262000040916005919062000890565b506040805180820190915260038082527f52535200000000000000000000000000000000000000000000000000000000006020909201918252620000879160069162000890565b506007805460ff191660121790556040805161048081018252738ad9c8ebe26eadab9251b8fc36cd06a1ec399a7f815273b268c230720d16c69a61cbee24731e3b2a3330a1602082015273082705fabf49bd30de8f0222821f6d940713b89d9181019190915273c3aa4ced5dea58a3d1ca76e507515c79ca1e443660608201527366f25f036eb4463d8a45c6594a325f9e89baa6db6080820152739e454fe7d8e087fcac4ec8c40562de781004477e60a0820152734fcc7ca22680aed155f981eeb13089383d624aa960c0820152735a66650e5345d76eb8136ea1490cbcce1c08072e60e082015273698a10b5d0972bffea306ba5950bd74d2af3c7ca61010082015273df437625216cca3d7148e18d09f4aab0d47c763b6101208201527324b4a6847ccb32972de40170c02fda121ddc6a30610140820152738d29a24f91df381feb4ee7f05405d3fb888c643e610160820152735a7350d95b9e644dcab4bc642707f43a361bf62861018082015273fc2e9a5cd1bb9b3953ffa7e6ddf0c0447eb95f116101a0820152733ac7a6c3a2ff08613b611485f795d07e785cbb956101c08201527347fc47cbcc5217740905e16c4c953b2f247369d26101e082015273d282337950ac6e936d0f0ebaaff1ffc3de79f3d561020082015273de59cd3aa43a2bf863723662b31906660c7d12b6610220820152735f84660cabb98f7b7764cd1ae2553442da91984e61024082015273efbaaf73fc22f70785515c1e2be3d5ba2fb8e9b06102608201527363c5ffb388d83477a15eb940cfa23991ca0b30f06102808201527314f018cce044f9d3fb1e1644db6f2fab70f6e3cb6102a082015273be30069d27a250f90c2ee5507bcaca5f868265f76102c082015273cfef27288bedcd587a1ed6e86a996c8c5b01d7c16102e0820152735f57bbccc7ffa4c46864b5ed999a271bc36bb0ce61030082015273bae85de9858375706dde5907c8c9c6ee22b19212610320820152735cf4bbb0ff093f3c725abec32fba8f34e4e98af161034082015273cb2d434bf72d3cd43d0c368493971183640ffe996103608201527302fc8e99401b970c265480140721b28bb3af85ab61038082015273e7ad11517d7254f6a0758cee932bffa328002dd06103a0820152736b39195c164d693d3b6518b70d99877d4f7c87ef6103c082015273c59119d8e4d129890036a108aed9d9fe94db1ba96103e082015273d28661e4c75d177d9c1f3c8b821902c1abd103a661040082015273ba385610025b1ea8091ae3e4a2e98913e2691ff761042082015273cd74834b8f3f71d2e82c6240ae0291c56378535661044082015273657a127639b9e0ccccfbe795a8e394d5ca1585266104608201526200047390600990602462000915565b503480156200048157600080fd5b506040516040806200182a833981016040528051602090910151600080620004b233640100000000620006d2810204565b6004805460ff19168155604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038087169382019390935290518694506200056b928692908616916370a08231916024808201926020929091908290030181600087803b1580156200052e57600080fd5b505af115801562000543573d6000803e3d6000fd5b505050506040513d60208110156200055a57600080fd5b505164010000000062000724810204565b5060005b600954811015620006c8576001600860006009848154811015156200059057fe5b600091825260208083209190910154600160a060020a031683528201929092526040019020805460ff19169115159190911790556009805462000673919083908110620005d957fe5b60009182526020909120015460098054600160a060020a03928316928616916370a0823191869081106200060957fe5b6000918252602080832090910154604080517c010000000000000000000000000000000000000000000000000000000063ffffffff8716028152600160a060020a0390921660048301525160248083019491928390030190829087803b1580156200052e57600080fd5b60098054829081106200068257fe5b6000918252602082200154604051600160a060020a03909116917f78be06d07afe380e04d6deeba0f33c892db454f303fb739d9b768987a5ec6aca91a26001016200056f565b50505050620009c2565b620006ed60038264010000000062000cbf620007e382021704565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b600160a060020a03821615156200073a57600080fd5b60025462000757908264010000000062000c616200083e82021704565b600255600160a060020a0382166000908152602081905260409020546200078d908264010000000062000c616200083e82021704565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a0381161515620007f957600080fd5b6200080e828264010000000062000858810204565b156200081957600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b6000828201838110156200085157600080fd5b9392505050565b6000600160a060020a03821615156200087057600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620008d357805160ff191683800117855562000903565b8280016001018555821562000903579182015b8281111562000903578251825591602001919060010190620008e6565b50620009119291506200097b565b5090565b8280548282559060005260206000209081019282156200096d579160200282015b828111156200096d5782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000936565b50620009119291506200099b565b6200099891905b8082111562000911576000815560010162000982565b90565b6200099891905b8082111562000911578054600160a060020a0319168155600101620009a2565b610e5880620009d26000396000f3006080604052600436106101065763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166303ff0fec811461010b57806306fdde0314610178578063095ea7b31461020257806318160ddd1461022657806323b872dd1461024d578063313ce5671461027757806339509351146102a25780633f4ba83a146102c657806346fbf68e146102dd5780635c975abb146102fe5780636ef8d66d1461031357806370a082311461032857806382dc1ec4146103495780638456cb591461036a57806391cdccec1461037f57806395d89b41146103a0578063a457c2d7146103b5578063a9059cbb146103d9578063dd62ed3e146103fd575b600080fd5b34801561011757600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526101649436949293602493928401919081908401838280828437509497506104249650505050505050565b604080519115158252519081900360200190f35b34801561018457600080fd5b5061018d610621565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101c75781810151838201526020016101af565b50505050905090810190601f1680156101f45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020e57600080fd5b50610164600160a060020a03600435166024356106af565b34801561023257600080fd5b5061023b6106d3565b60408051918252519081900360200190f35b34801561025957600080fd5b50610164600160a060020a03600435811690602435166044356106d9565b34801561028357600080fd5b5061028c610712565b6040805160ff9092168252519081900360200190f35b3480156102ae57600080fd5b50610164600160a060020a036004351660243561071b565b3480156102d257600080fd5b506102db610738565b005b3480156102e957600080fd5b50610164600160a060020a036004351661079c565b34801561030a57600080fd5b506101646107b5565b34801561031f57600080fd5b506102db6107be565b34801561033457600080fd5b5061023b600160a060020a03600435166107c9565b34801561035557600080fd5b506102db600160a060020a03600435166107e4565b34801561037657600080fd5b506102db610804565b34801561038b57600080fd5b50610164600160a060020a036004351661086a565b3480156103ac57600080fd5b5061018d61087f565b3480156103c157600080fd5b50610164600160a060020a03600435166024356108da565b3480156103e557600080fd5b50610164600160a060020a03600435166024356108f7565b34801561040957600080fd5b5061023b600160a060020a036004358116906024351661091e565b600060405160200180807f4920756e6465727374616e642074686174204920616d206c6f636b696e67206d81526020017f79206163636f756e7420666f72657665722c206f72206174206c65617374207581526020017f6e74696c20746865206e65787420746f6b656e20757067726164652e00000000815250605c0190506040516020818303038152906040526040518082805190602001908083835b602083106104e15780518252601f1990920191602091820191016104c2565b51815160209384036101000a60001901801990921691161790526040519190930181900381208751909550879450908301928392508401908083835b6020831061053c5780518252601f19909201916020918201910161051d565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040526040518082805190602001908083835b6020831061059f5780518252601f199092019160209182019101610580565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600019161415156105dc57600080fd5b33600081815260086020526040808220805460ff19166001179055517f78be06d07afe380e04d6deeba0f33c892db454f303fb739d9b768987a5ec6aca9190a2919050565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b820191906000526020600020905b81548152906001019060200180831161068a57829003601f168201915b505050505081565b60045460009060ff16156106c257600080fd5b6106cc8383610949565b9392505050565b60025490565b600160a060020a03831660009081526008602052604081205460ff16156106ff57600080fd5b61070a8484846109c7565b949350505050565b60075460ff1681565b60045460009060ff161561072e57600080fd5b6106cc83836109e5565b6107413361079c565b151561074c57600080fd5b60045460ff16151561075d57600080fd5b6004805460ff191690556040805133815290517f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa9181900360200190a1565b60006107af60038363ffffffff610a9516565b92915050565b60045460ff1690565b6107c733610acc565b565b600160a060020a031660009081526020819052604090205490565b6107ed3361079c565b15156107f857600080fd5b61080181610b14565b50565b61080d3361079c565b151561081857600080fd5b60045460ff161561082857600080fd5b6004805460ff191660011790556040805133815290517f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2589181900360200190a1565b60086020526000908152604090205460ff1681565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106a75780601f1061067c576101008083540402835291602001916106a7565b60045460009060ff16156108ed57600080fd5b6106cc8383610b5c565b3360009081526008602052604081205460ff161561091457600080fd5b6106cc8383610ba7565b600160a060020a03918216600090815260016020908152604080832093909416825291909152205490565b6000600160a060020a038316151561096057600080fd5b336000818152600160209081526040808320600160a060020a03881680855290835292819020869055805186815290519293927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929181900390910190a350600192915050565b60045460009060ff16156109da57600080fd5b61070a848484610bc4565b6000600160a060020a03831615156109fc57600080fd5b336000908152600160209081526040808320600160a060020a0387168452909152902054610a30908363ffffffff610c6116565b336000818152600160209081526040808320600160a060020a0389168085529083529281902085905580519485525191937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a350600192915050565b6000600160a060020a0382161515610aac57600080fd5b50600160a060020a03166000908152602091909152604090205460ff1690565b610add60038263ffffffff610c7316565b604051600160a060020a038216907fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e90600090a250565b610b2560038263ffffffff610cbf16565b604051600160a060020a038216907f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f890600090a250565b6000600160a060020a0383161515610b7357600080fd5b336000908152600160209081526040808320600160a060020a0387168452909152902054610a30908363ffffffff610d0d16565b60045460009060ff1615610bba57600080fd5b6106cc8383610d24565b600160a060020a0383166000908152600160209081526040808320338452909152812054821115610bf457600080fd5b600160a060020a0384166000908152600160209081526040808320338452909152902054610c28908363ffffffff610d0d16565b600160a060020a0385166000908152600160209081526040808320338452909152902055610c57848484610d3a565b5060019392505050565b6000828201838110156106cc57600080fd5b600160a060020a0381161515610c8857600080fd5b610c928282610a95565b1515610c9d57600080fd5b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0381161515610cd457600080fd5b610cde8282610a95565b15610ce857600080fd5b600160a060020a0316600090815260209190915260409020805460ff19166001179055565b60008083831115610d1d57600080fd5b5050900390565b6000610d31338484610d3a565b50600192915050565b600160a060020a038316600090815260208190526040902054811115610d5f57600080fd5b600160a060020a0382161515610d7457600080fd5b600160a060020a038316600090815260208190526040902054610d9d908263ffffffff610d0d16565b600160a060020a038085166000908152602081905260408082209390935590841681522054610dd2908263ffffffff610c6116565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35050505600a165627a7a72305820e18b1b13331dea7aba2e3e5057a49d7ccd3165c8cbcb998b457f754a57b8fb890029`

// DeployReserveRightsToken deploys a new Ethereum contract, binding an instance of ReserveRightsToken to it.
func DeployReserveRightsToken(auth *bind.TransactOpts, backend bind.ContractBackend, previousContract common.Address, reservePrimaryWallet common.Address) (common.Address, *types.Transaction, *ReserveRightsToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveRightsTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveRightsTokenBin), backend, previousContract, reservePrimaryWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReserveRightsToken{ReserveRightsTokenCaller: ReserveRightsTokenCaller{contract: contract}, ReserveRightsTokenTransactor: ReserveRightsTokenTransactor{contract: contract}, ReserveRightsTokenFilterer: ReserveRightsTokenFilterer{contract: contract}}, nil
}

// ReserveRightsToken is an auto generated Go binding around an Ethereum contract.
type ReserveRightsToken struct {
	ReserveRightsTokenCaller     // Read-only binding to the contract
	ReserveRightsTokenTransactor // Write-only binding to the contract
	ReserveRightsTokenFilterer   // Log filterer for contract events
}

// ReserveRightsTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveRightsTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveRightsTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveRightsTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveRightsTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveRightsTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveRightsTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveRightsTokenSession struct {
	Contract     *ReserveRightsToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReserveRightsTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveRightsTokenCallerSession struct {
	Contract *ReserveRightsTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ReserveRightsTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveRightsTokenTransactorSession struct {
	Contract     *ReserveRightsTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ReserveRightsTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveRightsTokenRaw struct {
	Contract *ReserveRightsToken // Generic contract binding to access the raw methods on
}

// ReserveRightsTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveRightsTokenCallerRaw struct {
	Contract *ReserveRightsTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveRightsTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveRightsTokenTransactorRaw struct {
	Contract *ReserveRightsTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserveRightsToken creates a new instance of ReserveRightsToken, bound to a specific deployed contract.
func NewReserveRightsToken(address common.Address, backend bind.ContractBackend) (*ReserveRightsToken, error) {
	contract, err := bindReserveRightsToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsToken{ReserveRightsTokenCaller: ReserveRightsTokenCaller{contract: contract}, ReserveRightsTokenTransactor: ReserveRightsTokenTransactor{contract: contract}, ReserveRightsTokenFilterer: ReserveRightsTokenFilterer{contract: contract}}, nil
}

// NewReserveRightsTokenCaller creates a new read-only instance of ReserveRightsToken, bound to a specific deployed contract.
func NewReserveRightsTokenCaller(address common.Address, caller bind.ContractCaller) (*ReserveRightsTokenCaller, error) {
	contract, err := bindReserveRightsToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenCaller{contract: contract}, nil
}

// NewReserveRightsTokenTransactor creates a new write-only instance of ReserveRightsToken, bound to a specific deployed contract.
func NewReserveRightsTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveRightsTokenTransactor, error) {
	contract, err := bindReserveRightsToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenTransactor{contract: contract}, nil
}

// NewReserveRightsTokenFilterer creates a new log filterer instance of ReserveRightsToken, bound to a specific deployed contract.
func NewReserveRightsTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveRightsTokenFilterer, error) {
	contract, err := bindReserveRightsToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenFilterer{contract: contract}, nil
}

// bindReserveRightsToken binds a generic wrapper to an already deployed contract.
func bindReserveRightsToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveRightsTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveRightsToken *ReserveRightsTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveRightsToken.Contract.ReserveRightsTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveRightsToken *ReserveRightsTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.ReserveRightsTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveRightsToken *ReserveRightsTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.ReserveRightsTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveRightsToken *ReserveRightsTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveRightsToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveRightsToken *ReserveRightsTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveRightsToken *ReserveRightsTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ReserveRightsToken.Contract.Allowance(&_ReserveRightsToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ReserveRightsToken.Contract.Allowance(&_ReserveRightsToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ReserveRightsToken.Contract.BalanceOf(&_ReserveRightsToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ReserveRightsToken.Contract.BalanceOf(&_ReserveRightsToken.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ReserveRightsToken *ReserveRightsTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ReserveRightsToken *ReserveRightsTokenSession) Decimals() (uint8, error) {
	return _ReserveRightsToken.Contract.Decimals(&_ReserveRightsToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) Decimals() (uint8, error) {
	return _ReserveRightsToken.Contract.Decimals(&_ReserveRightsToken.CallOpts)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "isPauser", account)
	return *ret0, err
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) IsPauser(account common.Address) (bool, error) {
	return _ReserveRightsToken.Contract.IsPauser(&_ReserveRightsToken.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) IsPauser(account common.Address) (bool, error) {
	return _ReserveRightsToken.Contract.IsPauser(&_ReserveRightsToken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveRightsToken *ReserveRightsTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveRightsToken *ReserveRightsTokenSession) Name() (string, error) {
	return _ReserveRightsToken.Contract.Name(&_ReserveRightsToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) Name() (string, error) {
	return _ReserveRightsToken.Contract.Name(&_ReserveRightsToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) Paused() (bool, error) {
	return _ReserveRightsToken.Contract.Paused(&_ReserveRightsToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) Paused() (bool, error) {
	return _ReserveRightsToken.Contract.Paused(&_ReserveRightsToken.CallOpts)
}

// ReserveTeamMemberOrEarlyInvestor is a free data retrieval call binding the contract method 0x91cdccec.
//
// Solidity: function reserveTeamMemberOrEarlyInvestor(address ) constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenCaller) ReserveTeamMemberOrEarlyInvestor(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "reserveTeamMemberOrEarlyInvestor", arg0)
	return *ret0, err
}

// ReserveTeamMemberOrEarlyInvestor is a free data retrieval call binding the contract method 0x91cdccec.
//
// Solidity: function reserveTeamMemberOrEarlyInvestor(address ) constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) ReserveTeamMemberOrEarlyInvestor(arg0 common.Address) (bool, error) {
	return _ReserveRightsToken.Contract.ReserveTeamMemberOrEarlyInvestor(&_ReserveRightsToken.CallOpts, arg0)
}

// ReserveTeamMemberOrEarlyInvestor is a free data retrieval call binding the contract method 0x91cdccec.
//
// Solidity: function reserveTeamMemberOrEarlyInvestor(address ) constant returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) ReserveTeamMemberOrEarlyInvestor(arg0 common.Address) (bool, error) {
	return _ReserveRightsToken.Contract.ReserveTeamMemberOrEarlyInvestor(&_ReserveRightsToken.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReserveRightsToken *ReserveRightsTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReserveRightsToken *ReserveRightsTokenSession) Symbol() (string, error) {
	return _ReserveRightsToken.Contract.Symbol(&_ReserveRightsToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) Symbol() (string, error) {
	return _ReserveRightsToken.Contract.Symbol(&_ReserveRightsToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveRightsToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenSession) TotalSupply() (*big.Int, error) {
	return _ReserveRightsToken.Contract.TotalSupply(&_ReserveRightsToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveRightsToken *ReserveRightsTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ReserveRightsToken.Contract.TotalSupply(&_ReserveRightsToken.CallOpts)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ReserveRightsToken *ReserveRightsTokenSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.AddPauser(&_ReserveRightsToken.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.AddPauser(&_ReserveRightsToken.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Approve(&_ReserveRightsToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Approve(&_ReserveRightsToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_ReserveRightsToken *ReserveRightsTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_ReserveRightsToken *ReserveRightsTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.DecreaseAllowance(&_ReserveRightsToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool success)
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.DecreaseAllowance(&_ReserveRightsToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_ReserveRightsToken *ReserveRightsTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_ReserveRightsToken *ReserveRightsTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.IncreaseAllowance(&_ReserveRightsToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool success)
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.IncreaseAllowance(&_ReserveRightsToken.TransactOpts, spender, addedValue)
}

// LockMyTokensForever is a paid mutator transaction binding the contract method 0x03ff0fec.
//
// Solidity: function lockMyTokensForever(string consent) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactor) LockMyTokensForever(opts *bind.TransactOpts, consent string) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "lockMyTokensForever", consent)
}

// LockMyTokensForever is a paid mutator transaction binding the contract method 0x03ff0fec.
//
// Solidity: function lockMyTokensForever(string consent) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) LockMyTokensForever(consent string) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.LockMyTokensForever(&_ReserveRightsToken.TransactOpts, consent)
}

// LockMyTokensForever is a paid mutator transaction binding the contract method 0x03ff0fec.
//
// Solidity: function lockMyTokensForever(string consent) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) LockMyTokensForever(consent string) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.LockMyTokensForever(&_ReserveRightsToken.TransactOpts, consent)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ReserveRightsToken *ReserveRightsTokenSession) Pause() (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Pause(&_ReserveRightsToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Pause(&_ReserveRightsToken.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ReserveRightsToken *ReserveRightsTokenSession) RenouncePauser() (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.RenouncePauser(&_ReserveRightsToken.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.RenouncePauser(&_ReserveRightsToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Transfer(&_ReserveRightsToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Transfer(&_ReserveRightsToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.TransferFrom(&_ReserveRightsToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.TransferFrom(&_ReserveRightsToken.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveRightsToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ReserveRightsToken *ReserveRightsTokenSession) Unpause() (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Unpause(&_ReserveRightsToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ReserveRightsToken *ReserveRightsTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _ReserveRightsToken.Contract.Unpause(&_ReserveRightsToken.TransactOpts)
}

// ReserveRightsTokenAccountLockedIterator is returned from FilterAccountLocked and is used to iterate over the raw logs and unpacked data for AccountLocked events raised by the ReserveRightsToken contract.
type ReserveRightsTokenAccountLockedIterator struct {
	Event *ReserveRightsTokenAccountLocked // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenAccountLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenAccountLocked)
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
		it.Event = new(ReserveRightsTokenAccountLocked)
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
func (it *ReserveRightsTokenAccountLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenAccountLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenAccountLocked represents a AccountLocked event raised by the ReserveRightsToken contract.
type ReserveRightsTokenAccountLocked struct {
	LockedAccount common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAccountLocked is a free log retrieval operation binding the contract event 0x78be06d07afe380e04d6deeba0f33c892db454f303fb739d9b768987a5ec6aca.
//
// Solidity: event AccountLocked(address indexed lockedAccount)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterAccountLocked(opts *bind.FilterOpts, lockedAccount []common.Address) (*ReserveRightsTokenAccountLockedIterator, error) {

	var lockedAccountRule []interface{}
	for _, lockedAccountItem := range lockedAccount {
		lockedAccountRule = append(lockedAccountRule, lockedAccountItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "AccountLocked", lockedAccountRule)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenAccountLockedIterator{contract: _ReserveRightsToken.contract, event: "AccountLocked", logs: logs, sub: sub}, nil
}

// WatchAccountLocked is a free log subscription operation binding the contract event 0x78be06d07afe380e04d6deeba0f33c892db454f303fb739d9b768987a5ec6aca.
//
// Solidity: event AccountLocked(address indexed lockedAccount)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchAccountLocked(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenAccountLocked, lockedAccount []common.Address) (event.Subscription, error) {

	var lockedAccountRule []interface{}
	for _, lockedAccountItem := range lockedAccount {
		lockedAccountRule = append(lockedAccountRule, lockedAccountItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "AccountLocked", lockedAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenAccountLocked)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "AccountLocked", log); err != nil {
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

// ReserveRightsTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ReserveRightsToken contract.
type ReserveRightsTokenApprovalIterator struct {
	Event *ReserveRightsTokenApproval // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenApproval)
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
		it.Event = new(ReserveRightsTokenApproval)
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
func (it *ReserveRightsTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenApproval represents a Approval event raised by the ReserveRightsToken contract.
type ReserveRightsTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ReserveRightsTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenApprovalIterator{contract: _ReserveRightsToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenApproval)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ReserveRightsTokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ReserveRightsToken contract.
type ReserveRightsTokenPausedIterator struct {
	Event *ReserveRightsTokenPaused // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenPaused)
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
		it.Event = new(ReserveRightsTokenPaused)
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
func (it *ReserveRightsTokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenPaused represents a Paused event raised by the ReserveRightsToken contract.
type ReserveRightsTokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterPaused(opts *bind.FilterOpts) (*ReserveRightsTokenPausedIterator, error) {

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenPausedIterator{contract: _ReserveRightsToken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenPaused) (event.Subscription, error) {

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenPaused)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ReserveRightsTokenPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the ReserveRightsToken contract.
type ReserveRightsTokenPauserAddedIterator struct {
	Event *ReserveRightsTokenPauserAdded // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenPauserAdded)
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
		it.Event = new(ReserveRightsTokenPauserAdded)
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
func (it *ReserveRightsTokenPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenPauserAdded represents a PauserAdded event raised by the ReserveRightsToken contract.
type ReserveRightsTokenPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterPauserAdded(opts *bind.FilterOpts, account []common.Address) (*ReserveRightsTokenPauserAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenPauserAddedIterator{contract: _ReserveRightsToken.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address indexed account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenPauserAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "PauserAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenPauserAdded)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ReserveRightsTokenPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the ReserveRightsToken contract.
type ReserveRightsTokenPauserRemovedIterator struct {
	Event *ReserveRightsTokenPauserRemoved // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenPauserRemoved)
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
		it.Event = new(ReserveRightsTokenPauserRemoved)
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
func (it *ReserveRightsTokenPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenPauserRemoved represents a PauserRemoved event raised by the ReserveRightsToken contract.
type ReserveRightsTokenPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterPauserRemoved(opts *bind.FilterOpts, account []common.Address) (*ReserveRightsTokenPauserRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenPauserRemovedIterator{contract: _ReserveRightsToken.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address indexed account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenPauserRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "PauserRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenPauserRemoved)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ReserveRightsTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ReserveRightsToken contract.
type ReserveRightsTokenTransferIterator struct {
	Event *ReserveRightsTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenTransfer)
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
		it.Event = new(ReserveRightsTokenTransfer)
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
func (it *ReserveRightsTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenTransfer represents a Transfer event raised by the ReserveRightsToken contract.
type ReserveRightsTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ReserveRightsTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenTransferIterator{contract: _ReserveRightsToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenTransfer)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ReserveRightsTokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ReserveRightsToken contract.
type ReserveRightsTokenUnpausedIterator struct {
	Event *ReserveRightsTokenUnpaused // Event containing the contract specifics and raw log

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
func (it *ReserveRightsTokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRightsTokenUnpaused)
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
		it.Event = new(ReserveRightsTokenUnpaused)
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
func (it *ReserveRightsTokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRightsTokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRightsTokenUnpaused represents a Unpaused event raised by the ReserveRightsToken contract.
type ReserveRightsTokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ReserveRightsTokenUnpausedIterator, error) {

	logs, sub, err := _ReserveRightsToken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ReserveRightsTokenUnpausedIterator{contract: _ReserveRightsToken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ReserveRightsToken *ReserveRightsTokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ReserveRightsTokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _ReserveRightsToken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRightsTokenUnpaused)
				if err := _ReserveRightsToken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
