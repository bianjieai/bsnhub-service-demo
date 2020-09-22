package ethereum

import (
	"github.com/bianjieai/bsnhub-service-demo/common"
	"github.com/spf13/viper"
)

const (
	Prefix = "eth"

	ChainID     = "chain_id"
	NodeRPCAddr = "node_rpc_addr"
	GasLimit    = "gas_limit"
	GasPrice    = "gas_price"
	Key         = "key"
	Passphrase  = "passphrase"

	NftCoreAddr = "nft_core_addr"
)

// Config represents the Ethereum chain config
type Config struct {
	ChainID     string `yaml:"chain_id"`
	NodeRPCAddr string `yaml:"node_rpc_addr"`
	GasLimit    uint64 `yaml:"gas_limit"`
	GasPrice    uint64 `yaml:"gas_price"`
	Key         string `yaml:"key"`
	Passphrase  string `yaml:"passphrase"`

	NftCoreAddr string `yaml:"nft_core_addr"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChainID:     v.GetString(common.GetConfigKey(Prefix, ChainID)),
		NodeRPCAddr: v.GetString(common.GetConfigKey(Prefix, NodeRPCAddr)),
		GasLimit:    v.GetUint64(common.GetConfigKey(Prefix, GasLimit)),
		GasPrice:    v.GetUint64(common.GetConfigKey(Prefix, GasPrice)),
		Key:         v.GetString(common.GetConfigKey(Prefix, Key)),
		Passphrase:  v.GetString(common.GetConfigKey(Prefix, Passphrase)),
		NftCoreAddr: v.GetString(common.GetConfigKey(Prefix, NftCoreAddr)),
	}
}
