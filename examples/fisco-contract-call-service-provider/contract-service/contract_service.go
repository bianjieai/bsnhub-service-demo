package contract_service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	ethcmn "github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco/config"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/mysql"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/server"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/types"
)

// ContractService defines the contract service
type ContractService struct {
	FISCOClient *fisco.FISCOChain
	Logger      *log.Logger
}

// BuildContractService builds a ContractService instance from the given config
func BuildContractService(v *viper.Viper, chainManager *server.ChainManager) (ContractService, error) {
	baseConfig, err := config.NewBaseConfig(v)
	if err != nil {
		return ContractService{}, err
	}

	return ContractService{
		FISCOClient: fisco.NewFISCOChain(*baseConfig, chainManager),
	}, nil
}

// Callback implements the iservice.RespondCallback interface
func (cs ContractService) Callback(reqCtxID, reqID, input string) (output string, result string) {

	cs.Logger.Infof("service request received, request id: %s", reqID)
	res := &types.Result{
		Code: 200,
	}

	//var status bool
	var txHash string
	var callResult []byte

	defer func() {
		resBz, _ := json.Marshal(res)
		result = string(resBz)

		if res.Code == 200 {
			var outputBz []byte
			outputBz, _ = json.Marshal(types.Output{Result: hex.EncodeToString(callResult)})
			output = fmt.Sprintf(`{"header":{},"body":%s}`, outputBz)
		}

		cs.Logger.Infof("request processed, result: %s, output: %s", result, output)
	}()

	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		//参数不符合规则，直接不处理
		res.Code = 204
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
		return
	}
	contractAddress := ethcmn.HexToAddress(request.Dest.EndpointAddress)
	requestID, err := hex.DecodeString(request.Header.ReqSequence)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not decode requestID: %s", err.Error())

		return
	}
	var requestIDByte32 [32]byte
	copy(requestIDByte32[:], requestID)

	chainParams, err := cs.FISCOClient.ChainManager.GetChainParams(request.Dest.ID)
	if err != nil {
		res.Code = 204
		res.Message = "chain params not exist"
		cs.Logger.Error("chain params not exist")
		return
	}

	mysql.OnServiceRequestReceived(reqID, request.Dest.ID)

	// instantiate the fisco client with the specified group id and chain id
	err = cs.FISCOClient.InstantiateClient(chainParams)
	if err != nil {
		res.Code = 500
		res.Message = "failed to connect to the fisco node"

		return
	}

	tx, _, err := cs.FISCOClient.IServiceCoreSession.CallService(requestIDByte32, contractAddress, request.CallData)
	if err != nil {
		mysql.TxErrCollection(reqID, err.Error())
		res.Code = 500
		res.Message = err.Error()
	}

	receipt, err := cs.FISCOClient.WaitForReceipt(tx, "CallService")
	if err != nil {
		mysql.TxErrCollection(reqID, err.Error())
		res.Code = 500
		res.Message = err.Error()
	}
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address, cs.FISCOClient.BaseConfig.IServiceCoreAddr) {
			continue
		}

		data, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			cs.Logger.Errorf("failed to decode the log data: %s", err)
			continue
		}

		var event fisco.IServiceCoreExCrossChainResponseSent
		err = cs.FISCOClient.IServiceCoreABI.Unpack(&event, "CrossChainResponseSent", data)
		if err != nil {
			cs.Logger.Errorf("failed to unpack the log data: %s", err)
			continue
		}
		if event.EventName == "CrossChainRequestSent"{
			continue
		}

		callResult = event.Result
		break
	}



	mysql.OnContractTxSend(reqID, txHash)

	return output, result
}
