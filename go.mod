module github.com/bianjieai/bsnhub-service-demo

require (
	github.com/bianjieai/irita-sdk-go v1.0.0
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.5.1
)

replace (
	github.com/bianjieai/irita-sdk-go => /Users/bianjie/github.com/bianjieai/irita-sdk-go
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.4-irita-200703
)

go 1.14
