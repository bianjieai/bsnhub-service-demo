package config

import (
	"encoding/json"
	"fmt"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/types"
	"math/rand"
	"strings"

	"github.com/spf13/viper"

	"github.com/FISCO-BCOS/go-sdk/conf"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
)

const (
	Prefix = "fisco"

	ChainId        = "chainId"
	ConnectionType = "connection_type"
	CAFile         = "ca_file"
	CertFile       = "cert_file"
	KeyFile        = "key_file"
	SMCrypto       = "sm_crypto"
	PrivateKeyFile = "priv_key_file"
	IServiceCoreAddr   = "iservice_core_addr"
)

// BaseConfig defines the base config
type BaseConfig struct {
	IsHTTP     bool
	ChainId    string
	CAFile     string
	KeyFile    string
	CertFile   string
	PrivateKey []byte
	IsSMCrypto bool
	NodesMap   map[string]string
	IServiceCoreAddr string
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURL []string `json:"nodes"`
	ChainID string    `json:"chainId"`
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	connType := v.GetString(common.GetConfigKey(Prefix, ConnectionType))
	caFile := v.GetString(common.GetConfigKey(Prefix, CAFile))
	certFile := v.GetString(common.GetConfigKey(Prefix, CertFile))
	keyFile := v.GetString(common.GetConfigKey(Prefix, KeyFile))
	smCrypto := v.GetBool(common.GetConfigKey(Prefix, SMCrypto))
	privKeyFile := v.GetString(common.GetConfigKey(Prefix, PrivateKeyFile))
	chainId := v.GetString(common.GetConfigKey(Prefix, ChainId))
	iServiceCoreAddr := v.GetString(common.GetConfigKey(Prefix, IServiceCoreAddr))
	config := new(BaseConfig)

	if strings.EqualFold(connType, "rpc") {
		config.IsHTTP = true
	} else if strings.EqualFold(connType, "channel") {
		config.IsHTTP = false
	} else {
		return nil, fmt.Errorf("connection type %s is not supported", connType)
	}

	config.IsSMCrypto = smCrypto

	keyBytes, curve, err := conf.LoadECPrivateKeyFromPEM(privKeyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key, err: %v", err)
	}

	if config.IsSMCrypto && curve != "sm2p256v1" {
		return nil, fmt.Errorf("smcrypto must use sm2p256v1 private key, but found %s", curve)
	}
	if !config.IsSMCrypto && curve != "secp256k1" {
		return nil, fmt.Errorf("must use secp256k1 private key, but found %s", curve)
	}

	config.ChainId = chainId
	config.PrivateKey = keyBytes
	config.CAFile = caFile
	config.CertFile = certFile
	config.KeyFile = keyFile
	config.NodesMap = v.GetStringMapString(common.ConfigKeyNodes)
	config.IServiceCoreAddr = iServiceCoreAddr
	common.Logger.Infof("config fisco nods : %v", config.NodesMap)

	return config, nil
}

// BuildClientConfig builds the FISCO client config from the given Config
func BuildClientConfig(config Config) *conf.Config {
	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := randURL(config.NodeURL)
	//获取配置的nodeURL
	nodeUrl, ok := config.NodesMap[nodeName]
	if ok {
		nodeName = nodeUrl
	}

	return &conf.Config{
		IsHTTP:     config.IsHTTP,
		CAFile:     config.CAFile,
		Key:        config.KeyFile,
		Cert:       config.CertFile,
		PrivateKey: config.PrivateKey,
		IsSMCrypto: config.IsSMCrypto,
		GroupID:    types.GetFiscoGroupID(config.ChainID),
		ChainID:    types.GetFiscoChainID(config.ChainID),
		NodeURL:    nodeName,
	}
}

// ValidateBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}

func randURL(m []string) string {
	if len(m) == 0 {
		return ""
	}
	for _, index := range rand.Perm(len(m)) {
		return m[index]
	}
	return ""
}
