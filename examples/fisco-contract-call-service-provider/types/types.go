package types

const (
	ServiceName = "fisco-contract-call"
)

type Input struct {
	GroupID         int    `json:"group_id"`
	ChainID         int64  `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
	CallData        string `json:"call_data"`
	Height          uint64 `json:"height"`
}

type Output struct {
	Result string `json:"result"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
