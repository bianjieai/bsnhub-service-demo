package bcos

import (
	"fmt"

	bcosclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"

	"github.com/bianjieai/bsnhub-service-demo/common"
)

// BCOSChain defines the BCOS chain
type BCOSChain struct {
	Client *bcosclient.Client
}

// NewBCOSChain constructs a new BCOSChain instance
func NewBCOSChain(
	configFile string,
) BCOSChain {
	configs, err := conf.ParseConfigFile(configFile)
	if err != nil {
		common.Logger.Panicf("failed to parse bcos config file: %s", err)
	}

	client, err := bcosclient.Dial(&configs[0])
	if err != nil {
		common.Logger.Panicf("failed to connect to bcos node: %s", err)
	}

	bcos := BCOSChain{
		Client: client,
	}

	return bcos
}

// MakeBCOSChain builds a BCOS chain from the given config
func MakeBCOSChain(config Config) BCOSChain {
	return NewBCOSChain(
		config.ConfigFile,
	)
}

// WaitForReceipt waits for the receipt of the given tx
func (b BCOSChain) WaitForReceipt(tx *types.Transaction, name string) (*types.Receipt, error) {
	common.Logger.Infof("%s: transaction sent, hash: %s", name, tx.Hash().Hex())

	receipt, err := b.Client.WaitMined(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to mint the transaction %s: %s", tx.Hash().Hex(), err)
	}

	if receipt.Status != types.Success {
		return nil, fmt.Errorf("transaction %s execution failed: %s", tx.Hash().Hex(), receipt.GetErrorMessage())
	}

	common.Logger.Infof("%s: transaction %s execution succeeded", name, tx.Hash().Hex())

	return receipt, nil
}
