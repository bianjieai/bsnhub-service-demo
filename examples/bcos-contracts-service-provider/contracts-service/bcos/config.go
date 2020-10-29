package bcos

import (
	"github.com/spf13/viper"

	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/common"
)

const (
	Prefix = "bcos"

	ConfigFile = "config_file"
)

// Config represents the BCOS chain config
type Config struct {
	ConfigFile string `yaml:"config_file"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ConfigFile: v.GetString(common.GetConfigKey(Prefix, ConfigFile)),
	}
}
