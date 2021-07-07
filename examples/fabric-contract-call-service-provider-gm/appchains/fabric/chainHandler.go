package fabric

import (
	"bsn-irita-fabric-provider/appchains/fabric/config"
	"bsn-irita-fabric-provider/appchains/fabric/entity"
	"bsn-irita-fabric-provider/appchains/fabric/metadata"
	"bsn-irita-fabric-provider/appchains/fabric/store"
	"bsn-irita-fabric-provider/common"
	"bsn-irita-fabric-provider/errors"
	"bsn-irita-fabric-provider/iservice"
	"bsn-irita-fabric-provider/types"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"
)

const (
	fabric_sdk_config    = "fabric.sdk_config"
	fabric_msp_user_name = "fabric.msp_user_name"
	fabric_org_name      = "fabric.org_name"
	base_mysql_conn      = "base.mysql_conn"
	base_city_code       = "base.city_code"
)

type FabricChainHandler struct {
	IServiceClient *iservice.ServiceClientWrapper

	serviceName string
	appChain    map[string]*FabricChain

	chainInfos map[string]*entity.FabricChainInfo

	sdkConf *config.FabricConfig

	logger *log.Logger
}

func NewFabricChainHandle(serviceName string, v *viper.Viper, iserviceClient *iservice.ServiceClientWrapper, logger *log.Logger) *FabricChainHandler {

	store.InitMysql(v.GetString(base_mysql_conn))

	conf := &config.FabricConfig{
		SdkConfig:   v.GetString(fabric_sdk_config),
		MspUserName: v.GetString(fabric_msp_user_name),
		OrgName:     v.GetString(fabric_org_name),
		OrgCode:     v.GetString(base_city_code),
	}
	common.Logger.Infof("fabric config is %v", conf)

	fabric := &FabricChainHandler{
		IServiceClient: iserviceClient,
		serviceName:    serviceName,
		logger:         logger,
		appChain:       make(map[string]*FabricChain),
		sdkConf:        conf,
		chainInfos:     make(map[string]*entity.FabricChainInfo),
	}

	return fabric
}

func (f *FabricChainHandler) Start() error {

	go f.listenerStart()
	//go f.check()

	return nil
}

func (f *FabricChainHandler) check() {

	time.Sleep(time.Minute)

	for _, chain := range f.appChain {
		if chain.LastCallTime.Add(5 * time.Minute).After(time.Now()) {
			chain.Close()
			delete(f.appChain, chain.chainInfo.GetChainId())
		}
	}
}

func (f *FabricChainHandler) listenerStart() error {
	f.logger.Infof("app started")
	err := f.IServiceClient.SubscribeServiceRequest(
		f.serviceName,
		f.Callback,
	)
	if err != nil {
		f.logger.Errorf("failed to subscribe service requests, err: %s", err.Error())
		return errors.New("failed to subscribe service requests, err: %s", err.Error())
	}

	return nil
}

func (f *FabricChainHandler) RegisterChain(data []byte) error {
	regData := &entity.RegisterData{}
	err := json.Unmarshal(data, regData)
	if err != nil {
		f.logger.Error("invalid JSON %v", err)
		return errors.New("invalid JSON")
	}
	//存到数据库
	sc := regData.ToStoreDate(f.sdkConf.OrgCode)
	err = store.StoreProviderAppInfo(sc)
	if err != nil {
		//日志
		f.logger.Errorf("RegisterChain store database failed %v", err)
		return errors.New("Storage DB failed")
	}
	//把数据库的对象存到 map chainInfos
	f.chainInfos[sc.GetChainId()] = sc

	return nil
}

func (f *FabricChainHandler) UpdateChain(data []byte) error {
	regDate := &entity.RegisterData{}
	err := json.Unmarshal(data, regDate)
	if err != nil {
		f.logger.Errorf("invalid JSON %v", err)
		return errors.New("invalid JSON")
	}
	sc := regDate.ToStoreDate(f.sdkConf.OrgCode)
	//更新数据库
	updateFabricProvider := regDate.UpdateStoreDate()
	err = store.UpdateProviderAppInfo(updateFabricProvider)
	if err != nil {
		f.logger.Errorf("RegisterChain update database failed %v", err)
		return errors.New("Update DB failed")
	}
	//更新map chainInfos
	f.chainInfos[sc.GetChainId()] = updateFabricProvider

	return nil
}

func (f *FabricChainHandler) DeleteChain(chainId string) error {
	sc := chainId
	//删除数据库
	f.logger.Infof("start delete ChainId  %s", sc)
	err := store.DeleteProviderAppInfo(sc)
	if err != nil {
		f.logger.Errorf("RegisterChain delete database failed %v", err)
		return errors.New("Delete DB failed")
	}
	//删除map
	bc, err := strconv.ParseUint(chainId, 10, 64)
	if err != nil {
		f.logger.Error("Invalid character %v", err)
		return errors.New("Invalid character")
	}
	rc := entity.ChainBase{
		ChainId: bc,
	}
	delete(f.chainInfos, rc.GetChainId())
	return nil
}

func (f *FabricChainHandler) GetChains() error {
	return nil
}

func (f *FabricChainHandler) getChainInfo(chainId string) *entity.FabricChainInfo {
	chainInfo, ok := f.chainInfos[chainId]
	if ok {
		common.Logger.Infof("chainID:%s", chainId)
		common.Logger.Infof("chainInfo GetChainId:%s", chainInfo.GetChainId())
		return chainInfo
	} else {

		rc := store.GetProviderAppInfo(chainId)
		if rc != nil {
			f.chainInfos[rc.GetChainId()] = rc
			return rc
		} else {
			return nil
		}

		//_, ok := f.chainInfos[rc.GetChainId()]
		//if !ok {
		//	return nil, nil
		//} else {
		//	f.chainInfos[rc.GetChainId()] = rc
		//}
		////todo 从数据库查询
		////存在，添加到 map
		////不存在，返回nil
		//
		//return nil, nil
	}

}

func (f *FabricChainHandler) pransInput(input string) (*metadata.CrossData, error) {
	data := &metadata.CrossData{}
	err := json.Unmarshal([]byte(input), data)
	return data, err
}

func checkInput(iutput *metadata.FabricIutput) (bool, string) {

	if strings.TrimSpace(iutput.ChainCode) == "" {
		return false, types.NewResult(types.Status_Params_Error, "chaincode can not be empty")
	}

	if strings.TrimSpace(iutput.FunType) == "" {
		return false, types.NewResult(types.Status_Params_Error, "function type can not be empty")
	}
	if len(iutput.Args) == 0 {
		return false, types.NewResult(types.Status_Params_Error, "args can not be empty")
	}

	return true, ""
}

func (f *FabricChainHandler) Callback(reqCtxID, reqID, input string) (output string, result string) {
	f.logger.Infof("Callback reqCtxID:%s", reqCtxID)
	f.logger.Infof("Callback reqID:%s", reqID)
	f.logger.Infof("Callback input:%s", input)

	crossData, err := f.pransInput(input)

	if err != nil {
		return entity.GetErrOutPut(), types.NewResult(types.Status_Chain_NotExist, "invalid JSON input")
	}

	inputData := crossData.Body
	if inputData == nil {
		return entity.GetErrOutPut(), types.NewResult(types.Status_Chain_NotExist, "invalid JSON input body")
	}

	chainId := inputData.GetChainId()
	fabricChain, ok := f.appChain[chainId]

	if !ok {
		chainInfo := f.getChainInfo(chainId)
		if chainInfo == nil {
			//不处理或者处理失败
			//如果不存在该chainId 信息，需要能返回不处理的信号，hub不用返回结果
			return entity.GetErrOutPut(), types.NewResult(types.Status_Chain_NotExist, "chain not exist")
		}

		fabricChain, err = NewFabricChain(f.sdkConf, chainInfo)

		if err != nil {
			return entity.GetErrOutPut(), types.NewResult(types.Status_Error, fmt.Sprintf("call chain %s has error : %s", inputData.GetChainId(), err.Error()))
		}

		f.appChain[chainId] = fabricChain
	}

	ok, checkRes := checkInput(inputData)
	if !ok {
		return entity.GetErrOutPut(), checkRes
	}

	var res *entity.FabricRespone

	if inputData.FunType == "invoke" {
		res, err = fabricChain.Invoke(inputData.ChainCode, inputData.Args)
	} else {
		res, err = fabricChain.Query(inputData.ChainCode, inputData.Args)
	}

	InsectCrossInfo := entity.CrossChainInfo{
		Ic_request_id:  reqID,
		To_chainid:     chainId,
		Tx_status:      1,
		Source_service: 1,
		Tx_createtime:  time.Now(),
	}

	//todo 目标链信息 insert
	// request_id reqID
	// to_chainid  chainId
	// to_tx  res.TxId

	if err != nil {
		f.logger.Errorf("Fabric ChainId %s Chaincode %s %s has error %v", chainId, inputData.ChainCode, inputData.FunType, err)

		InsectCrossInfo.Tx_status = 2
		InsectCrossInfo.Error = err.Error()
		store.TargetChainInfo(&InsectCrossInfo)
		//如果处理失败如何返回信息
		return entity.GetErrOutPut(), types.NewResult(types.Status_Error, fmt.Sprintf("call chain %s has error : %s", inputData.GetChainId(), err.Error()))
	}

	InsectCrossInfo.To_tx = res.TxId
	store.TargetChainInfo(&InsectCrossInfo)

	//resBytes, _ := json.Marshal(res)
	//f.logger.Infof("call fabric ")
	//todo 如果不存在该chainId 信息，需要能返回不处理的信号，hub不用返回结果
	return entity.GetSuccessOutPut(*res), types.NewResult(types.Status_OK, "success")
}

func (f *FabricChainHandler) DeployIService(schemas string, pricing string) error {
	f.logger.Infof("starting to deploy %s service", f.serviceName)

	_, err := f.IServiceClient.ServiceClient.QueryServiceDefinition(f.serviceName)
	if err != nil {
		f.logger.Infof("defining service")

		err := f.IServiceClient.DefineService(f.serviceName, "", "", nil, schemas)
		if err != nil {
			return fmt.Errorf("failed to define service: %s", err.Error())
		}
	} else {
		f.logger.Infof("service definition already exists")
	}

	_, provider, err2 := f.IServiceClient.ServiceClient.Find(f.IServiceClient.KeyName, f.IServiceClient.Passphrase)
	if err2 != nil {
		return err2
	}

	_, err = f.IServiceClient.ServiceClient.QueryServiceBinding(f.serviceName, provider.String()) //
	if err != nil {
		f.logger.Infof("binding service")

		err := f.IServiceClient.BindService(f.serviceName, "100000point", pricing, "{}", 100)
		if err != nil {
			return fmt.Errorf("failed to bind service: %s", err.Error())
		}
	} else {
		f.logger.Infof("service binding already exists")
	}

	f.logger.Infof("%s service deployment completed", f.serviceName)

	return nil
}

func NewOutPut(res *entity.FabricRespone) string {

	output := &types.Output{
		Status: res.TxValidationCode == 0,
		TxHash: res.TxId,
	}

	resBytes, _ := json.Marshal(res)

	output.Result = string(resBytes)

	outPutBytes, _ := json.Marshal(output)

	return string(outPutBytes)
}
