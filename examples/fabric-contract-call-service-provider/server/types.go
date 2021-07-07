package server

const (
	CODE_SUCCESS = 1
	CODE_ERROR   = 0
)

// SuccessResponse defines the response on success
type SuccessResponse struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg,omitempty"`
	Result interface{} `json:"data,omitempty"`
}

// ErrorResponse defines the response on error
type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"msg"`
}
