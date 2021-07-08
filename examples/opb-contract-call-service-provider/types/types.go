package types

import (
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/contract-service/opb/config"
	"strconv"
)

type Header struct {
	ReqSequence string `json:"req_sequence"`
	ChainID     string `json:"id"`
}

type Body struct {
	Source `json:"source"`
	Dest   `json:"dest"`
	Method string `json:"method"`
	CallData   []byte `json:"args"`
}

type Source struct {
	ID              string `json:"id"`
	ChainID         string `json:"chain_id"`
	SubChainID      string `json:"sub_chain_id"`
	EndpointType    string `json:"endpoint_type"`
	EndpointAddress string `json:"endpoint_address"`
	Sender          string `json:"sender"`
	TxHash          string `json:"tx_hash"`
}

type Dest struct {
	ID              string `json:"id"`
	ChainID         string `json:"chain_id"`
	SubChainID      string `json:"sub_chain_id"`
	EndpointType    string `json:"endpoint_type"`
	EndpointAddress string `json:"endpoint_address"`
}

// ServiceInput defines the service input
type Input struct {
	Header `json:"header"`
	Body   `json:"body"`
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


func GetChainIDString(chainID int64) string {
	return strconv.FormatInt(chainID, 10)
}

func GetChainID(chainID int64) string {
	//return strconv.FormatInt(chainID,10)
	return fmt.Sprintf("%s-%d", "opb", chainID)
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