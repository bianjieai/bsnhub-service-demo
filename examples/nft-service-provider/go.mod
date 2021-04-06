module github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/ethereum/go-ethereum v1.9.18
	github.com/irisnet/service-sdk-go v1.0.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tidwall/gjson v1.6.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

go 1.14
