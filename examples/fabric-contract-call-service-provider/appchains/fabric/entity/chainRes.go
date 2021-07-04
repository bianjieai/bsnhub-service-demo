package entity

import (
	"bsn-irita-fabric-provider/common"
	"encoding/hex"
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func NewFabricRes(fabRes channel.Response) *FabricRespone {
	res := &FabricRespone{
		TxValidationCode: int32(fabRes.TxValidationCode),
		TxId:             string(fabRes.TransactionID),
		Payload:          "0x" + hex.EncodeToString(fabRes.Payload),
		ChaincodeStatus:  fabRes.ChaincodeStatus,
	}

	if res.TxValidationCode != 0 {
		common.Logger.Errorf("call fabric res  TxValidationCode is %d", res.TxValidationCode)
	}
	return res
}

type FabricRespone struct {
	TxValidationCode int32  `json:"txValidationCode"`
	ChaincodeStatus  int32  `json:"chaincodeStatus"`
	TxId             string `json:"txId"`
	Payload          string `json:"payload"`
}

type OutPut struct {
	Header interface{}   `json:"header"`
	Body   FabricRespone `json:"body"`
}

func GetErrOutPut() string {

	outPut := &OutPut{
		Header: struct{}{},
		Body:   FabricRespone{},
	}

	jsonBytes, _ := json.Marshal(outPut)

	return string(jsonBytes)
}

func GetSuccessOutPut(res FabricRespone) string {
	outPut := &OutPut{
		Header: struct{}{},
		Body:   res,
	}
	jsonBytes, _ := json.Marshal(outPut)

	return string(jsonBytes)
}
