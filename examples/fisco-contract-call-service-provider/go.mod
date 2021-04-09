module github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider

require (
	github.com/FISCO-BCOS/go-sdk v0.10.1
	github.com/cockroachdb/pebble v0.0.0-20210406003833-3d4c32f510a8
	github.com/ethereum/go-ethereum v1.9.18
	github.com/gin-gonic/gin v1.7.1
	github.com/irisnet/service-sdk-go v1.0.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/tidwall/gjson v1.6.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

go 1.14
