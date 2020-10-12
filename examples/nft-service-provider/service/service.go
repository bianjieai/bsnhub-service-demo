package service

import (
	"encoding/json"
	"fmt"
	"math/big"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/service/ethereum"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/types"
)

// Service defines the service provided by the owner
type Service struct {
	EthClient ethereum.EthChain
	Logger    *log.Logger
}

// NewService constructs a new Service instance
func NewService(ethClient ethereum.EthChain, logger *log.Logger) Service {
	return Service{
		EthClient: ethClient,
		Logger:    logger,
	}
}

// MakeService builds a Service instance from the given config
func MakeService(config *viper.Viper) Service {
	return Service{
		EthClient: ethereum.MakeEthChain(ethereum.NewConfig(config)),
	}
}

// MintNft mints an NFT on ethereum
func (s Service) MintNft(
	to ethcmn.Address,
	amountToMint *big.Int,
	metaID string,
	setPrice *big.Int,
	isForSale bool,
) (string, error) {
	auth, err := s.EthClient.BuildAuthTransactor()
	if err != nil {
		return "", err
	}

	tx, err := s.EthClient.NftContract.BatchMint(auth, to, amountToMint, metaID, setPrice, isForSale)
	if err != nil {
		return "", fmt.Errorf("failed to send BatchMint transaction: %s", err)
	}

	nftID, err := s.EthClient.WaitForReceipt(tx, "BatchMint")
	if err != nil {
		return "", err
	}

	return nftID, nil
}

// Callback implements the iservice.RespondCallback interface
func (s Service) Callback(reqCtxID, reqID, input string) (result string, output string) {
	s.Logger.Infof("service request received, request context id: %s, request id: %s", reqCtxID, reqID)

	res := &types.Result{
		Code: 200,
	}

	var nftID string

	defer func() {
		resBz, _ := json.Marshal(res)
		result = string(resBz)

		if res.Code == 200 {
			outputBz, _ := json.Marshal(types.Output{NftID: nftID})
			output = string(outputBz)
		}

		s.Logger.Infof("request processed, result: %s, output: %s", result, output)
	}()

	var request types.Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		res.Code = 400
		res.Message = fmt.Sprintf("can not parse request [%s] input json string : %s", reqID, err.Error())

		return
	}

	// decode abi-encoded args
	if len(request.ABIEncoded) > 0 {
		var rawRequest types.RawInput
		err = s.EthClient.NftContractABI.Unpack(&rawRequest, "BatchMint", []byte(request.ABIEncoded))
		if err != nil {
			res.Code = 400
			res.Message = err.Error()

			return
		}

		request.To = rawRequest.To
		request.AmountToMint = rawRequest.AmountToMint
		request.MetaID = rawRequest.MetaID
		request.SetPrice = rawRequest.SetPrice
		request.IsForSale = rawRequest.IsForSale
	}

	// mint nft
	nftID, err = s.MintNft(
		request.To,
		request.AmountToMint,
		request.MetaID,
		request.SetPrice,
		request.IsForSale,
	)
	if err != nil {
		res.Code = 500
		res.Message = err.Error()
	}

	return result, output
}
