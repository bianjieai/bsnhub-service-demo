package redconfig

import (
	config "bsn-irita-fabric-provider/appchains/fabric/config/redconfig/configbackend"
)

const (
	peerNodeName          = "peers"
	channelsNodeName      = "channels"
	organizationsNodeName = "organizations"
	caNodeName            = "certificateAuthorities"
	ordererNodeName       = "orderers"
	entityMatchersName    = "entityMatchers"
)

//todo 在这里修改参数配置，单个修改为多个 orderer 已修改
func SetPeer(peers *[]config.PeerConfig) SetOption {
	return func(def *defConfigBackend) error {
		m := make(map[string]interface{})
		for _, item := range *peers {
			item.SetPeerConfig(&m)
		}
		def.Set(peerNodeName, m)
		return nil
	}
}

func SetChannel(ch *config.ChannelConfig) SetOption {
	return func(def *defConfigBackend) error {

		m := make(map[string]interface{})
		ch.SetChannelConfig(&m)
		def.Set(channelsNodeName, m)

		return nil
	}
}

func SetOrg(org *config.OrganizationConfig) SetOption {
	return func(def *defConfigBackend) error {

		m := make(map[string]interface{})
		org.SetOrganizationConfig(&m)
		def.Set(organizationsNodeName, m)

		return nil
	}
}

func SetCa(ca *config.CertificateAuthoritiesConfig) SetOption {
	return func(def *defConfigBackend) error {

		m := make(map[string]interface{})
		ca.SetCertificateAuthoritiesConfig(&m)
		def.Set(caNodeName, m)

		return nil
	}
}

func SetOrderer(orderers *[]config.OrdererConfig) SetOption {
	return func(def *defConfigBackend) error {

		m := make(map[string]interface{})
		for _, o := range *orderers {
			o.SetOrderer(&m)
		}

		def.Set(ordererNodeName, m)

		return nil
	}
}

func SetEntityMatchers(entitys *config.EntityMatchersConfig) SetOption {
	return func(def *defConfigBackend) error {

		m := make(map[string]interface{})
		entitys.SetEntityMatchersConfig(&m)

		def.Set(entityMatchersName, m)

		return nil
	}
}
