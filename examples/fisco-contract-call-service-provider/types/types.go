package types

import (
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco/config"
	"strconv"
)

const (
//ServiceName = "cross_service_dev"
)

type Input struct {
	OptType string `json:"opt_type"`
	//GroupID         int    `json:"group_id"`
	ChainID         int64  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	CallData        string `json:"call_data"`
	//Height          uint64 `json:"height"`
}

type Output struct {
	Result string `json:"result,omitempty"`
	Status bool   `json:"status,omitempty"`
	TxHash string `json:"tx_hash,omitempty"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetChainID returns the unique chain id from the specified chain params
func GetChainIDString(chainID int64) string {
	return strconv.FormatInt(chainID, 10)
}

func GetChainID(chainID int64) string {
	//return strconv.FormatInt(chainID,10)
	return fmt.Sprintf("%s-%d-%d", "fisco", 11, chainID)
}

// GetChainIDFromBytes returns the unique chain id from the given chain params bytes
func GetChainIDFromBytes(params []byte) (string, error) {
	var chainParams config.ChainParams
	err := json.Unmarshal(params, &chainParams)
	if err != nil {
		return "", err
	}

	return GetChainID(chainParams.ChainID), nil
}
