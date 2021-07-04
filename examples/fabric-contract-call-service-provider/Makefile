#!/usr/bin/make -f

build: go.sum
ifeq ($(OS),Windows_NT)
	@echo "building relayer..."
	@go build -mod=readonly -o build/ cmd
else
	@echo "building relayer..."
	@go build -mod=readonly -o build/ cmd
endif

install: go.sum
	@echo "installing provider..."
	@go build -mod=readonly -o build/irita-fabric-provider ./cmd

