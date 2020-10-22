#!/usr/bin/make -f

export GO111MODULE = on

include examples/nft-service-provider/Makefile
include examples/bcos-store-service-provider/Makefile
include examples/bcos-contracts-service-provider/Makefile