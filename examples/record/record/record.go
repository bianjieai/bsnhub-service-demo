package record

import "github.com/bianjieai/bsnhub-service-demo/examples/record/types"

type Record interface {
	SaveRecord(request types.Input) (recordId string, error string)
}
