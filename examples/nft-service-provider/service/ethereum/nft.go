// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// NFTABI is the input ABI used to generate the binding from.
const NFTABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"giver\",\"type\":\"address\"},{\"name\":\"recipients\",\"type\":\"address[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroyAndSend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintWithTokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"buyThing\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amountToMint\",\"type\":\"uint256\"},{\"name\":\"metaId\",\"type\":\"string\"},{\"name\":\"setPrice\",\"type\":\"uint256\"},{\"name\":\"isForSale\",\"type\":\"bool\"}],\"name\":\"batchMint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"items\",\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"metaId\",\"type\":\"string\"},{\"name\":\"state\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"setTokenState\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchBurn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"setPrice\",\"type\":\"uint256\"}],\"name\":\"setTokenPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"fee\",\"type\":\"address\"},{\"name\":\"creator\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"error\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ErrorOut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"metaId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"recipients\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"BatchTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metaId\",\"type\":\"string\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"metaId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"BatchBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"metaId\",\"type\":\"string\"}],\"name\":\"BatchForSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"metaId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Destroy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"}]"

// NFT is an auto generated Go binding around an Ethereum contract.
type NFT struct {
	NFTCaller     // Read-only binding to the contract
	NFTTransactor // Write-only binding to the contract
	NFTFilterer   // Log filterer for contract events
}

// NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTSession struct {
	Contract     *NFT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCallerSession struct {
	Contract *NFTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTTransactorSession struct {
	Contract     *NFTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTRaw struct {
	Contract *NFT // Generic contract binding to access the raw methods on
}

// NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCallerRaw struct {
	Contract *NFTCaller // Generic read-only contract binding to access the raw methods on
}

// NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTTransactorRaw struct {
	Contract *NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFT creates a new instance of NFT, bound to a specific deployed contract.
func NewNFT(address common.Address, backend bind.ContractBackend) (*NFT, error) {
	contract, err := bindNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFT{NFTCaller: NFTCaller{contract: contract}, NFTTransactor: NFTTransactor{contract: contract}, NFTFilterer: NFTFilterer{contract: contract}}, nil
}

// NewNFTCaller creates a new read-only instance of NFT, bound to a specific deployed contract.
func NewNFTCaller(address common.Address, caller bind.ContractCaller) (*NFTCaller, error) {
	contract, err := bindNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCaller{contract: contract}, nil
}

// NewNFTTransactor creates a new write-only instance of NFT, bound to a specific deployed contract.
func NewNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTTransactor, error) {
	contract, err := bindNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTTransactor{contract: contract}, nil
}

// NewNFTFilterer creates a new log filterer instance of NFT, bound to a specific deployed contract.
func NewNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTFilterer, error) {
	contract, err := bindNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTFilterer{contract: contract}, nil
}

// bindNFT binds a generic wrapper to an already deployed contract.
func bindNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTABI))
	if err != nil {
		return nil, err
	}

	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFT *NFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFT *NFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFT.Contract.BalanceOf(&_NFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NFT *NFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NFT.Contract.BalanceOf(&_NFT.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_NFT *NFTCaller) BaseUri(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "baseUri")
	return *ret0, err
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_NFT *NFTSession) BaseUri() (string, error) {
	return _NFT.Contract.BaseUri(&_NFT.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_NFT *NFTCallerSession) BaseUri() (string, error) {
	return _NFT.Contract.BaseUri(&_NFT.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFT *NFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "getApproved", tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFT *NFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.GetApproved(&_NFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NFT *NFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.GetApproved(&_NFT.CallOpts, tokenId)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_NFT *NFTCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "id")
	return *ret0, err
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_NFT *NFTSession) Id() (*big.Int, error) {
	return _NFT.Contract.Id(&_NFT.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_NFT *NFTCallerSession) Id() (*big.Int, error) {
	return _NFT.Contract.Id(&_NFT.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFT *NFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "isApprovedForAll", owner, operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFT *NFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFT.Contract.IsApprovedForAll(&_NFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NFT *NFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NFT.Contract.IsApprovedForAll(&_NFT.CallOpts, owner, operator)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_NFT *NFTCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_NFT *NFTSession) IsMinter(account common.Address) (bool, error) {
	return _NFT.Contract.IsMinter(&_NFT.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) view returns(bool)
func (_NFT *NFTCallerSession) IsMinter(account common.Address) (bool, error) {
	return _NFT.Contract.IsMinter(&_NFT.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NFT *NFTCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NFT *NFTSession) IsOwner() (bool, error) {
	return _NFT.Contract.IsOwner(&_NFT.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_NFT *NFTCallerSession) IsOwner() (bool, error) {
	return _NFT.Contract.IsOwner(&_NFT.CallOpts)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 tokenId, uint256 price, string metaId, uint8 state)
func (_NFT *NFTCaller) Items(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId *big.Int
	Price   *big.Int
	MetaId  string
	State   uint8
}, error) {
	ret := new(struct {
		TokenId *big.Int
		Price   *big.Int
		MetaId  string
		State   uint8
	})
	out := ret
	err := _NFT.contract.Call(opts, out, "items", arg0)
	return *ret, err
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 tokenId, uint256 price, string metaId, uint8 state)
func (_NFT *NFTSession) Items(arg0 *big.Int) (struct {
	TokenId *big.Int
	Price   *big.Int
	MetaId  string
	State   uint8
}, error) {
	return _NFT.Contract.Items(&_NFT.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0xbfb231d2.
//
// Solidity: function items(uint256 ) view returns(uint256 tokenId, uint256 price, string metaId, uint8 state)
func (_NFT *NFTCallerSession) Items(arg0 *big.Int) (struct {
	TokenId *big.Int
	Price   *big.Int
	MetaId  string
	State   uint8
}, error) {
	return _NFT.Contract.Items(&_NFT.CallOpts, arg0)
}

// Maker is a free data retrieval call binding the contract method 0x50655d8c.
//
// Solidity: function maker() view returns(address)
func (_NFT *NFTCaller) Maker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "maker")
	return *ret0, err
}

// Maker is a free data retrieval call binding the contract method 0x50655d8c.
//
// Solidity: function maker() view returns(address)
func (_NFT *NFTSession) Maker() (common.Address, error) {
	return _NFT.Contract.Maker(&_NFT.CallOpts)
}

// Maker is a free data retrieval call binding the contract method 0x50655d8c.
//
// Solidity: function maker() view returns(address)
func (_NFT *NFTCallerSession) Maker() (common.Address, error) {
	return _NFT.Contract.Maker(&_NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFT *NFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFT *NFTSession) Name() (string, error) {
	return _NFT.Contract.Name(&_NFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NFT *NFTCallerSession) Name() (string, error) {
	return _NFT.Contract.Name(&_NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFT *NFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFT *NFTSession) Owner() (common.Address, error) {
	return _NFT.Contract.Owner(&_NFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFT *NFTCallerSession) Owner() (common.Address, error) {
	return _NFT.Contract.Owner(&_NFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFT *NFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "ownerOf", tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFT *NFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.OwnerOf(&_NFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NFT *NFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NFT.Contract.OwnerOf(&_NFT.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "supportsInterface", interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFT.Contract.SupportsInterface(&_NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFT.Contract.SupportsInterface(&_NFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFT *NFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFT *NFTSession) Symbol() (string, error) {
	return _NFT.Contract.Symbol(&_NFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NFT *NFTCallerSession) Symbol() (string, error) {
	return _NFT.Contract.Symbol(&_NFT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NFT *NFTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "tokenByIndex", index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NFT *NFTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenByIndex(&_NFT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NFT *NFTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenByIndex(&_NFT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NFT *NFTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "tokenOfOwnerByIndex", owner, index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NFT *NFTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenOfOwnerByIndex(&_NFT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NFT *NFTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenOfOwnerByIndex(&_NFT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_NFT *NFTCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_NFT *NFTSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _NFT.Contract.TokenURI(&_NFT.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_NFT *NFTCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _NFT.Contract.TokenURI(&_NFT.CallOpts, _tokenId)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_NFT *NFTCaller) TokensOfOwner(opts *bind.CallOpts, owner common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "tokensOfOwner", owner)
	return *ret0, err
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_NFT *NFTSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _NFT.Contract.TokensOfOwner(&_NFT.CallOpts, owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address owner) view returns(uint256[])
func (_NFT *NFTCallerSession) TokensOfOwner(owner common.Address) ([]*big.Int, error) {
	return _NFT.Contract.TokensOfOwner(&_NFT.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFT *NFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _NFT.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFT *NFTSession) TotalSupply() (*big.Int, error) {
	return _NFT.Contract.TotalSupply(&_NFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NFT *NFTCallerSession) TotalSupply() (*big.Int, error) {
	return _NFT.Contract.TotalSupply(&_NFT.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_NFT *NFTTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_NFT *NFTSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.AddMinter(&_NFT.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_NFT *NFTTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.AddMinter(&_NFT.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFT *NFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFT *NFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Approve(&_NFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Approve(&_NFT.TransactOpts, to, tokenId)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xdc8e92ea.
//
// Solidity: function batchBurn(uint256[] tokenIds) returns()
func (_NFT *NFTTransactor) BatchBurn(opts *bind.TransactOpts, tokenIds []*big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "batchBurn", tokenIds)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xdc8e92ea.
//
// Solidity: function batchBurn(uint256[] tokenIds) returns()
func (_NFT *NFTSession) BatchBurn(tokenIds []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BatchBurn(&_NFT.TransactOpts, tokenIds)
}

// BatchBurn is a paid mutator transaction binding the contract method 0xdc8e92ea.
//
// Solidity: function batchBurn(uint256[] tokenIds) returns()
func (_NFT *NFTTransactorSession) BatchBurn(tokenIds []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BatchBurn(&_NFT.TransactOpts, tokenIds)
}

// BatchMint is a paid mutator transaction binding the contract method 0x98588a2b.
//
// Solidity: function batchMint(address to, uint256 amountToMint, string metaId, uint256 setPrice, bool isForSale) returns()
func (_NFT *NFTTransactor) BatchMint(opts *bind.TransactOpts, to common.Address, amountToMint *big.Int, metaId string, setPrice *big.Int, isForSale bool) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "batchMint", to, amountToMint, metaId, setPrice, isForSale)
}

// BatchMint is a paid mutator transaction binding the contract method 0x98588a2b.
//
// Solidity: function batchMint(address to, uint256 amountToMint, string metaId, uint256 setPrice, bool isForSale) returns()
func (_NFT *NFTSession) BatchMint(to common.Address, amountToMint *big.Int, metaId string, setPrice *big.Int, isForSale bool) (*types.Transaction, error) {
	return _NFT.Contract.BatchMint(&_NFT.TransactOpts, to, amountToMint, metaId, setPrice, isForSale)
}

// BatchMint is a paid mutator transaction binding the contract method 0x98588a2b.
//
// Solidity: function batchMint(address to, uint256 amountToMint, string metaId, uint256 setPrice, bool isForSale) returns()
func (_NFT *NFTTransactorSession) BatchMint(to common.Address, amountToMint *big.Int, metaId string, setPrice *big.Int, isForSale bool) (*types.Transaction, error) {
	return _NFT.Contract.BatchMint(&_NFT.TransactOpts, to, amountToMint, metaId, setPrice, isForSale)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x1239ec8c.
//
// Solidity: function batchTransfer(address giver, address[] recipients, uint256[] values) returns()
func (_NFT *NFTTransactor) BatchTransfer(opts *bind.TransactOpts, giver common.Address, recipients []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "batchTransfer", giver, recipients, values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x1239ec8c.
//
// Solidity: function batchTransfer(address giver, address[] recipients, uint256[] values) returns()
func (_NFT *NFTSession) BatchTransfer(giver common.Address, recipients []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BatchTransfer(&_NFT.TransactOpts, giver, recipients, values)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x1239ec8c.
//
// Solidity: function batchTransfer(address giver, address[] recipients, uint256[] values) returns()
func (_NFT *NFTTransactorSession) BatchTransfer(giver common.Address, recipients []common.Address, values []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BatchTransfer(&_NFT.TransactOpts, giver, recipients, values)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NFT *NFTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NFT *NFTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Burn(&_NFT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Burn(&_NFT.TransactOpts, tokenId)
}

// BuyThing is a paid mutator transaction binding the contract method 0x80482491.
//
// Solidity: function buyThing(uint256 _tokenId) payable returns(bool)
func (_NFT *NFTTransactor) BuyThing(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "buyThing", _tokenId)
}

// BuyThing is a paid mutator transaction binding the contract method 0x80482491.
//
// Solidity: function buyThing(uint256 _tokenId) payable returns(bool)
func (_NFT *NFTSession) BuyThing(_tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BuyThing(&_NFT.TransactOpts, _tokenId)
}

// BuyThing is a paid mutator transaction binding the contract method 0x80482491.
//
// Solidity: function buyThing(uint256 _tokenId) payable returns(bool)
func (_NFT *NFTTransactorSession) BuyThing(_tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BuyThing(&_NFT.TransactOpts, _tokenId)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0x45cefcec.
//
// Solidity: function destroyAndSend() returns()
func (_NFT *NFTTransactor) DestroyAndSend(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "destroyAndSend")
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0x45cefcec.
//
// Solidity: function destroyAndSend() returns()
func (_NFT *NFTSession) DestroyAndSend() (*types.Transaction, error) {
	return _NFT.Contract.DestroyAndSend(&_NFT.TransactOpts)
}

// DestroyAndSend is a paid mutator transaction binding the contract method 0x45cefcec.
//
// Solidity: function destroyAndSend() returns()
func (_NFT *NFTTransactorSession) DestroyAndSend() (*types.Transaction, error) {
	return _NFT.Contract.DestroyAndSend(&_NFT.TransactOpts)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_NFT *NFTTransactor) MintWithTokenURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "mintWithTokenURI", to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_NFT *NFTSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _NFT.Contract.MintWithTokenURI(&_NFT.TransactOpts, to, tokenId, tokenURI)
}

// MintWithTokenURI is a paid mutator transaction binding the contract method 0x50bb4e7f.
//
// Solidity: function mintWithTokenURI(address to, uint256 tokenId, string tokenURI) returns(bool)
func (_NFT *NFTTransactorSession) MintWithTokenURI(to common.Address, tokenId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _NFT.Contract.MintWithTokenURI(&_NFT.TransactOpts, to, tokenId, tokenURI)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_NFT *NFTTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_NFT *NFTSession) RenounceMinter() (*types.Transaction, error) {
	return _NFT.Contract.RenounceMinter(&_NFT.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_NFT *NFTTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _NFT.Contract.RenounceMinter(&_NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFT *NFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFT *NFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFT.Contract.RenounceOwnership(&_NFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFT *NFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFT.Contract.RenounceOwnership(&_NFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFT *NFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFT *NFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom0(&_NFT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NFT *NFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom0(&_NFT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_NFT *NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, to common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setApprovalForAll", to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_NFT *NFTSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.Contract.SetApprovalForAll(&_NFT.TransactOpts, to, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address to, bool approved) returns()
func (_NFT *NFTTransactorSession) SetApprovalForAll(to common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.Contract.SetApprovalForAll(&_NFT.TransactOpts, to, approved)
}

// SetTokenPrice is a paid mutator transaction binding the contract method 0xe564da69.
//
// Solidity: function setTokenPrice(uint256[] ids, uint256 setPrice) returns()
func (_NFT *NFTTransactor) SetTokenPrice(opts *bind.TransactOpts, ids []*big.Int, setPrice *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setTokenPrice", ids, setPrice)
}

// SetTokenPrice is a paid mutator transaction binding the contract method 0xe564da69.
//
// Solidity: function setTokenPrice(uint256[] ids, uint256 setPrice) returns()
func (_NFT *NFTSession) SetTokenPrice(ids []*big.Int, setPrice *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SetTokenPrice(&_NFT.TransactOpts, ids, setPrice)
}

// SetTokenPrice is a paid mutator transaction binding the contract method 0xe564da69.
//
// Solidity: function setTokenPrice(uint256[] ids, uint256 setPrice) returns()
func (_NFT *NFTTransactorSession) SetTokenPrice(ids []*big.Int, setPrice *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SetTokenPrice(&_NFT.TransactOpts, ids, setPrice)
}

// SetTokenState is a paid mutator transaction binding the contract method 0xd4477b3c.
//
// Solidity: function setTokenState(uint256[] ids, bool isEnabled) returns()
func (_NFT *NFTTransactor) SetTokenState(opts *bind.TransactOpts, ids []*big.Int, isEnabled bool) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setTokenState", ids, isEnabled)
}

// SetTokenState is a paid mutator transaction binding the contract method 0xd4477b3c.
//
// Solidity: function setTokenState(uint256[] ids, bool isEnabled) returns()
func (_NFT *NFTSession) SetTokenState(ids []*big.Int, isEnabled bool) (*types.Transaction, error) {
	return _NFT.Contract.SetTokenState(&_NFT.TransactOpts, ids, isEnabled)
}

// SetTokenState is a paid mutator transaction binding the contract method 0xd4477b3c.
//
// Solidity: function setTokenState(uint256[] ids, bool isEnabled) returns()
func (_NFT *NFTTransactorSession) SetTokenState(ids []*big.Int, isEnabled bool) (*types.Transaction, error) {
	return _NFT.Contract.SetTokenState(&_NFT.TransactOpts, ids, isEnabled)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.TransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NFT *NFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.TransferFrom(&_NFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFT *NFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFT *NFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFT.Contract.TransferOwnership(&_NFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFT *NFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFT.Contract.TransferOwnership(&_NFT.TransactOpts, newOwner)
}

// NFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NFT contract.
type NFTApprovalIterator struct {
	Event *NFTApproval // Event containing the contract specifics and raw log

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
func (it *NFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTApproval)
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
		it.Event = new(NFTApproval)
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
func (it *NFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTApproval represents a Approval event raised by the NFT contract.
type NFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFT *NFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTApprovalIterator{contract: _NFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFT *NFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTApproval)
				if err := _NFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NFT *NFTFilterer) ParseApproval(log types.Log) (*NFTApproval, error) {
	event := new(NFTApproval)
	if err := _NFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFT contract.
type NFTApprovalForAllIterator struct {
	Event *NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTApprovalForAll)
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
		it.Event = new(NFTApprovalForAll)
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
func (it *NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTApprovalForAll represents a ApprovalForAll event raised by the NFT contract.
type NFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFT *NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTApprovalForAllIterator{contract: _NFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFT *NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTApprovalForAll)
				if err := _NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NFT *NFTFilterer) ParseApprovalForAll(log types.Log) (*NFTApprovalForAll, error) {
	event := new(NFTApprovalForAll)
	if err := _NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTBatchBurnedIterator is returned from FilterBatchBurned and is used to iterate over the raw logs and unpacked data for BatchBurned events raised by the NFT contract.
type NFTBatchBurnedIterator struct {
	Event *NFTBatchBurned // Event containing the contract specifics and raw log

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
func (it *NFTBatchBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBatchBurned)
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
		it.Event = new(NFTBatchBurned)
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
func (it *NFTBatchBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBatchBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBatchBurned represents a BatchBurned event raised by the NFT contract.
type NFTBatchBurned struct {
	MetaId string
	Ids    []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchBurned is a free log retrieval operation binding the contract event 0x68f2536b9ff968bd2acc006b7cf7be7ca83c2f4a462c355e925bb354e5ce43d5.
//
// Solidity: event BatchBurned(string metaId, uint256[] ids)
func (_NFT *NFTFilterer) FilterBatchBurned(opts *bind.FilterOpts) (*NFTBatchBurnedIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "BatchBurned")
	if err != nil {
		return nil, err
	}
	return &NFTBatchBurnedIterator{contract: _NFT.contract, event: "BatchBurned", logs: logs, sub: sub}, nil
}

// WatchBatchBurned is a free log subscription operation binding the contract event 0x68f2536b9ff968bd2acc006b7cf7be7ca83c2f4a462c355e925bb354e5ce43d5.
//
// Solidity: event BatchBurned(string metaId, uint256[] ids)
func (_NFT *NFTFilterer) WatchBatchBurned(opts *bind.WatchOpts, sink chan<- *NFTBatchBurned) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "BatchBurned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBatchBurned)
				if err := _NFT.contract.UnpackLog(event, "BatchBurned", log); err != nil {
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

// ParseBatchBurned is a log parse operation binding the contract event 0x68f2536b9ff968bd2acc006b7cf7be7ca83c2f4a462c355e925bb354e5ce43d5.
//
// Solidity: event BatchBurned(string metaId, uint256[] ids)
func (_NFT *NFTFilterer) ParseBatchBurned(log types.Log) (*NFTBatchBurned, error) {
	event := new(NFTBatchBurned)
	if err := _NFT.contract.UnpackLog(event, "BatchBurned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTBatchForSaleIterator is returned from FilterBatchForSale and is used to iterate over the raw logs and unpacked data for BatchForSale events raised by the NFT contract.
type NFTBatchForSaleIterator struct {
	Event *NFTBatchForSale // Event containing the contract specifics and raw log

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
func (it *NFTBatchForSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBatchForSale)
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
		it.Event = new(NFTBatchForSale)
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
func (it *NFTBatchForSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBatchForSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBatchForSale represents a BatchForSale event raised by the NFT contract.
type NFTBatchForSale struct {
	Ids    []*big.Int
	MetaId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBatchForSale is a free log retrieval operation binding the contract event 0xefd56044628f374cec68c952a0580a76d23b4bd06d17bb887193f2176ba6c48b.
//
// Solidity: event BatchForSale(uint256[] ids, string metaId)
func (_NFT *NFTFilterer) FilterBatchForSale(opts *bind.FilterOpts) (*NFTBatchForSaleIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "BatchForSale")
	if err != nil {
		return nil, err
	}
	return &NFTBatchForSaleIterator{contract: _NFT.contract, event: "BatchForSale", logs: logs, sub: sub}, nil
}

// WatchBatchForSale is a free log subscription operation binding the contract event 0xefd56044628f374cec68c952a0580a76d23b4bd06d17bb887193f2176ba6c48b.
//
// Solidity: event BatchForSale(uint256[] ids, string metaId)
func (_NFT *NFTFilterer) WatchBatchForSale(opts *bind.WatchOpts, sink chan<- *NFTBatchForSale) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "BatchForSale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBatchForSale)
				if err := _NFT.contract.UnpackLog(event, "BatchForSale", log); err != nil {
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

// ParseBatchForSale is a log parse operation binding the contract event 0xefd56044628f374cec68c952a0580a76d23b4bd06d17bb887193f2176ba6c48b.
//
// Solidity: event BatchForSale(uint256[] ids, string metaId)
func (_NFT *NFTFilterer) ParseBatchForSale(log types.Log) (*NFTBatchForSale, error) {
	event := new(NFTBatchForSale)
	if err := _NFT.contract.UnpackLog(event, "BatchForSale", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTBatchTransferedIterator is returned from FilterBatchTransfered and is used to iterate over the raw logs and unpacked data for BatchTransfered events raised by the NFT contract.
type NFTBatchTransferedIterator struct {
	Event *NFTBatchTransfered // Event containing the contract specifics and raw log

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
func (it *NFTBatchTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBatchTransfered)
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
		it.Event = new(NFTBatchTransfered)
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
func (it *NFTBatchTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBatchTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBatchTransfered represents a BatchTransfered event raised by the NFT contract.
type NFTBatchTransfered struct {
	MetaId     string
	Recipients []common.Address
	Ids        []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfered is a free log retrieval operation binding the contract event 0xe0bdff057bc74b2545c13ac89393038d336f254b2245c876efa9669ddcc09b97.
//
// Solidity: event BatchTransfered(string metaId, address[] recipients, uint256[] ids)
func (_NFT *NFTFilterer) FilterBatchTransfered(opts *bind.FilterOpts) (*NFTBatchTransferedIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "BatchTransfered")
	if err != nil {
		return nil, err
	}
	return &NFTBatchTransferedIterator{contract: _NFT.contract, event: "BatchTransfered", logs: logs, sub: sub}, nil
}

// WatchBatchTransfered is a free log subscription operation binding the contract event 0xe0bdff057bc74b2545c13ac89393038d336f254b2245c876efa9669ddcc09b97.
//
// Solidity: event BatchTransfered(string metaId, address[] recipients, uint256[] ids)
func (_NFT *NFTFilterer) WatchBatchTransfered(opts *bind.WatchOpts, sink chan<- *NFTBatchTransfered) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "BatchTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBatchTransfered)
				if err := _NFT.contract.UnpackLog(event, "BatchTransfered", log); err != nil {
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

// ParseBatchTransfered is a log parse operation binding the contract event 0xe0bdff057bc74b2545c13ac89393038d336f254b2245c876efa9669ddcc09b97.
//
// Solidity: event BatchTransfered(string metaId, address[] recipients, uint256[] ids)
func (_NFT *NFTFilterer) ParseBatchTransfered(log types.Log) (*NFTBatchTransfered, error) {
	event := new(NFTBatchTransfered)
	if err := _NFT.contract.UnpackLog(event, "BatchTransfered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the NFT contract.
type NFTBoughtIterator struct {
	Event *NFTBought // Event containing the contract specifics and raw log

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
func (it *NFTBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBought)
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
		it.Event = new(NFTBought)
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
func (it *NFTBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBought represents a Bought event raised by the NFT contract.
type NFTBought struct {
	TokenId *big.Int
	MetaId  string
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0x46e70e87b058d8f3cde334fc179a61ac8acadcc4fdc26735e4195106b3c88eb3.
//
// Solidity: event Bought(uint256 tokenId, string metaId, uint256 value)
func (_NFT *NFTFilterer) FilterBought(opts *bind.FilterOpts) (*NFTBoughtIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return &NFTBoughtIterator{contract: _NFT.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0x46e70e87b058d8f3cde334fc179a61ac8acadcc4fdc26735e4195106b3c88eb3.
//
// Solidity: event Bought(uint256 tokenId, string metaId, uint256 value)
func (_NFT *NFTFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *NFTBought) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Bought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBought)
				if err := _NFT.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0x46e70e87b058d8f3cde334fc179a61ac8acadcc4fdc26735e4195106b3c88eb3.
//
// Solidity: event Bought(uint256 tokenId, string metaId, uint256 value)
func (_NFT *NFTFilterer) ParseBought(log types.Log) (*NFTBought, error) {
	event := new(NFTBought)
	if err := _NFT.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTDestroyIterator is returned from FilterDestroy and is used to iterate over the raw logs and unpacked data for Destroy events raised by the NFT contract.
type NFTDestroyIterator struct {
	Event *NFTDestroy // Event containing the contract specifics and raw log

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
func (it *NFTDestroyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTDestroy)
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
		it.Event = new(NFTDestroy)
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
func (it *NFTDestroyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTDestroyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTDestroy represents a Destroy event raised by the NFT contract.
type NFTDestroy struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDestroy is a free log retrieval operation binding the contract event 0xf58fef8e187ef8dfd7bef096c1ef3e4f3c54f98d95b8ad5659349b07e61204df.
//
// Solidity: event Destroy()
func (_NFT *NFTFilterer) FilterDestroy(opts *bind.FilterOpts) (*NFTDestroyIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Destroy")
	if err != nil {
		return nil, err
	}
	return &NFTDestroyIterator{contract: _NFT.contract, event: "Destroy", logs: logs, sub: sub}, nil
}

// WatchDestroy is a free log subscription operation binding the contract event 0xf58fef8e187ef8dfd7bef096c1ef3e4f3c54f98d95b8ad5659349b07e61204df.
//
// Solidity: event Destroy()
func (_NFT *NFTFilterer) WatchDestroy(opts *bind.WatchOpts, sink chan<- *NFTDestroy) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Destroy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTDestroy)
				if err := _NFT.contract.UnpackLog(event, "Destroy", log); err != nil {
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

// ParseDestroy is a log parse operation binding the contract event 0xf58fef8e187ef8dfd7bef096c1ef3e4f3c54f98d95b8ad5659349b07e61204df.
//
// Solidity: event Destroy()
func (_NFT *NFTFilterer) ParseDestroy(log types.Log) (*NFTDestroy, error) {
	event := new(NFTDestroy)
	if err := _NFT.contract.UnpackLog(event, "Destroy", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTErrorOutIterator is returned from FilterErrorOut and is used to iterate over the raw logs and unpacked data for ErrorOut events raised by the NFT contract.
type NFTErrorOutIterator struct {
	Event *NFTErrorOut // Event containing the contract specifics and raw log

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
func (it *NFTErrorOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTErrorOut)
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
		it.Event = new(NFTErrorOut)
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
func (it *NFTErrorOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTErrorOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTErrorOut represents a ErrorOut event raised by the NFT contract.
type NFTErrorOut struct {
	Error   string
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterErrorOut is a free log retrieval operation binding the contract event 0x62f7925b321cb13eb5bae75960341d045933b23884aa3bb6cf1deaecf307aec1.
//
// Solidity: event ErrorOut(string error, uint256 tokenId)
func (_NFT *NFTFilterer) FilterErrorOut(opts *bind.FilterOpts) (*NFTErrorOutIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ErrorOut")
	if err != nil {
		return nil, err
	}
	return &NFTErrorOutIterator{contract: _NFT.contract, event: "ErrorOut", logs: logs, sub: sub}, nil
}

// WatchErrorOut is a free log subscription operation binding the contract event 0x62f7925b321cb13eb5bae75960341d045933b23884aa3bb6cf1deaecf307aec1.
//
// Solidity: event ErrorOut(string error, uint256 tokenId)
func (_NFT *NFTFilterer) WatchErrorOut(opts *bind.WatchOpts, sink chan<- *NFTErrorOut) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ErrorOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTErrorOut)
				if err := _NFT.contract.UnpackLog(event, "ErrorOut", log); err != nil {
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

// ParseErrorOut is a log parse operation binding the contract event 0x62f7925b321cb13eb5bae75960341d045933b23884aa3bb6cf1deaecf307aec1.
//
// Solidity: event ErrorOut(string error, uint256 tokenId)
func (_NFT *NFTFilterer) ParseErrorOut(log types.Log) (*NFTErrorOut, error) {
	event := new(NFTErrorOut)
	if err := _NFT.contract.UnpackLog(event, "ErrorOut", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the NFT contract.
type NFTMintedIterator struct {
	Event *NFTMinted // Event containing the contract specifics and raw log

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
func (it *NFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMinted)
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
		it.Event = new(NFTMinted)
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
func (it *NFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMinted represents a Minted event raised by the NFT contract.
type NFTMinted struct {
	Id     *big.Int
	MetaId string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xadef11a3979b8ceb0573eb6ef0678134a09c23a0d94e5ea47cd18ac3a9fc0194.
//
// Solidity: event Minted(uint256 id, string metaId)
func (_NFT *NFTFilterer) FilterMinted(opts *bind.FilterOpts) (*NFTMintedIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &NFTMintedIterator{contract: _NFT.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xadef11a3979b8ceb0573eb6ef0678134a09c23a0d94e5ea47cd18ac3a9fc0194.
//
// Solidity: event Minted(uint256 id, string metaId)
func (_NFT *NFTFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *NFTMinted) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMinted)
				if err := _NFT.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xadef11a3979b8ceb0573eb6ef0678134a09c23a0d94e5ea47cd18ac3a9fc0194.
//
// Solidity: event Minted(uint256 id, string metaId)
func (_NFT *NFTFilterer) ParseMinted(log types.Log) (*NFTMinted, error) {
	event := new(NFTMinted)
	if err := _NFT.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the NFT contract.
type NFTMinterAddedIterator struct {
	Event *NFTMinterAdded // Event containing the contract specifics and raw log

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
func (it *NFTMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMinterAdded)
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
		it.Event = new(NFTMinterAdded)
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
func (it *NFTMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMinterAdded represents a MinterAdded event raised by the NFT contract.
type NFTMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_NFT *NFTFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*NFTMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &NFTMinterAddedIterator{contract: _NFT.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_NFT *NFTFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *NFTMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMinterAdded)
				if err := _NFT.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_NFT *NFTFilterer) ParseMinterAdded(log types.Log) (*NFTMinterAdded, error) {
	event := new(NFTMinterAdded)
	if err := _NFT.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the NFT contract.
type NFTMinterRemovedIterator struct {
	Event *NFTMinterRemoved // Event containing the contract specifics and raw log

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
func (it *NFTMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMinterRemoved)
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
		it.Event = new(NFTMinterRemoved)
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
func (it *NFTMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMinterRemoved represents a MinterRemoved event raised by the NFT contract.
type NFTMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_NFT *NFTFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*NFTMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &NFTMinterRemovedIterator{contract: _NFT.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_NFT *NFTFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *NFTMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMinterRemoved)
				if err := _NFT.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_NFT *NFTFilterer) ParseMinterRemoved(log types.Log) (*NFTMinterRemoved, error) {
	event := new(NFTMinterRemoved)
	if err := _NFT.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFT contract.
type NFTOwnershipTransferredIterator struct {
	Event *NFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTOwnershipTransferred)
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
		it.Event = new(NFTOwnershipTransferred)
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
func (it *NFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTOwnershipTransferred represents a OwnershipTransferred event raised by the NFT contract.
type NFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFT *NFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTOwnershipTransferredIterator{contract: _NFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFT *NFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTOwnershipTransferred)
				if err := _NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFT *NFTFilterer) ParseOwnershipTransferred(log types.Log) (*NFTOwnershipTransferred, error) {
	event := new(NFTOwnershipTransferred)
	if err := _NFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFT contract.
type NFTTransferIterator struct {
	Event *NFTTransfer // Event containing the contract specifics and raw log

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
func (it *NFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTTransfer)
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
		it.Event = new(NFTTransfer)
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
func (it *NFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTTransfer represents a Transfer event raised by the NFT contract.
type NFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFT *NFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTTransferIterator{contract: _NFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFT *NFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTTransfer)
				if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NFT *NFTFilterer) ParseTransfer(log types.Log) (*NFTTransfer, error) {
	event := new(NFTTransfer)
	if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
