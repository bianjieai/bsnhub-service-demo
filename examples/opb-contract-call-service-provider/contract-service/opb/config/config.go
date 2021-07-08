package config

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"

	"github.com/spf13/viper"

	"github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider/common"
)

const (
	Prefix = "opb"

	ChainId         = "chainId"
	KeyPath         = "key_path"
	KeyName         = "key_name"
	Passphrase      = "passphrase"
	RpcAddrsMap     = "rpc_addrs"
	GrpcAddrsMap    = "grpc_addrs"
	Timeout         = "timeout"
)

// BaseConfig defines the base config
type BaseConfig struct {
	KeyPath         string
	KeyName         string
	Passphrase      string
	RpcAddrsMap     map[string]string
	GrpcAddrsMap    map[string]string
	ChainId         string
	Timeout         uint
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURLs        []string `json:"nodes"`
	ChainID        int64    `json:"chainId"`
	TargetCoreAddr string   `json:"targetCoreAddr"`
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	config := new(BaseConfig)
	config.KeyPath = v.GetString(common.GetConfigKey(Prefix, KeyPath))
	config.ChainId = v.GetString(common.GetConfigKey(Prefix, ChainId))
	config.KeyName = v.GetString(common.GetConfigKey(Prefix, KeyName))
	config.Passphrase = v.GetString(common.GetConfigKey(Prefix, Passphrase))
	config.RpcAddrsMap = v.GetStringMapString(common.GetConfigKey(Prefix, RpcAddrsMap))
	config.GrpcAddrsMap = v.GetStringMapString(common.GetConfigKey(Prefix, GrpcAddrsMap))
	config.Timeout = v.GetUint(common.GetConfigKey(Prefix, Timeout))

	return config, nil
}


// ValidateBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}

func RandURL(m []string) string {
	if len(m) == 0 {
		return ""
	}
	for _, index := range rand.Perm(len(m)) {
		return m[index]
	}
	return ""
}

// GetChainID returns the unique opb chain id from the ChainID
func GetOpbChainID(ChainID string) int64 {
	chainInfos := strings.Split(ChainID, "-")
	opbChainID, _ := strconv.ParseInt(chainInfos[2], 10, 64)
	return opbChainID
}

// GetGroupID returns the unique opb group id from the ChainID
func GetOpbGroupID(ChainID string) int {
	chainInfos := strings.Split(ChainID, "-")
	opbGroupID, _ := strconv.Atoi(chainInfos[1])
	return opbGroupID
}
