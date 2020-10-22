#!/usr/bin/make -f

export GO111MODULE = on

install-contracts-service:
	@echo "installing bcos contracts service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/bcos-contracts-service-provider github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/cmd
