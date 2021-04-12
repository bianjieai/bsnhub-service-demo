package server

import "fmt"

const (
	CODE_SUCCESS = 0
	CODE_ERROR   = 1
)

// AddChainRequest defines the request to add an app chain
type AddChainRequest struct {
	ChainParams string `json:"chain_params"`
}

// AddChainResult defines the result for adding an app chain
type AddChainResult struct {
	ChainID string `json:"chain_id"`
}

// SuccessResponse defines the response on success
type SuccessResponse struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result,omitempty"`
}

// ErrorResponse defines the response on error
type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// ValidateChainID validates the given chain ID
func ValidateChainID(chainID string) error {
	if len(chainID) == 0 {
		return fmt.Errorf("chain ID can not be empty")
	}

	return nil
}
