package ethereum

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ParseABI parses the given JSON-formatted abi
func ParseABI(abiJSON string) (abi.ABI, error) {
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return abi.ABI{}, err
	}

	return parsedABI, nil
}
