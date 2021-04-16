package contract_service

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/types"
	config2 "github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco/config"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/server"
)

// ContractService defines the contract service
type ContractService struct {
	FISCOClient *fisco.FISCOChain
	Logger      *log.Logger
}

// NewContractService constructs a new ContractService instance
func NewContractService(fiscoChain *fisco.FISCOChain, logger *log.Logger) ContractService {
	return ContractService{
		FISCOClient: fiscoChain,
		Logger:      logger,
	}
}

// BuildContractService builds a ContractService instance from the given config
func BuildContractService(config *viper.Viper, chainManager *server.ChainManager) (ContractService, error) {
	baseConfig, err := config2.NewBaseConfig(config)
	if err != nil {
		return ContractService{}, err
	}

	return ContractService{
		FISCOClient: fisco.NewFISCOChain(*baseConfig, chainManager),
	}, nil
}

// CallContract initiates a contract call with the given contract address and data
func (cs ContractService) CallContract(
	contractAddress ethcmn.Address,
	data []byte,
) (result []byte, err error) {
	return cs.FISCOClient.CallContract(contractAddress, data)
}

// SendContractTx initiates a contract tx with the given contract address and data
func (cs ContractService) SendContractTx(
	contractAddress ethcmn.Address,
	data []byte,
) (status bool, txHash string, err error) {
	return cs.FISCOClient.SendContractTx(contractAddress, data)
}

// Callback implements the iservice.RespondCallback interface
func (cs ContractService) Callback(reqCtxID, reqID, input string) (output string, result string) {
	cs.Logger.Infof("service request received, request id: %s", reqID)

	res := &types.Result{
		Code: 200,
	}

	var status bool
	var txHash string
	var optType string
	var callResult []byte

	defer func() {
		resBz, _ := json.Marshal(res)
		result = string(resBz)

		if res.Code == 200 {
			var outputBz []byte
			if optType == "call"{
				outputBz, _ = json.Marshal(types.Output{Result: hex.EncodeToString(callResult)})
			}else if optType == "tx"{
				outputBz, _ = json.Marshal(types.Output{Status: status, TxHash: txHash})
			}
			output = fmt.Sprintf(`{"header":{},"body":%s}`, outputBz)
		}

		cs.Logger.Infof("request processed, result: %s, output: %s", result, output)
	}()

	input = gjson.Get(input, "body").String()

	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())

		return
	}
	optType = request.OptType

	contractAddress := ethcmn.HexToAddress(request.ContractAddress)

	callData, err := hex.DecodeString(request.CallData)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not decode call data: %s", err.Error())

		return
	}

	// instantiate the fisco client with the specified group id and chain id
	err = cs.FISCOClient.InstantiateClient(request.GroupID, request.ChainID)
	if err != nil {
		res.Code = 500
		res.Message = "failed to connect to the fisco node"

		return
	}
	if optType == "call"{
		callResult, err = cs.CallContract(contractAddress, callData)
	}else if optType == "tx"{
		// send contract tx
		status, txHash, err = cs.SendContractTx(contractAddress, callData)
	}else{
		err = errors.New("wrong operation type")
	}
	if err != nil {
		res.Code = 500
		res.Message = err.Error()
	}

	return output, result
}
