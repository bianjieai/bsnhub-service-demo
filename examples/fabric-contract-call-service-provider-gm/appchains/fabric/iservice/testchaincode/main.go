package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"bsn-irita-fabric-provider-gm/appchains/fabric/iservice/testchaincode/testChainCode"
)

func main() {

	err := shim.Start(new(testChainCode.TestChainCode))
	if err != nil {
		fmt.Printf("Error starting CrossChainCode: %s", err)
	}

}
