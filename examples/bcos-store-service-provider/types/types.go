package types

const (
	ServiceName = "bcos-store"
)

type Input struct {
	Value string `json:"value"`
}

type Output struct {
	Key string `json:"key"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
