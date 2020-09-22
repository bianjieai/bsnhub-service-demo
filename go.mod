module github.com/bianjieai/bsnhub-service-demo

require (
	github.com/bianjieai/irita-sdk-go v1.0.0
	github.com/ethereum/go-ethereum v1.9.18
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.6.1
)

replace github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703

go 1.14
