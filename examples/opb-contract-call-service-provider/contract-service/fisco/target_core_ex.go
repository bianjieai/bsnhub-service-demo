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

// TargetCoreExABI is the input ABI used to generate the binding from.
const TargetCoreExABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_RequestID\",\"type\":\"bytes32\"},{\"name\":\"_endpointAddress\",\"type\":\"address\"},{\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"callService\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_RequestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"CrossChainResponseSent\",\"type\":\"event\"}]"

// TargetCoreExBin is the compiled bytecode used for deploying new contracts.
var TargetCoreExBin = "0x608060405234801561001057600080fd5b506102ef806100206000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063e9f7862d14610046575b600080fd5b34801561005257600080fd5b506100db6004803603810190808035600019169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506100dd565b005b6000606060008560001515600080836000191660001916815260200190815260200160002060009054906101000a900460ff1615151415156101ad576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f6953657276696365436f726545783a206475706c69636174656420726571756581526020017f737421000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b84519350602085016000808683348b5af1925082600181146101d657600081146101fd57610202565b3d6040519550601f19601f6020830101168601604052808652806000602088013e50610202565b600080fd5b505060018214156102ba577f335ab62b3c41f74c3265d91b5a6e7a87c73608ea084d3d3f6c4ee5636fa10956878460405180836000191660001916815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561027e578082015181840152602081019050610263565b50505050905090810190601f1680156102ab5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15b505050505050505600a165627a7a72305820c321dc5286125286d15e65c7ad4b66d39d2540df0a81694e0bf8b982a32333180029"

// DeployTargetCoreEx deploys a new contract, binding an instance of TargetCoreEx to it.
func DeployTargetCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TargetCoreEx, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetCoreExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TargetCoreExBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TargetCoreEx{TargetCoreExCaller: TargetCoreExCaller{contract: contract}, TargetCoreExTransactor: TargetCoreExTransactor{contract: contract}, TargetCoreExFilterer: TargetCoreExFilterer{contract: contract}}, nil
}

func AsyncDeployTargetCoreEx(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetCoreExABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(TargetCoreExBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// TargetCoreEx is an auto generated Go binding around a Solidity contract.
type TargetCoreEx struct {
	TargetCoreExCaller     // Read-only binding to the contract
	TargetCoreExTransactor // Write-only binding to the contract
	TargetCoreExFilterer   // Log filterer for contract events
}

// TargetCoreExCaller is an auto generated read-only Go binding around a Solidity contract.
type TargetCoreExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetCoreExTransactor is an auto generated write-only Go binding around a Solidity contract.
type TargetCoreExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetCoreExFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type TargetCoreExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TargetCoreExSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type TargetCoreExSession struct {
	Contract     *TargetCoreEx     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TargetCoreExCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type TargetCoreExCallerSession struct {
	Contract *TargetCoreExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TargetCoreExTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type TargetCoreExTransactorSession struct {
	Contract     *TargetCoreExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TargetCoreExRaw is an auto generated low-level Go binding around a Solidity contract.
type TargetCoreExRaw struct {
	Contract *TargetCoreEx // Generic contract binding to access the raw methods on
}

// TargetCoreExCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type TargetCoreExCallerRaw struct {
	Contract *TargetCoreExCaller // Generic read-only contract binding to access the raw methods on
}

// TargetCoreExTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type TargetCoreExTransactorRaw struct {
	Contract *TargetCoreExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTargetCoreEx creates a new instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreEx(address common.Address, backend bind.ContractBackend) (*TargetCoreEx, error) {
	contract, err := bindTargetCoreEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TargetCoreEx{TargetCoreExCaller: TargetCoreExCaller{contract: contract}, TargetCoreExTransactor: TargetCoreExTransactor{contract: contract}, TargetCoreExFilterer: TargetCoreExFilterer{contract: contract}}, nil
}

// NewTargetCoreExCaller creates a new read-only instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreExCaller(address common.Address, caller bind.ContractCaller) (*TargetCoreExCaller, error) {
	contract, err := bindTargetCoreEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TargetCoreExCaller{contract: contract}, nil
}

// NewTargetCoreExTransactor creates a new write-only instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreExTransactor(address common.Address, transactor bind.ContractTransactor) (*TargetCoreExTransactor, error) {
	contract, err := bindTargetCoreEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TargetCoreExTransactor{contract: contract}, nil
}

// NewTargetCoreExFilterer creates a new log filterer instance of TargetCoreEx, bound to a specific deployed contract.
func NewTargetCoreExFilterer(address common.Address, filterer bind.ContractFilterer) (*TargetCoreExFilterer, error) {
	contract, err := bindTargetCoreEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TargetCoreExFilterer{contract: contract}, nil
}

// bindTargetCoreEx binds a generic wrapper to an already deployed contract.
func bindTargetCoreEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TargetCoreExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TargetCoreEx *TargetCoreExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TargetCoreEx.Contract.TargetCoreExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TargetCoreEx *TargetCoreExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.Contract.TargetCoreExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TargetCoreEx *TargetCoreExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.Contract.TargetCoreExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TargetCoreEx *TargetCoreExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TargetCoreEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TargetCoreEx *TargetCoreExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TargetCoreEx *TargetCoreExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.Contract.contract.Transact(opts, method, params...)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _RequestID, address _endpointAddress, bytes _callData) returns()
func (_TargetCoreEx *TargetCoreExTransactor) CallService(opts *bind.TransactOpts, _RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.contract.Transact(opts, "callService", _RequestID, _endpointAddress, _callData)
}

func (_TargetCoreEx *TargetCoreExTransactor) AsyncCallService(handler func(*types.Receipt, error), opts *bind.TransactOpts, _RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _TargetCoreEx.contract.AsyncTransact(opts, handler, "callService", _RequestID, _endpointAddress, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _RequestID, address _endpointAddress, bytes _callData) returns()
func (_TargetCoreEx *TargetCoreExSession) CallService(_RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.Contract.CallService(&_TargetCoreEx.TransactOpts, _RequestID, _endpointAddress, _callData)
}

func (_TargetCoreEx *TargetCoreExSession) AsyncCallService(handler func(*types.Receipt, error), _RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.AsyncCallService(handler, &_TargetCoreEx.TransactOpts, _RequestID, _endpointAddress, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _RequestID, address _endpointAddress, bytes _callData) returns()
func (_TargetCoreEx *TargetCoreExTransactorSession) CallService(_RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _TargetCoreEx.Contract.CallService(&_TargetCoreEx.TransactOpts, _RequestID, _endpointAddress, _callData)
}

func (_TargetCoreEx *TargetCoreExTransactorSession) AsyncCallService(handler func(*types.Receipt, error), _RequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _TargetCoreEx.Contract.AsyncCallService(handler, &_TargetCoreEx.TransactOpts, _RequestID, _endpointAddress, _callData)
}

// TargetCoreExCrossChainResponseSentIterator is returned from FilterCrossChainResponseSent and is used to iterate over the raw logs and unpacked data for CrossChainResponseSent events raised by the TargetCoreEx contract.
type TargetCoreExCrossChainResponseSentIterator struct {
	Event *TargetCoreExCrossChainResponseSent // Event containing the contract specifics and raw log

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
func (it *TargetCoreExCrossChainResponseSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TargetCoreExCrossChainResponseSent)
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
		it.Event = new(TargetCoreExCrossChainResponseSent)
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
func (it *TargetCoreExCrossChainResponseSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TargetCoreExCrossChainResponseSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TargetCoreExCrossChainResponseSent represents a CrossChainResponseSent event raised by the TargetCoreEx contract.
type TargetCoreExCrossChainResponseSent struct {
	RequestID [32]byte
	Result    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCrossChainResponseSent is a free log retrieval operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _RequestID, bytes _result)
func (_TargetCoreEx *TargetCoreExFilterer) FilterCrossChainResponseSent(opts *bind.FilterOpts) (*TargetCoreExCrossChainResponseSentIterator, error) {

	logs, sub, err := _TargetCoreEx.contract.FilterLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return &TargetCoreExCrossChainResponseSentIterator{contract: _TargetCoreEx.contract, event: "CrossChainResponseSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainResponseSent is a free log subscription operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _RequestID, bytes _result)
func (_TargetCoreEx *TargetCoreExFilterer) WatchCrossChainResponseSent(opts *bind.WatchOpts, sink chan<- *TargetCoreExCrossChainResponseSent) (event.Subscription, error) {

	logs, sub, err := _TargetCoreEx.contract.WatchLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TargetCoreExCrossChainResponseSent)
				if err := _TargetCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
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

// ParseCrossChainResponseSent is a log parse operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _RequestID, bytes _result)
func (_TargetCoreEx *TargetCoreExFilterer) ParseCrossChainResponseSent(log types.Log) (*TargetCoreExCrossChainResponseSent, error) {
	event := new(TargetCoreExCrossChainResponseSent)
	if err := _TargetCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
		return nil, err
	}
	return event, nil
}
