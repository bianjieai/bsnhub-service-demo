package iservice

import (
	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
)

// ServiceClientWrapper defines a wrapper for service client
type ServiceClientWrapper struct {
	ChainID      string
	NodeRPCAddr  string
	NodeGRPCAddr string

	KeyPath    string
	KeyName    string
	Passphrase string

	ServiceClient servicesdk.ServiceClient
}

// NewServiceClientWrapper constructs a new ServiceClientWrapper
func NewServiceClientWrapper(
	chainID string,
	nodeRPCAddr string,
	nodeGRPCAddr string,
	keyPath string,
	keyName string,
	passphrase string,
) ServiceClientWrapper {
	if len(chainID) == 0 {
		chainID = defaultChainID
	}

	if len(nodeRPCAddr) == 0 {
		nodeRPCAddr = defaultNodeRPCAddr
	}

	if len(nodeGRPCAddr) == 0 {
		nodeGRPCAddr = defaultNodeGRPCAddr
	}

	if len(keyPath) == 0 {
		keyPath = defaultKeyPath
	}

	fee, err := types.ParseDecCoins(defaultFee)
	if err != nil {
		panic(err)
	}

	config := types.ClientConfig{
		NodeURI:  nodeRPCAddr,
		GRPCAddr: nodeGRPCAddr,
		ChainID:  chainID,
		Gas:      defaultGas,
		Fee:      fee,
		Mode:     defaultBroadcastMode,
		Algo:     defaultKeyAlgorithm,
		KeyDAO:   store.NewFileDAO(keyPath),
	}

	wrapper := ServiceClientWrapper{
		ChainID:       chainID,
		NodeRPCAddr:   nodeRPCAddr,
		NodeGRPCAddr:  nodeGRPCAddr,
		KeyPath:       keyPath,
		KeyName:       keyName,
		Passphrase:    passphrase,
		ServiceClient: servicesdk.NewServiceClient(config),
	}

	return wrapper
}

// MakeServiceClientWrapper builds a ServiceClientWrapper from the given config
func MakeServiceClientWrapper(config Config) ServiceClientWrapper {
	return NewServiceClientWrapper(
		config.ChainID,
		config.NodeRPCAddr,
		config.NodeGRPCAddr,
		config.KeyPath,
		config.KeyName,
		config.Passphrase,
	)
}

// SubscribeServiceRequest wraps service.SubscribeServiceRequest
func (s ServiceClientWrapper) SubscribeServiceRequest(serviceName string, cb service.RespondCallback) error {
	_, err := s.ServiceClient.SubscribeServiceRequest(serviceName, cb, s.BuildBaseTx())
	return err
}

// DefineService wraps iservice.DefineService
func (s ServiceClientWrapper) DefineService(
	serviceName string,
	description string,
	authorDescription string,
	tags []string,
	schemas string,
) error {
	request := service.DefineServiceRequest{
		ServiceName:       serviceName,
		Description:       description,
		AuthorDescription: authorDescription,
		Tags:              tags,
		Schemas:           schemas,
	}

	_, err := s.ServiceClient.DefineService(request, s.BuildBaseTx())

	return err
}

// BindService wraps iservice.BindService
func (s ServiceClientWrapper) BindService(
	serviceName string,
	deposit string,
	pricing string,
	options string,
	qos uint64,
) error {
	depositCoins, err := types.ParseDecCoins(deposit)
	if err != nil {
		return err
	}

	provider, err := s.ShowKey(s.KeyName, s.Passphrase)
	if err != nil {
		return err
	}

	request := service.BindServiceRequest{
		ServiceName: serviceName,
		Deposit:     depositCoins,
		Pricing:     pricing,
		Options:     options,
		QoS:         qos,
		Provider:    provider,
	}

	_, err = s.ServiceClient.BindService(request, s.BuildBaseTx())

	return err
}

// BuildBaseTx builds a base tx
func (s ServiceClientWrapper) BuildBaseTx() types.BaseTx {
	return types.BaseTx{
		From:     s.KeyName,
		Password: s.Passphrase,
	}
}
