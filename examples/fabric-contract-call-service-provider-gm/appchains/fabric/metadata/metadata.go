package metadata

import "strconv"

type CrossData struct {
	Header interface{}   `json:"header,omitempty"`
	Body   *FabricIutput `json:"body,omitempty"`
}

type FabricIutput struct {
	ChainId   uint64   `json:"chainId"`
	ChainCode string   `json:"chainCode"`
	FunType   string   `json:"funType"`
	Args      []string `json:"args"`
}

func (r *FabricIutput) GetChainId() string {
	return strconv.FormatUint(r.ChainId, 10)
}
