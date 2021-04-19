package types

import (
	"fmt"
	"encoding/json"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service/fisco/config"
)

const (
	ServiceName = "fisco-contract-call"
)

type Input struct {
	OptType         string `json:"opt_type"`
	GroupID         int    `json:"group_id"`
	ChainID         int64  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	CallData        string `json:"call_data"`
	Height          uint64 `json:"height"`
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
func GetChainID(chainID int64, groupID int) string {
	return fmt.Sprintf("%s-%d-%d", "fisco", groupID, chainID)
}

// GetChainIDFromBytes returns the unique chain id from the given chain params bytes
func GetChainIDFromBytes(params []byte) (string, error) {
	var chainParams config.ChainParams
	err := json.Unmarshal(params, &chainParams)
	if err != nil {
		return "", err
	}

	return GetChainID(chainParams.ChainID, chainParams.GroupID), nil
}
