package bcos

import (
	"github.com/spf13/viper"

	"github.com/bianjieai/bsnhub-service-demo/common"
)

const (
	Prefix = "bcos"

	ConfigFile        = "config_file"
	StoreContractAddr = "store_contract_addr"
)

// Config represents the BCOS chain config
type Config struct {
	ConfigFile        string `yaml:"config_file"`
	StoreContractAddr string `yaml:"store_contract_addr"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ConfigFile:        v.GetString(common.GetConfigKey(Prefix, ConfigFile)),
		StoreContractAddr: v.GetString(common.GetConfigKey(Prefix, StoreContractAddr)),
	}
}
