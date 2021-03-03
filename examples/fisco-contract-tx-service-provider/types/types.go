package types

const (
	ServiceName = "fisco-contract-tx"
)

type Input struct {
	GroupID         int    `json:"group_id"`
	ChainID         int64  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	CallData        string `json:"call_data"`
}

type Output struct {
	Status bool   `json:"status"`
	TxHash string `json:"tx_hash"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
