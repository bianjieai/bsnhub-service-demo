package contracts_service

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"

	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/contracts-service/bcos"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/types"
)

// ContractsService defines the contracts service
type ContractsService struct {
	BCOSClient bcos.BCOSChain
	Logger     *log.Logger
}

// NewContractsService constructs a new ContractsService instance
func NewContractsService(bcosClient bcos.BCOSChain, logger *log.Logger) ContractsService {
	return ContractsService{
		BCOSClient: bcosClient,
		Logger:     logger,
	}
}

// MakeContractsService builds a ContractsService instance from the given config
func MakeContractsService(config *viper.Viper) ContractsService {
	return ContractsService{
		BCOSClient: bcos.MakeBCOSChain(bcos.NewConfig(config)),
	}
}

// Invoke initiates the function invocation with the given args
func (s ContractsService) Invoke(
	args string,
) (string, error) {
	// TODO
	return "", nil
}

// Callback implements the iservice.RespondCallback interface
func (s ContractsService) Callback(reqCtxID, reqID, input string) (output string, result string) {
	s.Logger.Infof("service request received, request id: %s", reqID)

	res := &types.Result{
		Code: 200,
	}

	var key string

	defer func() {
		resBz, _ := json.Marshal(res)
		result = string(resBz)

		if res.Code == 200 {
			outputBz, _ := json.Marshal(types.Output{Key: key})
			output = fmt.Sprintf(`{"header":{},"body":%s}`, outputBz)
		}

		s.Logger.Infof("request processed, result: %s, output: %s", result, output)
	}()

	input = gjson.Get(input, "body").String()

	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())

		return
	}

	// store
	key, err = s.Invoke(request.Value)
	if err != nil {
		res.Code = 500
		res.Message = err.Error()
	}

	return output, result
}
