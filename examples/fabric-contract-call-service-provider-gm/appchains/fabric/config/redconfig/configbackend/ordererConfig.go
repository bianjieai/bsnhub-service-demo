package configbackend

type OrdererConfig struct {
	OrdererName    string
	OrdererUrl     string
	TlsCACertsPath string
}

func (o *OrdererConfig) GetOrdererConfig() map[string]interface{} {

	p := make(map[string]interface{})

	p["url"] = o.OrdererUrl
	p["grpcOptions"] = GetDefGrpcOptions(o.OrdererName)

	SetTlsCaCertPath(&p, o.TlsCACertsPath)

	return p
}

func (o *OrdererConfig) SetOrderer(m *map[string]interface{}) {
	mi := *m
	mi[o.OrdererName] = o.GetOrdererConfig()
}
