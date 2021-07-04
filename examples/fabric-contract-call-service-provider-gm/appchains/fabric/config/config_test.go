package config

import (
	"fmt"
	"testing"
)

func TestNewSdkConfig(t *testing.T) {

	conf := Config{
		ChannelId:      "channel8899",
		PeerName:       "peer1",
		ConfigTmpPath:  "./conf/sdkconfig.yaml",
		PeerUrl:        "127.0.0.1",
		PeerTlsPath:    "string",
		OrgName:        "string",
		OrgMspId:       "string",
		OrgCryptoPath:  "string",
		UserName:       "string",
		UserSecret:     "string",
		IsDefUser:      true,
		CAAddress:      "string",
		CAEnrollId:     "string",
		CAEnrollSecret: "string",
	}

	configProvider := NewSdkConfig(&conf)

	ch, err := configProvider()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(ch))

	for i := range ch {
		m, ok := ch[i].Lookup("channels")
		fmt.Println(m, ok)
		k, ok := ch[i].Lookup("peers")
		fmt.Println(k, ok)
		p, ok := ch[i].Lookup("organizations")
		fmt.Println(p, ok)
		x, ok := ch[i].Lookup("orderers")
		fmt.Println(x, ok)
		g, ok := ch[i].Lookup("client")
		fmt.Println(g, ok)
	}

}
