package contract_service

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/contract-service/opb"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/contract-service/opb/config"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/mysql"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/server"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/types"
	"github.com/bianjieai/irita-sdk-go/modules/wasm"
)

// ContractService defines the contract service
type ContractService struct {
	opbClient   *opb.OpbChain
	Logger      *log.Logger
}

// BuildContractService builds a ContractService instance from the given config
func BuildContractService(v *viper.Viper, chainManager *server.ChainManager) (ContractService, error) {
	baseConfig, err := config.NewBaseConfig(v)
	if err != nil {
		return ContractService{}, err
	}

	return ContractService{
		opbClient: opb.NewOpbChain(*baseConfig, chainManager),
	}, nil
}

// Callback implements the iservice.RespondCallback interface
func (cs ContractService) Callback(reqCtxID, reqID, input string) (output string, result string) {

	cs.Logger.Infof("service request received, request id: %s", reqID)
	res := &types.Result{
		Code: 200,
	}

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
	cs.Logger.Infof("Input is %s ", input)
	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		//参数不符合规则，直接不处理
		res.Code = 204
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
		return
	}

	chainParams, err := cs.opbClient.ChainManager.GetChainParams(request.Dest.ID)
	if err != nil {
		res.Code = 204
		res.Message = "chain params not exist"
		cs.Logger.Error("chain params not exist")
		return
	}

	mysql.OnServiceRequestReceived(reqID, request.Dest.ID)

	// instantiate the opb client with the specified group id and chain id
	err = cs.opbClient.InstantiateClient(chainParams)
	if err != nil {
		res.Code = 500
		res.Message = "failed to connect to the opb node"

		return
	}

	execAbi := wasm.NewContractABI().
		WithMethod("call_service").
		WithArgs("request_id", request.ReqSequence).
		WithArgs("endpoint_address", request.Dest.EndpointAddress).
		WithArgs("call_data", string(request.CallData))
	resultTx, err := cs.opbClient.OpbClient.WASM.Execute(chainParams.TargetCoreAddr, execAbi, nil, cs.opbClient.BuildBaseTx())
	if err != nil {
		mysql.TxErrCollection(reqID, err.Error())
		res.Code = 500
		res.Message = err.Error()
	}
	txHash = resultTx.Hash

	err = cs.opbClient.WaitForSuccess(resultTx.Hash, "callService")
	if err != nil {
		mysql.TxErrCollection(reqID, err.Error())
		res.Code = 500
		res.Message = err.Error()
	}
	callResult = resultTx.Data

	mysql.OnContractTxSend(reqID, txHash)

	return output, result
}
