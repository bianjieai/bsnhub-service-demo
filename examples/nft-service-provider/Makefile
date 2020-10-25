#!/usr/bin/make -f

export GO111MODULE = on

install-nft:
	@echo "installing nft service provider..."
	@go build -mod=readonly -o $${GOBIN-$${GOPATH-$$HOME/go}/bin}/nft-service-provider github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/cmd
