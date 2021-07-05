package metadata

import "strconv"

type Header struct {
	ReqSequence string `json:"req_sequence"`
	ChainID     string `json:"id"`
}

type Body struct {
	*Source `json:"source"`
	*Dest   `json:"dest"`
	Method string `json:"method"`
	CallData   []byte `json:"args"`
}

type Source struct {
	ID              string `json:"id"`
	ChainID         string `json:"chain_id"`
	SubChainID      string `json:"sub_chain_id"`
	EndpointType    string `json:"endpoint_type"`
	EndpointAddress string `json:"endpoint_address"`
	Sender          string `json:"sender"`
	TxHash          string `json:"tx_hash"`
}

type Dest struct {
	ID              string `json:"id"`
	ChainID         string `json:"chain_id"`
	SubChainID      string `json:"sub_chain_id"`
	EndpointType    string `json:"endpoint_type"`
	EndpointAddress string `json:"endpoint_address"`
}

// ServiceInput defines the service input
type CrossData struct {
	*Header `json:"header"`
	*Body   `json:"body"`
}

type FabricIutput struct {
	ChainId   uint64   `json:"chainId"`
	ChainCode string   `jso n:"chainCode"`
	FunType   string   `json:"funType"`
	Args      []string `json:"args"`
}

func (r *FabricIutput) GetChainId() string {
	return strconv.FormatUint(r.ChainId, 10)
}
