package configbackend

type ChannelConfig struct {
	ChannelId string
	PeerName  string
}

////此方法用于设置config配置文件的channels部分的节点配置
func (o *ChannelConfig) SetChannelConfig(m *map[string]interface{}) {

	p := *m

	pm := make(map[string]interface{})

	pmp := make(map[string]interface{})
	cp := ChannelPeerConfig{PeerName: o.PeerName}
	cp.GetChannelPeerConfig(&pmp)

	pm["peers"] = pmp
	pm["policies"] = GetPoliciesConfig()
	p[o.ChannelId] = pm

}

//TODO：缺少注释
func GetPoliciesConfig() map[string]interface{} {
	m := make(map[string]interface{})

	mq := make(map[string]interface{})
	ro := make(map[string]interface{})

	mq["minResponses"] = 1
	mq["maxTargets"] = 1

	ro["attempts"] = 5
	ro["initialBackoff"] = "500ms"
	ro["maxBackoff"] = "5s"
	ro["backoffFactor"] = 2.0

	mq["retryOpts"] = ro

	m["queryChannelConfig"] = mq
	return m
}

type ChannelPeerConfig struct {
	PeerName string
}

//此方法用于设置config配置文件的channels-通道名-peers-节点部分的配置
func (o *ChannelPeerConfig) GetChannelPeerConfig(m *map[string]interface{}) map[string]interface{} {
	p := *m
	peerparameters := make(map[string]interface{})
	peerparameters["endorsingPeer"] = true
	peerparameters["chaincodeQuery"] = true
	peerparameters["ledgerQuery"] = true
	peerparameters["eventSource"] = true
	p[o.PeerName] = peerparameters
	return p
}
