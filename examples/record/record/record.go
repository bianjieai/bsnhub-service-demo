package record

import "github.com/bianjieai/bsnhub-service-demo/examples/record/types"

type Record interface {
	Create(request types.Input) (recordId string, error string)
}
