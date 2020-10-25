package service

import (
	"encoding/json"
	"fmt"

	"github.com/bianjieai/bsnhub-service-demo/examples/market/market"
	"github.com/bianjieai/irita-sdk-go/modules/service"
)

var serviceMap = make(map[string]service.RespondCallback)

const (
	PriceServiceName = "price_service"
)

func init() {
	serviceMap[PriceServiceName] = priceService
}

func GetServiceCallBack(serviceName string) service.RespondCallback {
	return serviceMap[serviceName]
}

func priceService(reqCtxID, reqID, input string) (output string, result string) {
	var request Input
	res := Result{
		Code: 200,
	}
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
	}

	// get price from public market
	mk := market.GetMarket()
	price, errMsg := mk.GetPrice(request.Base, request.Quote)

	if len(errMsg) > 0 {
		res.Code = 500
		res.Message = errMsg
	}

	if res.Code == 200 {
		outputBz, _ := json.Marshal(Output{Price: price})
		output = string(outputBz)
	}

	resBz, _ := json.Marshal(res)
	result = string(resBz)
	return output, result
}

type Input struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type Output struct {
	Price float64 `json:"price"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
