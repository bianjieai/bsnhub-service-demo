package iservice

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/irisnet/service-sdk-go/types"

	"bsn-irita-fabric-provider/common"
)

// default config variables
var (
	defaultChainID       = "irita-hub"
	defaultNodeRPCAddr   = "http://127.0.0.1:26657"
	defaultNodeGRPCAddr  = "127.0.0.1:9090"
	defaultKeyPath       = os.ExpandEnv(filepath.Join("$HOME", ".iritacli"))
	defaultGas           = uint64(200000)
	defaultFee           = "4point"
	defaultBroadcastMode = types.Commit
	defaultKeyAlgorithm  = "sm2"
)

const (
	Prefix       = "iservice"
	ChainID      = "chain_id"
	NodeRPCAddr  = "node_rpc_addr"
	NodeGRPCAddr = "node_grpc_addr"
	KeyPath      = "key_path"
	KeyName      = "key_name"
	Passphrase   = "passphrase"
)

// Config is a config struct for iservice
type Config struct {
	ChainID      string `yaml:"chain_id"`
	NodeRPCAddr  string `yaml:"node_rpc_addr"`
	NodeGRPCAddr string `yaml:"node_grpc_addr"`
	KeyPath      string `yaml:"key_path"`
	KeyName      string `yaml:"key_name"`
	Passphrase   string `yaml:"passphrase"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ChainID:      v.GetString(common.GetConfigKey(Prefix, ChainID)),
		NodeRPCAddr:  v.GetString(common.GetConfigKey(Prefix, NodeRPCAddr)),
		NodeGRPCAddr: v.GetString(common.GetConfigKey(Prefix, NodeGRPCAddr)),
		KeyPath:      v.GetString(common.GetConfigKey(Prefix, KeyPath)),
		KeyName:      v.GetString(common.GetConfigKey(Prefix, KeyName)),
		Passphrase:   v.GetString(common.GetConfigKey(Prefix, Passphrase)),
	}
}
