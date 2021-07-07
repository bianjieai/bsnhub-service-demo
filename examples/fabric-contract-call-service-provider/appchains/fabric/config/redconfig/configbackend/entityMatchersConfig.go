/**
 * @Time : 2020-07-15 11:18
 * @Author : yz
 */

package configbackend

const (
	_EntityPeers    = "peer"
	_EntityOrderers = "orderer"
	_EntityTag      = "(\\w*)"
)

type EntityMatchersConfig struct {
	EntityPeers  []EntityPeer
	EntityOrders []EntityOrder
}

type EntityPeer struct {
	PeerName     string
	PeerUrl      string
	PeerEventUrl string
}

type EntityOrder struct {
	OrderName string
	OrderUrl  string
}

func (entity *EntityMatchersConfig) GetPeerEntityMatchersConfig() []interface{} {
	e := make([]interface{}, 0)

	for _, peer := range entity.EntityPeers {
		p := make(map[string]string)
		p["pattern"] = _EntityTag + peer.PeerName + _EntityTag
		p["urlSubstitutionExp"] = peer.PeerUrl
		p["eventUrlSubstitutionExp"] = peer.PeerEventUrl
		p["sslTargetOverrideUrlSubstitutionExp"] = peer.PeerName
		p["mappedHost"] = peer.PeerName
		e = append(e, p)
	}

	return e
}

func (entity *EntityMatchersConfig) GetOrderEntityMatchersConfig() []interface{} {
	e := make([]interface{}, 0)

	for _, order := range entity.EntityOrders {
		p := make(map[string]string)
		p["pattern"] = _EntityTag + order.OrderName + _EntityTag
		p["urlSubstitutionExp"] = order.OrderUrl
		p["sslTargetOverrideUrlSubstitutionExp"] = order.OrderName
		p["mappedHost"] = order.OrderName
		e = append(e, p)
	}

	return e
}

func (entity *EntityMatchersConfig) GetEntityMatchersConfig() map[string]interface{} {
	e := make(map[string]interface{})
	e[_EntityPeers] = entity.GetPeerEntityMatchersConfig()
	e[_EntityOrderers] = entity.GetOrderEntityMatchersConfig()
	return e
}

func (e *EntityMatchersConfig) SetEntityMatchersConfig(m *map[string]interface{}) {
	*m = e.GetEntityMatchersConfig()
}
