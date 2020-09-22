package types

import "github.com/bianjieai/irita-sdk-go/modules/record"

type Input struct {
	Contents []record.Content `json:"contents"`
}

type Output struct {
	RecordId string `json:"record_id"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
