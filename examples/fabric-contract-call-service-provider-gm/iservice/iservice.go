package iservice

import (
	"bsn-irita-fabric-provider/common"
	"encoding/json"

	"bsn-irita-fabric-provider/appchains/fabric/entity"
	txStore "bsn-irita-fabric-provider/appchains/fabric/store"
	"bsn-irita-fabric-provider/types"
	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	sdk "github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
)

const (
	eventTypeNewBatchRequestProvider = "new_batch_request_provider"
	attributeKeyServiceName          = "service_name"
	attributeKeyProvider             = "provider"
	attributeKeyRequests             = "requests"
	attributeKeyRequestID            = "request_id"
)

// ServiceClientWrapper defines a wrapper for service client
type ServiceClientWrapper struct {
	ChainID     string
	NodeRPCAddr string

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

	fee, err := sdk.ParseDecCoins(defaultFee)
	if err != nil {
		panic(err)
	}

	config := sdk.ClientConfig{
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
	baseTx := s.BuildBaseTx()
	provider, e := s.ServiceClient.QueryAddress(baseTx.From, baseTx.Password)
	if e != nil {
		return sdk.Wrap(e)
	}

	builder := sdk.NewEventQueryBuilder().AddCondition(
		sdk.NewCond(eventTypeNewBatchRequestProvider, attributeKeyServiceName).EQ(sdk.EventValue(serviceName)),
	).AddCondition(
		sdk.NewCond(eventTypeNewBatchRequestProvider, attributeKeyProvider).EQ(sdk.EventValue(provider.String())),
	)

	_, err := s.ServiceClient.SubscribeNewBlock(builder, func(block sdk.EventDataNewBlock) {
		msgs := s.GenServiceResponseMsgs(block.ResultEndBlock.Events, serviceName, provider, cb)
		if msgs == nil || len(msgs) == 0 {
			s.ServiceClient.Logger().Error("no message created",
				"serviceName", serviceName,
				"provider", provider,
			)
		}
		for _, msg := range msgs {

			msg, ok := msg.(*service.MsgRespondService)

			// Submit to hub
			resTx, err := s.ServiceClient.BuildAndSend([]sdk.Msg{msg}, baseTx)

			if err != nil {
				if ok {
					// TODO 交易记录 向HUB返回响应信息的交易
					// error     err.Error() err字段
					// tx_status    2

					//mysql.TxErrCollection(msg.RequestId, err.Error())
					data := entity.CrossChainInfo{
						Tx_status:      2,
						Error:          err.Error(),
						Ic_request_id:  msg.RequestId,
						Source_service: 1,
					}
					txStore.RequestHubInfoFail(&data)
				}

				s.ServiceClient.Logger().Error("provider respond failed", "errMsg", err.Error())

			} else {

				if ok {
					// TODO 交易记录 向HUB返回响应信息的交易
					// hub_res_tx  resTx.Hash
					// tx_status    1
					//mysql.OnInterchainResponseSent(msg.RequestId, resTx.Hash)
					data := entity.CrossChainInfo{
						Tx_status:      1,
						Hub_res_tx:     resTx.Hash,
						Ic_request_id:  msg.RequestId,
						Source_service: 1,
					}
					txStore.RequestHubInfoSucc(&data)

				}

			}

		}
	})
	return err
}

func (s ServiceClientWrapper) GenServiceResponseMsgs(events sdk.StringEvents, serviceName string,
	provider sdk.AccAddress,
	handler service.RespondCallback) (msgs []sdk.Msg) {

	var ids []string
	for _, e := range events {
		if e.Type != eventTypeNewBatchRequestProvider {
			continue
		}
		attributes := sdk.Attributes(e.Attributes)
		svcName := attributes.GetValue(attributeKeyServiceName)
		prov := attributes.GetValue(attributeKeyProvider)
		if svcName == serviceName && prov == provider.String() {
			reqIDsStr := attributes.GetValue(attributeKeyRequests)
			var idsTemp []string
			if err := json.Unmarshal([]byte(reqIDsStr), &idsTemp); err != nil {
				s.ServiceClient.Logger().Error(
					"service request don't exist",
					attributeKeyRequestID, reqIDsStr,
					attributeKeyServiceName, serviceName,
					attributeKeyProvider, provider.String(),
					"errMsg", err.Error(),
				)
				return
			}
			ids = append(ids, idsTemp...)
		}
	}

	for _, reqID := range ids {
		request, err := s.ServiceClient.QueryServiceRequest(reqID)
		if err != nil {
			s.ServiceClient.Logger().Error(
				"service request don't exist",
				attributeKeyRequestID, reqID,
				attributeKeyServiceName, serviceName,
				attributeKeyProvider, provider.String(),
				"errMsg", err.Error(),
			)
			continue
		}
		//check again
		providerStr := provider.String()
		if providerStr == request.Provider && request.ServiceName == serviceName {
			output, result := handler(request.RequestContextID, reqID, request.Input)
			var resultObj types.Result
			json.Unmarshal([]byte(result), &resultObj)

			common.Logger.Infof("GenServiceResponseMsgs resultObj.Code:%d", resultObj.Code)
			if resultObj.Code != 204 {
				msgs = append(msgs, &service.MsgRespondService{
					RequestId: reqID,
					Provider:  providerStr,
					Output:    output,
					Result:    result,
				})
			}
		}
	}
	return msgs
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
	depositCoins, err := sdk.ParseDecCoins(deposit)
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
func (s ServiceClientWrapper) BuildBaseTx() sdk.BaseTx {
	return sdk.BaseTx{
		From:     s.KeyName,
		Password: s.Passphrase,
	}
}
