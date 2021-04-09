package fisco

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	fiscocfg "github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco/config"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/server"
	"math/big"

	"github.com/ethereum/go-ethereum"
	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	fiscoclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
)

// FISCOChain defines the FISCO chain
type FISCOChain struct {
	Client       *fiscoclient.Client
	ChainManager *server.ChainManager
	BaseConfig   fiscocfg.BaseConfig
}

// NewFISCOChain constructs a new FISCOChain instance
func NewFISCOChain(
	baseConfig fiscocfg.BaseConfig,
	chainManager *server.ChainManager,
) *FISCOChain {
	return &FISCOChain{
		BaseConfig:   baseConfig,
		ChainManager: chainManager,
	}
}

// InstantiateClient instantiates the fisco client according to the given chain params
func (f *FISCOChain) InstantiateClient(
	groupID int,
	chainID int64,
) error {
	chainParams, err := f.ChainManager.GetChainParams(chainID, groupID)
	if err != nil {
		return fmt.Errorf("chainparams not exist!", err)
	}
	config := fiscocfg.Config{
		BaseConfig:  f.BaseConfig,
		ChainParams: chainParams,
	}

	clientConfig := fiscocfg.BuildClientConfig(config)

	client, err := fiscoclient.Dial(clientConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to fisco node: %s", err)
	}

	f.Client = client
	return nil
}

// CallContract calls the specified contract with the given contract address and data
func (f *FISCOChain) CallContract(
	contractAddress ethcmn.Address,
	data []byte,
) (result []byte, err error) {
	if f.Client == nil {
		return nil, fmt.Errorf("client is not instantiated")
	}

	opts := f.Client.GetCallOpts()

	msg := ethereum.CallMsg{
		From: opts.From,
		To:   &contractAddress,
		Data: data,
	}

	return f.Client.CallContract(context.Background(), msg, nil)
}

// SendContractTx sends the specified contract transaction
func (f *FISCOChain) SendContractTx(
	contractAddress ethcmn.Address,
	data []byte,
) (status bool, txHash string, err error) {
	if f.Client == nil {
		return false, "", fmt.Errorf("client is not instantiated")
	}

	opts := f.Client.GetTransactOpts()

	signedTx, err := f.generateSignedTx(opts, &contractAddress, data)
	if err != nil {
		return false, "", err
	}

	_, err = f.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, "", err
	}

	_, err = f.WaitForReceipt(signedTx, hex.EncodeToString(data[0:4]))
	if err != nil {
		return false, signedTx.Hash().Hex(), err
	}

	return true, signedTx.Hash().Hex(), nil
}

// WaitForReceipt waits for the receipt of the given tx
func (f *FISCOChain) WaitForReceipt(tx *types.Transaction, name string) (*types.Receipt, error) {
	common.Logger.Infof("%s: transaction sent, hash: %s", name, tx.Hash().Hex())

	receipt, err := f.Client.WaitMined(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to mint the transaction %s: %s", tx.Hash().Hex(), err)
	}

	if receipt.Status != types.Success {
		return nil, fmt.Errorf("transaction %s execution failed: %s", tx.Hash().Hex(), receipt.GetErrorMessage())
	}

	common.Logger.Infof("%s: transaction %s execution succeeded", name, tx.Hash().Hex())

	return receipt, nil
}

func (f *FISCOChain) generateSignedTx(opts *bind.TransactOpts, contractAddress *ethcmn.Address, input []byte) (*types.Transaction, error) {
	var err error

	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}

	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(250), nil).Sub(max, big.NewInt(1))

	nonce, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %v", err)
	}

	gasPrice := opts.GasPrice
	if gasPrice == nil {
		gasPrice = big.NewInt(30000000)
	}

	gasLimit := opts.GasLimit
	if gasLimit == nil {
		if contractAddress != nil {
			if code, err := f.Client.PendingCodeAt(context.Background(), *contractAddress); err != nil {
				return nil, err
			} else if len(code) == 0 {
				return nil, fmt.Errorf("no code for address %v", contractAddress)
			}
		}

		gasLimit = big.NewInt(30000000)
	}

	var blockLimit *big.Int
	blockLimit, err = f.Client.GetBlockLimit(context.Background())
	if err != nil {
		return nil, err
	}

	var chainID *big.Int
	chainID, err = f.Client.GetChainID(context.Background())
	if err != nil {
		return nil, err
	}

	var groupID *big.Int
	groupID = f.Client.GetGroupID()
	if groupID == nil {
		return nil, fmt.Errorf("failed to get the group ID")
	}

	var rawTx *types.Transaction
	str := ""
	extraData := []byte(str)
	if contractAddress == nil {
		rawTx = types.NewContractCreation(nonce, value, gasLimit, gasPrice, blockLimit, input, chainID, groupID, extraData, f.Client.SMCrypto())
	} else {
		rawTx = types.NewTransaction(nonce, *contractAddress, value, gasLimit, gasPrice, blockLimit, input, chainID, groupID, extraData, f.Client.SMCrypto())
	}

	if opts.Signer == nil {
		return nil, errors.New("no signer to authorize the transaction with")
	}

	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
