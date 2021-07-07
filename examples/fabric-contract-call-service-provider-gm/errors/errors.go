package errors

import "fmt"

func New(f string, a ...interface{}) error {

	return &RelayerError{
		code: "9999",
		msg:  fmt.Sprintf(f, a...),
	}

}

func NewErrCode(code string) error {
	return &RelayerError{
		code: code,
	}
}

type RelayerError struct {
	code string
	msg  string
}

func (r *RelayerError) Error() string {
	return r.msg
}

func (r *RelayerError) Code() string {
	return r.code
}

type ChanError struct {
	HasError bool
	Err      error
}

func NewChanError(err error) *ChanError {

	return &ChanError{
		HasError: true,
		Err:      err,
	}
}

func NewChanSuccess() *ChanError {

	return &ChanError{
		HasError: false,
		Err:      nil,
	}
}
