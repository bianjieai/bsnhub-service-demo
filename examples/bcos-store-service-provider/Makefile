#!/usr/bin/make -f

export GO111MODULE = on

install-store:
	@echo "installing bcos store service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/bcos-store-service-provider github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/cmd
