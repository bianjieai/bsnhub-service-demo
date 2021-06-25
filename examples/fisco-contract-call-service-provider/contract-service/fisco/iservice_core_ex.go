// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fisco

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

// IServiceCoreExABI is the input ABI used to generate the binding from.
const IServiceCoreExABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"name\":\"_errMsg\",\"type\":\"string\"},{\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_endpointInfo\",\"type\":\"string\"},{\"name\":\"_method\",\"type\":\"string\"},{\"name\":\"_callData\",\"type\":\"bytes\"},{\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"sendRequest\",\"outputs\":[{\"name\":\"requestID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_icRequestID\",\"type\":\"bytes32\"},{\"name\":\"_endpointAddress\",\"type\":\"address\"},{\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"callService\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_eventName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_endpointInfo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_method\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_callData\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CrossChainRequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_eventName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_icRequestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"CrossChainResponseSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IServiceCoreExBin is the compiled bytecode used for deploying new contracts.
var IServiceCoreExBin = "0x60806040526040805190810160405280600981526020017f666973636f2d312d310000000000000000000000000000000000000000000000815250600190805190602001906200005192919062000247565b503480156200005f57600080fd5b50604051602080620018e283398101806040528101908080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515620001bd5780600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000217565b620001d66200021e640100000000026401000000009004565b600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50620002f6565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028a57805160ff1916838001178555620002bb565b82800160010185558215620002bb579182015b82811115620002ba5782518255916020019190600101906200029d565b5b509050620002ca9190620002ce565b5090565b620002f391905b80821115620002ef576000816000905550600101620002d5565b5090565b90565b6115dc80620003066000396000f300608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063161e3b9a1461009e57806316cad12a146101175780633a938f2d1461015a5780634f8da64f146101855780635089e2c8146102555780637a1f2dea146102ac578063e9667acd146102ef578063e9f7862d14610346578063ede8e529146103dd575b600080fd5b3480156100aa57600080fd5b506100fd600480360381019080803560001916906020019092919080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939050505061040c565b604051808215151515815260200191505060405180910390f35b34801561012357600080fd5b50610158600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108f3565b005b34801561016657600080fd5b5061016f61097b565b6040518082815260200191505060405180910390f35b34801561019157600080fd5b50610237600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610981565b60405180826000191660001916815260200191505060405180910390f35b34801561026157600080fd5b5061026a610e50565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102b857600080fd5b506102ed600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e79565b005b3480156102fb57600080fd5b50610304611004565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561035257600080fd5b506103db6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061102a565b005b3480156103e957600080fd5b506103f261124a565b604051808215151515815260200191505060405180910390f35b6000610416611561565b60606000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610505576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f6953657276696365436f726545783a2073656e646572206973206e6f7420746881526020017f652072656c61796572000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b886000151560026000836000191660001916815260200190815260200160002060009054906101000a900460ff1615151415156105d0576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001807f6953657276696365436f726545783a207265717565737420686173206265656e81526020017f20726573706f6e6465640000000000000000000000000000000000000000000081525060400191505060405180910390fd5b600360008b600019166000191681526020019081526020016000206040805190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a90047c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050935060008989905011156107055788888080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050925061073b565b86868080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505092505b6001600260008c6000191660001916815260200190815260200160002060006101000a81548160ff021916908315150217905550836000015173ffffffffffffffffffffffffffffffffffffffff1684602001518b8560405160240180836000191660001916815260200180602001828103825283818151815260200191508051906020019080838360005b838110156107e25780820151818401526020810190506107c7565b50505050905090810190601f16801561080f5780820380516001836020036101000a031916815260200191505b509350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405180828051906020019080838360005b8381101561089f578082015181840152602081019050610884565b50505050905090810190601f1680156108cc5780820380516001836020036101000a031916815260200191505b509150506000604051808303816000865af191505091508194505050505095945050505050565b6108fb61124a565b151561096f576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b610978816112a1565b50565b60055481565b600088888080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505087878080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505086868080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505060008351111515610abb576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252602c8152602001807f6953657276696365436f726545783a2064657374436861696e49442063616e2081526020017f6e6f7420626520656d707479000000000000000000000000000000000000000081525060400191505060405180910390fd5b60008251111515610b5a576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260278152602001807f6953657276696365436f726545783a206d6574686f642063616e206e6f74206281526020017f6520656d7074790000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60008151111515610bf9576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f6953657276696365436f726545783a2063616c6c446174612063616e206e6f7481526020017f20626520656d707479000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b60016005546040516020018083805460018160011615610100020316600290048015610c5c5780601f10610c3a576101008083540402835291820191610c5c565b820191906000526020600020905b815481529060010190602001808311610c48575b5050828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083101515610cac5780518252602082019150602081019050602083039250610c87565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902093506005600081548092919060010191905055507fffb1b4b73f42832bdbd4342fcd29c78b84aa9ec2517e0cfb4aab9e28729d5f3a848d8d8d8d8d8d3360405180806020018a600019166000191681526020018060200180602001806020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001858103855260158152602001807f43726f7373436861696e5265717565737453656e74000000000000000000000081525060200185810384528c8c82818152602001925080828437820191505085810383528a8a828181526020019250808284378201915050858103825288888281815260200192508082843782019150509c5050505050505050505050505060405180910390a1610e0a84878761142a565b600060026000866000191660001916815260200190815260200160002060006101000a81548160ff02191690831515021790555083935050505098975050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610e8161124a565b1515610ef5576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610fc0576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f6953657276696365436f726545783a2072656c6179657220616464726573732081526020017f63616e206e6f74206265207a65726f000000000000000000000000000000000081525060400191505060405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060606000856000151560046000836000191660001916815260200190815260200160002060009054906101000a900460ff1615151415156110fb576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b8152602001807f6953657276696365436f726545783a206963526571756573742068617320626581526020017f656e20726563656976656400000000000000000000000000000000000000000081525060400191505060405180910390fd5b84519350602085016000808683348b5af192508260018114611124576000811461114b57611150565b3d6040519550601f19601f6020830101168601604052808652806000602088013e50611150565b600080fd5b50506001821415611241577f695d8fe31b685f93fcaa25b0e4ff24eb6ed49a1eee43b7acc656f33046c4dcdf87846040518080602001846000191660001916815260200180602001838103835260168152602001807f43726f7373436861696e526573706f6e736553656e7400000000000000000000815250602001838103825284818151815260200191508051906020019080838360005b838110156112045780820151818401526020810190506111e9565b50505050905090810190601f1680156112315780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15b50505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561136c576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526020017f646472657373000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b611432611561565b82816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508181602001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815250508060036000866000191660001916815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff02191690837c01000000000000000000000000000000000000000000000000000000009004021790555090505050505050565b6040805190810160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815250905600a165627a7a723058200de40999560d64751a92495c627abb473cd7a5bcf3377403294c45b8e858b7280029"

// DeployIServiceCoreEx deploys a new contract, binding an instance of IServiceCoreEx to it.
func DeployIServiceCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend, _relayer common.Address) (common.Address, *types.Transaction, *IServiceCoreEx, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IServiceCoreExBin), backend, _relayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

func AsyncDeployIServiceCoreEx(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _relayer common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(IServiceCoreExBin), backend, _relayer)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// IServiceCoreEx is an auto generated Go binding around a Solidity contract.
type IServiceCoreEx struct {
	IServiceCoreExCaller     // Read-only binding to the contract
	IServiceCoreExTransactor // Write-only binding to the contract
	IServiceCoreExFilterer   // Log filterer for contract events
}

// IServiceCoreExCaller is an auto generated read-only Go binding around a Solidity contract.
type IServiceCoreExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExTransactor is an auto generated write-only Go binding around a Solidity contract.
type IServiceCoreExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IServiceCoreExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IServiceCoreExSession struct {
	Contract     *IServiceCoreEx   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceCoreExCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IServiceCoreExCallerSession struct {
	Contract *IServiceCoreExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IServiceCoreExTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IServiceCoreExTransactorSession struct {
	Contract     *IServiceCoreExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IServiceCoreExRaw is an auto generated low-level Go binding around a Solidity contract.
type IServiceCoreExRaw struct {
	Contract *IServiceCoreEx // Generic contract binding to access the raw methods on
}

// IServiceCoreExCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IServiceCoreExCallerRaw struct {
	Contract *IServiceCoreExCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceCoreExTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IServiceCoreExTransactorRaw struct {
	Contract *IServiceCoreExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIServiceCoreEx creates a new instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreEx(address common.Address, backend bind.ContractBackend) (*IServiceCoreEx, error) {
	contract, err := bindIServiceCoreEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

// NewIServiceCoreExCaller creates a new read-only instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExCaller(address common.Address, caller bind.ContractCaller) (*IServiceCoreExCaller, error) {
	contract, err := bindIServiceCoreEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCaller{contract: contract}, nil
}

// NewIServiceCoreExTransactor creates a new write-only instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExTransactor(address common.Address, transactor bind.ContractTransactor) (*IServiceCoreExTransactor, error) {
	contract, err := bindIServiceCoreEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExTransactor{contract: contract}, nil
}

// NewIServiceCoreExFilterer creates a new log filterer instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExFilterer(address common.Address, filterer bind.ContractFilterer) (*IServiceCoreExFilterer, error) {
	contract, err := bindIServiceCoreEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExFilterer{contract: contract}, nil
}

// bindIServiceCoreEx binds a generic wrapper to an already deployed contract.
func bindIServiceCoreEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceCoreEx *IServiceCoreExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceCoreEx.Contract.IServiceCoreExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceCoreEx *IServiceCoreExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceCoreEx *IServiceCoreExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceCoreEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExCallerSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "relayer")
	return *ret0, err
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "requestCount")
	return *ret0, err
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCallerSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _icRequestID, address _endpointAddress, bytes _callData) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) CallService(opts *bind.TransactOpts, _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "callService", _icRequestID, _endpointAddress, _callData)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncCallService(handler func(*types.Receipt, error), opts *bind.TransactOpts, _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "callService", _icRequestID, _endpointAddress, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _icRequestID, address _endpointAddress, bytes _callData) returns()
func (_IServiceCoreEx *IServiceCoreExSession) CallService(_icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncCallService(handler func(*types.Receipt, error), _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService(handler, &_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _icRequestID, address _endpointAddress, bytes _callData) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) CallService(_icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncCallService(handler func(*types.Receipt, error), _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService(handler, &_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactor) SendRequest(opts *bind.TransactOpts, _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "sendRequest", _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSendRequest(handler func(*types.Receipt, error), opts *bind.TransactOpts, _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "sendRequest", _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExSession) SendRequest(_endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SendRequest(&_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSendRequest(handler func(*types.Receipt, error), _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSendRequest(handler, &_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SendRequest(_endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SendRequest(&_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSendRequest(handler func(*types.Receipt, error), _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSendRequest(handler, &_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setRelayer", _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetRelayer(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExSession) SetRelayer(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetRelayer(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetRelayer(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetRelayer(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetRelayer(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactor) SetResponse(opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setResponse", _requestID, _errMsg, _output)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setResponse", _requestID, _errMsg, _output)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) SetResponse(_requestID [32]byte, _errMsg string, _output string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetResponse(handler func(*types.Receipt, error), _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetResponse(handler, &_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetResponse(_requestID [32]byte, _errMsg string, _output string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetResponse(handler func(*types.Receipt, error), _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetResponse(handler, &_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncTransferOwnership(handler, &_IServiceCoreEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncTransferOwnership(handler, &_IServiceCoreEx.TransactOpts, newOwner)
}

// IServiceCoreExCrossChainRequestSentIterator is returned from FilterCrossChainRequestSent and is used to iterate over the raw logs and unpacked data for CrossChainRequestSent events raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainRequestSentIterator struct {
	Event *IServiceCoreExCrossChainRequestSent // Event containing the contract specifics and raw log

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
func (it *IServiceCoreExCrossChainRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExCrossChainRequestSent)
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
		it.Event = new(IServiceCoreExCrossChainRequestSent)
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
func (it *IServiceCoreExCrossChainRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExCrossChainRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExCrossChainRequestSent represents a CrossChainRequestSent event raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainRequestSent struct {
	EventName    string
	RequestID    [32]byte
	EndpointInfo string
	Method       string
	CallData     []byte
	Sender       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCrossChainRequestSent is a free log retrieval operation binding the contract event 0xbdf5c91926e766003ce81f528fb359c226a22c61137fe24429574c5e0c62994b.
//
// Solidity: event CrossChainRequestSent(string _eventName, bytes32 _requestID, string _endpointInfo, string _method, bytes _callData, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterCrossChainRequestSent(opts *bind.FilterOpts) (*IServiceCoreExCrossChainRequestSentIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "CrossChainRequestSent")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCrossChainRequestSentIterator{contract: _IServiceCoreEx.contract, event: "CrossChainRequestSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainRequestSent is a free log subscription operation binding the contract event 0xbdf5c91926e766003ce81f528fb359c226a22c61137fe24429574c5e0c62994b.
//
// Solidity: event CrossChainRequestSent(string _eventName, bytes32 _requestID, string _endpointInfo, string _method, bytes _callData, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchCrossChainRequestSent(opts *bind.WatchOpts, sink chan<- *IServiceCoreExCrossChainRequestSent) (event.Subscription, error) {

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "CrossChainRequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExCrossChainRequestSent)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainRequestSent", log); err != nil {
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

// ParseCrossChainRequestSent is a log parse operation binding the contract event 0xbdf5c91926e766003ce81f528fb359c226a22c61137fe24429574c5e0c62994b.
//
// Solidity: event CrossChainRequestSent(string _eventName, bytes32 _requestID, string _endpointInfo, string _method, bytes _callData, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseCrossChainRequestSent(log types.Log) (*IServiceCoreExCrossChainRequestSent, error) {
	event := new(IServiceCoreExCrossChainRequestSent)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainRequestSent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceCoreExCrossChainResponseSentIterator is returned from FilterCrossChainResponseSent and is used to iterate over the raw logs and unpacked data for CrossChainResponseSent events raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainResponseSentIterator struct {
	Event *IServiceCoreExCrossChainResponseSent // Event containing the contract specifics and raw log

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
func (it *IServiceCoreExCrossChainResponseSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExCrossChainResponseSent)
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
		it.Event = new(IServiceCoreExCrossChainResponseSent)
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
func (it *IServiceCoreExCrossChainResponseSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExCrossChainResponseSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExCrossChainResponseSent represents a CrossChainResponseSent event raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainResponseSent struct {
	EventName   string
	IcRequestID [32]byte
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCrossChainResponseSent is a free log retrieval operation binding the contract event 0x982b787eefda447c92d33a6dca5a03ed134eb2ec3d1c0135f628dcb761bf1709.
//
// Solidity: event CrossChainResponseSent(string _eventName, bytes32 _icRequestID, bytes _result)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterCrossChainResponseSent(opts *bind.FilterOpts) (*IServiceCoreExCrossChainResponseSentIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCrossChainResponseSentIterator{contract: _IServiceCoreEx.contract, event: "CrossChainResponseSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainResponseSent is a free log subscription operation binding the contract event 0x982b787eefda447c92d33a6dca5a03ed134eb2ec3d1c0135f628dcb761bf1709.
//
// Solidity: event CrossChainResponseSent(string _eventName, bytes32 _icRequestID, bytes _result)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchCrossChainResponseSent(opts *bind.WatchOpts, sink chan<- *IServiceCoreExCrossChainResponseSent) (event.Subscription, error) {

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExCrossChainResponseSent)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
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

// ParseCrossChainResponseSent is a log parse operation binding the contract event 0x982b787eefda447c92d33a6dca5a03ed134eb2ec3d1c0135f628dcb761bf1709.
//
// Solidity: event CrossChainResponseSent(string _eventName, bytes32 _icRequestID, bytes _result)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseCrossChainResponseSent(log types.Log) (*IServiceCoreExCrossChainResponseSent, error) {
	event := new(IServiceCoreExCrossChainResponseSent)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceCoreExOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IServiceCoreEx contract.
type IServiceCoreExOwnershipTransferredIterator struct {
	Event *IServiceCoreExOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IServiceCoreExOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExOwnershipTransferred)
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
		it.Event = new(IServiceCoreExOwnershipTransferred)
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
func (it *IServiceCoreExOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExOwnershipTransferred represents a OwnershipTransferred event raised by the IServiceCoreEx contract.
type IServiceCoreExOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IServiceCoreExOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExOwnershipTransferredIterator{contract: _IServiceCoreEx.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IServiceCoreExOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExOwnershipTransferred)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseOwnershipTransferred(log types.Log) (*IServiceCoreExOwnershipTransferred, error) {
	event := new(IServiceCoreExOwnershipTransferred)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
