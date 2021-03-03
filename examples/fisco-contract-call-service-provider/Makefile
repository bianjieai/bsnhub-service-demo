#!/usr/bin/make -f

export GO111MODULE = on

install:
	@echo "installing fisco contract call service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/fisco-contract-call-sp github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/cmd
