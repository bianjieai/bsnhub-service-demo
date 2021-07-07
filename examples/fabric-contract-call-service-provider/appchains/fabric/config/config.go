package config

import (
	"bsn-irita-fabric-provider/appchains/fabric/config/redconfig"
	"bsn-irita-fabric-provider/appchains/fabric/config/redconfig/configbackend"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
)

type Config struct {
	ConfigTmpPath string

	ChannelId   string
	PeerName    string
	PeerUrl     string
	PeerTlsPath string

	OrgName       string
	OrgMspId      string
	OrgCryptoPath string

	UserName       string
	UserSecret     string
	IsDefUser      bool
	CAAddress      string
	CAEnrollId     string
	CAEnrollSecret string
}

func NewSdkConfig(conf *Config) core.ConfigProvider {

	ch := configbackend.ChannelConfig{ChannelId: conf.ChannelId, PeerName: conf.PeerName}

	var peers []configbackend.PeerConfig
	peer := configbackend.PeerConfig{PeerName: conf.PeerName, PeerUrl: conf.PeerUrl, PeerEventUrl: conf.PeerUrl, TlsCACertsPath: conf.PeerTlsPath}
	peers = append(peers, peer)

	org := configbackend.OrganizationConfig{OrgName: conf.OrgName, MspId: conf.OrgMspId, CryptoPath: conf.OrgCryptoPath, Peers: []string{conf.PeerName}}

	var s []redconfig.SetOption
	s = append(s, redconfig.SetPeer(&peers))
	s = append(s, redconfig.SetChannel(&ch))
	s = append(s, redconfig.SetOrg(&org))

	configProvider := redconfig.FromFile(conf.ConfigTmpPath, s)
	return configProvider

}
