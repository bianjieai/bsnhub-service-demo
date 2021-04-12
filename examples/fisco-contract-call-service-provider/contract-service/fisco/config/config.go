package config

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"github.com/FISCO-BCOS/go-sdk/conf"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
)

const (
	Prefix = "fisco"

	ConnectionType = "connection_type"
	CAFile         = "ca_file"
	CertFile       = "cert_file"
	KeyFile        = "key_file"
	SMCrypto       = "sm_crypto"
	PrivateKeyFile = "priv_key_file"
)

// BaseConfig defines the base config
type BaseConfig struct {
	IsHTTP     bool
	CAFile     string
	KeyFile    string
	CertFile   string
	PrivateKey []byte
	IsSMCrypto bool
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURL string `json:"node_url"`
	GroupID int   `json:"group_id"`
	ChainID int64 `json:"chain_id"`
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

	config.PrivateKey = keyBytes
	config.CAFile = caFile
	config.CertFile = certFile
	config.KeyFile = keyFile

	return config, nil
}

// NewConfig constructs a new Config instance
func NewConfig(baseConfig BaseConfig, chainParams ChainParams) *Config {
	return &Config{
		BaseConfig:  baseConfig,
		ChainParams: chainParams,
	}
}

// BuildClientConfig builds the FISCO client config from the given Config
func BuildClientConfig(config Config) *conf.Config {

	return &conf.Config{
		IsHTTP:     config.IsHTTP,
		CAFile:     config.CAFile,
		Key:        config.KeyFile,
		Cert:       config.CertFile,
		PrivateKey: config.PrivateKey,
		IsSMCrypto: config.IsSMCrypto,
		GroupID:    config.GroupID,
		ChainID:    config.ChainID,
		NodeURL:    config.NodeURL,
	}
}

// ValidateBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}
