package irishub

import (
	"github.com/bianjieai/bsnhub-service-demo/examples/record/types"
	sdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/modules/record"
	iritatypes "github.com/bianjieai/irita-sdk-go/types"
)

type IrisRecord struct {
	Client sdk.IRITAClient
	baseTx iritatypes.BaseTx
}

type Result struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (ir IrisRecord) Create(request types.Input) (recordId string, error string) {
	recordId, err := ir.Client.Record.CreateRecord(
		record.CreateRecordRequest{
			Contents: request.Contents,
		},
		ir.baseTx,
	)
	if err != nil {
		error = err.Error()
	}

	return
}
