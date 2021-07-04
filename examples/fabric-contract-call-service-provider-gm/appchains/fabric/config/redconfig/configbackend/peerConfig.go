package configbackend

//peers:

type PeerConfig struct {
	PeerName string

	PeerUrl        string
	PeerEventUrl   string
	TlsCACertsPath string
}

func (c *PeerConfig) GetPeerConfig() map[string]interface{} {

	p := make(map[string]interface{})
	p["url"] = c.PeerUrl
	p["eventUrl"] = c.PeerEventUrl
	p["grpcOptions"] = GetDefGrpcOptions(c.PeerName)

	SetTlsCaCertPath(&p, c.TlsCACertsPath)

	return p
}

func (o *PeerConfig) SetPeerConfig(m *map[string]interface{}) {
	mi := *m
	mi[o.PeerName] = o.GetPeerConfig()
}
