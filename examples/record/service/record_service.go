package service

import (
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/record/record/irishub"
	"github.com/bianjieai/bsnhub-service-demo/examples/record/types"

	"github.com/bianjieai/irita-sdk-go/modules/service"
)

var serviceMap = make(map[string]service.RespondCallback)

const (
	RecordServiceName = "record_service"
)

func init() {
	serviceMap[RecordServiceName] = recordService
}

func GetServiceCallBack(serviceName string) service.RespondCallback {
	return serviceMap[serviceName]
}

func recordService(reqCtxID, reqID, input string) (output string, result string) {
	var request types.Input
	res := types.Result{
		Code: 200,
	}
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
	}

	// save record
	recordId, errMsg := irishub.IrisRecord{}.SaveRecord(request)

	if len(errMsg) > 0 {
		res.Code = 500
		res.Message = errMsg
	}

	if res.Code == 200 {
		outputBz, _ := json.Marshal(types.Output{RecordId: recordId})
		output = string(outputBz)
	}

	resBz, _ := json.Marshal(res)
	result = string(resBz)
	return output, result
}
