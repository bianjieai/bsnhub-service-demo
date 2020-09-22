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

	NftCoreAddr string // nft Core Extension contract address

	NftCoreABI      abi.ABI // nft Core Extension contract ABI
	NftCoreContract *NFT    // iService Core Extension contract
}

// NewEthChain constructs a new EthChain instance
func NewEthChain(
	chainID string,
	nodeRPCAddr string,
	gasLimit uint64,
	gasPrice uint64,
	key string,
	passphrase string,
	nftCoreAddr string,
) EthChain {
	nftCoreABI, err := ParseABI(NFTABI)
	if err != nil {
		common.Logger.Panicf("failed to parse iservice core abi: %s", err)
	}

	client, err := ethclient.Dial(nodeRPCAddr)
	if err != nil {
		common.Logger.Panicf("failed to connect to %s: %s", nodeRPCAddr, err)
	}

	nftContract, err := NewNFT(ethcmn.HexToAddress(nftCoreAddr), client)
	if err != nil {
		common.Logger.Panicf("failed to instantiate the iservice market contract: %s", err)
	}

	eth := EthChain{
		ChainID:         chainID,
		NodeRPCAddr:     nodeRPCAddr,
		Client:          client,
		GasLimit:        gasLimit,
		GasPrice:        big.NewInt(int64(gasPrice)),
		Key:             key,
		Passphrase:      passphrase,
		NftCoreAddr:     nftCoreAddr,
		NftCoreABI:      nftCoreABI,
		NftCoreContract: nftContract,
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
		config.NftCoreAddr,
	)
}

// GetChainID implements AppChainI
func (ec EthChain) GetChainID() string {
	return ec.ChainID
}

// MintNft mint a nft on ethereum
func (ec EthChain) MintNft(to ethcmn.Address, amountToMint *big.Int, metaId string, setPrice *big.Int, isForSale bool) (string, error) {
	auth, err := ec.buildAuthTransactor()
	if err != nil {
		return "", err
	}

	tx, err := ec.NftCoreContract.BatchMint(auth, to, amountToMint, metaId, setPrice, isForSale)
	if err != nil {
		return "", fmt.Errorf("failed to send MintWithTokenURI transaction: %s", err)
	}

	nftId, err := ec.waitForReceipt(tx, "MintWithTokenURI")
	if err != nil {
		return "", err
	}

	return nftId, nil
}

// buildAuthTransactor builds an authenticated transactor
func (ec EthChain) buildAuthTransactor() (*bind.TransactOpts, error) {
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

// waitForReceipt waits for the receipt of the given tx
func (ec EthChain) waitForReceipt(tx *ethtypes.Transaction, name string) (string, error) {
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
		err := ec.NftCoreABI.Unpack(&nftMinted, "Minted", log.Data)
		if err == nil {
			return nftMinted, nil
		}
	}
	return nftMinted, errors.New("can not find NFTMinted event in logs")
}
