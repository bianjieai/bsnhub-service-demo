package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bianjieai/bsnhub-service-demo/common"
)

// EthChain defines the Ethereum chain
type EthChain struct {
	ChainID     string
	NodeRPCAddr string
	Client      *ethclient.Client
	GasLimit    uint64
	GasPrice    *big.Int
	Key         string
	Passphrase  string

	NftContractAddr string // nft contract address

	NftContractABI abi.ABI // nft contract ABI
	NftContract    *NFT    // nft contract
}

// NewEthChain constructs a new EthChain instance
func NewEthChain(
	chainID string,
	nodeRPCAddr string,
	gasLimit uint64,
	gasPrice uint64,
	key string,
	passphrase string,
	nftContractAddr string,
) EthChain {
	nftContractABI, err := ParseABI(NFTABI)
	if err != nil {
		common.Logger.Panicf("failed to parse nft contract abi: %s", err)
	}

	client, err := ethclient.Dial(nodeRPCAddr)
	if err != nil {
		common.Logger.Panicf("failed to connect to %s: %s", nodeRPCAddr, err)
	}

	nftContract, err := NewNFT(ethcmn.HexToAddress(nftContractAddr), client)
	if err != nil {
		common.Logger.Panicf("failed to instantiate the nft contract: %s", err)
	}

	eth := EthChain{
		ChainID:         chainID,
		NodeRPCAddr:     nodeRPCAddr,
		Client:          client,
		GasLimit:        gasLimit,
		GasPrice:        big.NewInt(int64(gasPrice)),
		Key:             key,
		Passphrase:      passphrase,
		NftContractAddr: nftContractAddr,
		NftContractABI:  nftContractABI,
		NftContract:     nftContract,
	}

	return eth
}

// MakeEthChain builds an Ethereum chain from the given config
func MakeEthChain(config Config) EthChain {
	return NewEthChain(
		config.ChainID,
		config.NodeRPCAddr,
		config.GasLimit,
		config.GasPrice,
		config.Key,
		config.Passphrase,
		config.NftContractAddr,
	)
}

// GetChainID implements AppChainI
func (ec EthChain) GetChainID() string {
	return ec.ChainID
}

// BuildAuthTransactor builds an authenticated transactor
func (ec EthChain) BuildAuthTransactor() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(ec.Key)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)

	nextNonce, err := ec.Client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}

	auth.GasLimit = ec.GasLimit
	auth.GasPrice = ec.GasPrice
	auth.Nonce = big.NewInt(int64(nextNonce))

	return auth, nil
}

// logListener listens to the log sent by the given channel and handles it with the specified handler
func (ec EthChain) logListener(sub ethereum.Subscription, logChan chan ethtypes.Log, handler func(log ethtypes.Log)) {
	for {
		select {
		case log := <-logChan:
			handler(log)
		case err := <-sub.Err():
			common.Logger.Errorf("Error on log subscription: %s", err)
		}
	}
}

// WaitForReceipt waits for the receipt of the given tx
func (ec EthChain) WaitForReceipt(tx *ethtypes.Transaction, name string) (string, error) {
	common.Logger.Infof("%s: transaction sent to %s, hash: %s", name, ec.GetChainID(), tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		return "", fmt.Errorf("%s: failed to mint the transaction %s: %s", name, tx.Hash().Hex(), err)
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return "", fmt.Errorf("%s: transaction %s execution failed", name, tx.Hash().Hex())
	}

	nftMinted, err := ec.parseLogs(receipt.Logs)
	if err != nil {
		return "", err
	}

	return nftMinted.Id.String(), nil
}

// parseLogs parses the given logs to NFTMinted
func (ec EthChain) parseLogs(logs []*ethtypes.Log) (NFTMinted, error) {
	var nftMinted NFTMinted

	for _, log := range logs {
		err := ec.NftContractABI.Unpack(&nftMinted, "Minted", log.Data)
		if err == nil {
			return nftMinted, nil
		}
	}

	return nftMinted, errors.New("can not find NFTMinted event in logs")
}
