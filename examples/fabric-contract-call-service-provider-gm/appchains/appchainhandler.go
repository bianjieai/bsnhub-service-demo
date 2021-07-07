package appchains

import (
	"bsn-irita-fabric-provider/appchains/fabric"
	"bsn-irita-fabric-provider/iservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewAppChainHandler(serviceName string, v *viper.Viper, iserviceClient *iservice.ServiceClientWrapper, logger *log.Logger) AppChainHandlerI {
	return fabric.NewFabricChainHandle(serviceName, v, iserviceClient, logger)
}

type AppChainHandlerI interface {
	Start() error

	RegisterChain([]byte) error

	UpdateChain([]byte) error

	DeleteChain(chainId string) error

	GetChains() error

	Callback(reqCtxID, reqID, input string) (output string, result string)

	DeployIService(schemas string, pricing string) error
}
