package types

import "encoding/json"

const (

	//ok
	Status_OK = 200

	//目标链信息不存在
	Status_Chain_NotExist = 204

	//参数不正确
	Status_Params_Error = 400

	//访问异常
	Status_Error = 500
)

func NewResult(code int, msg string) string {

	result := &Result{
		Code:    code,
		Message: msg,
	}
	jsonBytes, _ := json.Marshal(result)
	return string(jsonBytes)
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Output struct {
	Result string `json:"result,omitempty"`
	Status bool   `json:"status,omitempty"`
	TxHash string `json:"tx_hash,omitempty"`
}
