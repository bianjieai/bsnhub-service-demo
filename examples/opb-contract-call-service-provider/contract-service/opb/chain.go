package opb

import (
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/common"
	opbcfg "github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/contract-service/opb/config"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/server"
	sdk "github.com/bianjieai/irita-sdk-go"
	sdktypes "github.com/bianjieai/irita-sdk-go/types"
	sdkstore "github.com/bianjieai/irita-sdk-go/types/store"
)

// OpbChain defines the Opb chain
type OpbChain struct {
	OpbClient    sdk.IRITAClient
	ChainManager *server.ChainManager
	BaseConfig   opbcfg.BaseConfig
}

// NewOpbChain constructs a new OpbChain instance
func NewOpbChain(
	baseConfig opbcfg.BaseConfig,
	chainManager *server.ChainManager,
) *OpbChain {
	return &OpbChain{
		BaseConfig:   baseConfig,
		ChainManager: chainManager,
	}
}

// BuildBaseTx builds a base tx
func (opb *OpbChain) BuildBaseTx() sdktypes.BaseTx {
	return sdktypes.BaseTx{
		From:     opb.BaseConfig.KeyName,
		Password: opb.BaseConfig.Passphrase,
		Mode:     sdktypes.Commit,
	}
}

// InstantiateClient instantiates the opb client according to the given chain params
func (f *OpbChain) InstantiateClient(
	chainParams opbcfg.ChainParams,
) error {
	config := opbcfg.Config{
		BaseConfig:  f.BaseConfig,
		ChainParams: chainParams,
	}

	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := opbcfg.RandURL(config.ChainParams.NodeURLs)
	var rpcAddr string
	var grpcAddr string
	//获取配置的nodeURL
	rpcAddrstr, ok := config.RpcAddrsMap[nodeName]
	if ok {
		rpcAddr = rpcAddrstr
	}
	grpcAddrstr, ok := config.GrpcAddrsMap[nodeName]
	if ok {
		grpcAddr = grpcAddrstr
	}

	options := []sdktypes.Option{
		sdktypes.CachedOption(true),
		sdktypes.KeyDAOOption(sdkstore.NewFileDAO(config.KeyPath)),
	}

	clientConfig, err := sdktypes.NewClientConfig(rpcAddr, grpcAddr, config.BaseConfig.ChainId, options...)

	if err != nil {
		common.Logger.Errorf("failed to get the sdk clientConfig: %s", err)
		return fmt.Errorf("failed to get the sdk clientConfig: %s", err)
	}

	opbClient := sdk.NewIRITAClient(clientConfig)
	f.OpbClient = opbClient
	return nil
}

// waitForSuccess waits for the receipt of the given tx
func (opb *OpbChain) WaitForSuccess(txHash string, name string) error {
	common.Logger.Infof("%s: transaction sent to %s, hash: %s", name, opb.BaseConfig.ChainId, txHash)

	tx, _ := opb.OpbClient.QueryTx(txHash)
	if tx.Result.Code != 0 {
		return fmt.Errorf("transaction %s execution failed: %s", txHash, tx.Result.Log)
	}

	common.Logger.Infof("%s: transaction %s execution succeeded", name, txHash)

	return nil
}
