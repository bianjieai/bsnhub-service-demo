// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcos

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"store\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x608060405234801561001057600080fd5b50610400806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80631dd6728f1461003b5780633590b49f146100e2575b600080fd5b6100676004803603602081101561005157600080fd5b810190808035906020019092919050505061015b565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100a757808201518184015260208101905061008c565b50505050905090810190601f1680156100d45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610159600480360360208110156100f857600080fd5b810190808035906020019064010000000081111561011557600080fd5b82018360208201111561012757600080fd5b8035906020019184600183028401116401000000008311171561014957600080fd5b909192939192939050505061020b565b005b60006020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102035780601f106101d857610100808354040283529160200191610203565b820191906000526020600020905b8154815290600101906020018083116101e657829003601f168201915b505050505081565b60006102156102a9565b905082826000808481526020019081526020016000209190610238929190610325565b507fdb95fb276e0954c43831fb01b85ce465e414ac32625ec7494ed7fc651857425a81848460405180848152602001806020018281038252848482818152602001925080828437600081840152601f19601f82011690508083019250505094505050505060405180910390a1505050565b60008030600154604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1660601b815260140182815260200192505050604051602081830303815290604052805190602001209050600180600082825401925050819055508091505090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061036657803560ff1916838001178555610394565b82800160010185558215610394579182015b82811115610393578235825591602001919060010190610378565b5b5090506103a191906103a5565b5090565b6103c791905b808211156103c35760008160009055506001016103ab565b5090565b9056fea2646970667358221220cc9efada0e20db004ecbc962d8ede83a64374d5e818db8896dcd8d58a776f3ac64736f6c634300060a0033"

// DeployStore deploys a new contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

func AsyncDeployStore(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Store is an auto generated Go binding around a Solidity contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around a Solidity contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around a Solidity contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around a Solidity contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Store is a free data retrieval call binding the contract method 0x1dd6728f.
//
// Solidity: function store(bytes32 ) constant returns(string)
func (_Store *StoreCaller) Store(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "store", arg0)
	return *ret0, err
}

// Store is a free data retrieval call binding the contract method 0x1dd6728f.
//
// Solidity: function store(bytes32 ) constant returns(string)
func (_Store *StoreSession) Store(arg0 [32]byte) (string, error) {
	return _Store.Contract.Store(&_Store.CallOpts, arg0)
}

// Store is a free data retrieval call binding the contract method 0x1dd6728f.
//
// Solidity: function store(bytes32 ) constant returns(string)
func (_Store *StoreCallerSession) Store(arg0 [32]byte) (string, error) {
	return _Store.Contract.Store(&_Store.CallOpts, arg0)
}

// Set is a paid mutator transaction binding the contract method 0x3590b49f.
//
// Solidity: function set(string value) returns()
func (_Store *StoreTransactor) Set(opts *bind.TransactOpts, value string) (*types.Transaction, *types.Receipt, error) {
	return _Store.contract.Transact(opts, "set", value)
}

func (_Store *StoreTransactor) AsyncSet(handler func(*types.Receipt, error), opts *bind.TransactOpts, value string) (*types.Transaction, error) {
	return _Store.contract.AsyncTransact(opts, handler, "set", value)
}

// Set is a paid mutator transaction binding the contract method 0x3590b49f.
//
// Solidity: function set(string value) returns()
func (_Store *StoreSession) Set(value string) (*types.Transaction, *types.Receipt, error) {
	return _Store.Contract.Set(&_Store.TransactOpts, value)
}

func (_Store *StoreSession) AsyncSet(handler func(*types.Receipt, error), value string) (*types.Transaction, error) {
	return _Store.Contract.AsyncSet(handler, &_Store.TransactOpts, value)
}

// Set is a paid mutator transaction binding the contract method 0x3590b49f.
//
// Solidity: function set(string value) returns()
func (_Store *StoreTransactorSession) Set(value string) (*types.Transaction, *types.Receipt, error) {
	return _Store.Contract.Set(&_Store.TransactOpts, value)
}

func (_Store *StoreTransactorSession) AsyncSet(handler func(*types.Receipt, error), value string) (*types.Transaction, error) {
	return _Store.Contract.AsyncSet(handler, &_Store.TransactOpts, value)
}

// StoreSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Store contract.
type StoreSetIterator struct {
	Event *StoreSet // Event containing the contract specifics and raw log

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
func (it *StoreSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreSet)
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
		it.Event = new(StoreSet)
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
func (it *StoreSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreSet represents a Set event raised by the Store contract.
type StoreSet struct {
	Key   [32]byte
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000db95fb27.
//
// Solidity: event Set(bytes32 key, string value)
func (_Store *StoreFilterer) FilterSet(opts *bind.FilterOpts) (*StoreSetIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &StoreSetIterator{contract: _Store.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000db95fb27.
//
// Solidity: event Set(bytes32 key, string value)
func (_Store *StoreFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *StoreSet) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreSet)
				if err := _Store.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000db95fb27.
//
// Solidity: event Set(bytes32 key, string value)
func (_Store *StoreFilterer) ParseSet(log types.Log) (*StoreSet, error) {
	event := new(StoreSet)
	if err := _Store.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}
