package fisco

import (
	"fmt"
	fiscoclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	contract_service "github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service"
	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
	fiscocfg "github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco/config"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/server"
)

// FISCOChain defines the FISCO chain
type FISCOChain struct {
	Client       *fiscoclient.Client
	ChainManager *server.ChainManager
	BaseConfig   fiscocfg.BaseConfig
	IServiceCoreSession  *contract_service.IServiceCoreExSession   // iService Core Extension contract session
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
	chainParams fiscocfg.ChainParams,
) error {
	config := fiscocfg.Config{
		BaseConfig:  f.BaseConfig,
		ChainParams: chainParams,
	}

	clientConfig := fiscocfg.BuildClientConfig(config)

	client, err := fiscoclient.Dial(clientConfig)
	if err != nil {
		common.Logger.Errorf("failed to connect to fisco node: %s", err)
		return fmt.Errorf("failed to connect to fisco node: %s", err)
	}

	iServiceCore, err := contract_service.NewIServiceCoreEx(ethcmn.HexToAddress(f.BaseConfig.IServiceCoreAddr), client)
	if err != nil {
		common.Logger.Errorf("failed to instantiate the iservice core contract: %s", err)
	}
	f.Client = client
	f.IServiceCoreSession = &contract_service.IServiceCoreExSession{Contract: iServiceCore, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	return nil
}

// CallContract calls the specified contract with the given contract address and data

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
