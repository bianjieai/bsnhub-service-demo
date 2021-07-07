package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"bsn-irita-fabric-provider/appchains/fabric/iservice/chaincode/testChainCode"
)

func main() {

	err := shim.Start(new(testChainCode.TestChainCode))
	if err != nil {
		fmt.Printf("Error starting CrossChainCode: %s", err)
	}

}
