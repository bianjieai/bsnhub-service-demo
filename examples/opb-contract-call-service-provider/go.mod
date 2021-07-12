module github.com/bianjieai/bsnhub-service-demo/examples/opb-contract-call-service-provider

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/bianjieai/irita-sdk-go v1.1.1-0.20210707070124-79ed0124b3de
	github.com/cockroachdb/pebble v0.0.0-20210406003833-3d4c32f510a8
	github.com/gin-gonic/gin v1.7.1
	github.com/go-sql-driver/mysql v1.4.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace (
	github.com/bianjieai/irita-sdk-go => /Users/bianjie/irita/irita-sdk-go
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)

go 1.14
