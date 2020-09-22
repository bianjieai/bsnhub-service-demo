package service

import (
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft/nft/ethereum"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft/types"

	"github.com/bianjieai/irita-sdk-go/modules/service"
)

var serviceMap = make(map[string]service.RespondCallback)

const (
	NftServiceName = "nft_service"
)

func init() {
	serviceMap[NftServiceName] = nftService
}

func GetServiceCallBack(serviceName string) service.RespondCallback {
	return serviceMap[serviceName]
}

func nftService(reqCtxID, reqID, input string) (output string, result string) {
	var request types.Input
	res := types.Result{
		Code: 200,
	}
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())
	}

	// mint nft
	nftId, err := ethereum.EthChain{}.MintNft(request.To,
		request.AmountToMint, request.MetaId, request.SetPrice, request.IsForSale)

	if err != nil {
		res.Code = 500
		res.Message = err.Error()
	}

	if res.Code == 200 {
		outputBz, _ := json.Marshal(types.Output{NftId: nftId})
		output = string(outputBz)
	}

	resBz, _ := json.Marshal(res)
	result = string(resBz)
	return output, result
}
