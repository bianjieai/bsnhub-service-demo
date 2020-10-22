package store

import (
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"

	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/store/bcos"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/types"
)

// StoreService defines the store service
type StoreService struct {
	BCOSClient bcos.BCOSChain
	Logger     *log.Logger
}

// NewStoreService constructs a new StoreService instance
func NewStoreService(bcosClient bcos.BCOSChain, logger *log.Logger) StoreService {
	return StoreService{
		BCOSClient: bcosClient,
		Logger:     logger,
	}
}

// MakeStoreService builds a StoreService instance from the given config
func MakeStoreService(config *viper.Viper) StoreService {
	return StoreService{
		BCOSClient: bcos.MakeBCOSChain(bcos.NewConfig(config)),
	}
}

// Store stores the given data to BCOS store contract
func (s StoreService) Store(
	value string,
) (string, error) {
	tx, _, err := s.BCOSClient.StoreSession.Set(value)
	if err != nil {
		return "", fmt.Errorf("failed to send Store transaction: %s", err)
	}

	receipt, err := s.BCOSClient.WaitForReceipt(tx, "Store")
	if err != nil {
		return "", err
	}

	return receipt.BlockHash, nil
}

// Callback implements the iservice.RespondCallback interface
func (s StoreService) Callback(reqCtxID, reqID, input string) (output string, result string) {
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
	key, err = s.Store(request.Value)
	if err != nil {
		res.Code = 500
		res.Message = err.Error()
	}

	return output, result
}
