package main

import (
	"bsn-irita-fabric-provider/appchains/fabric/iservice/targetchaincode/targetChainCode"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func main() {

	err := shim.Start(new(targetChainCode.TargetChainCode))
	if err != nil {
		fmt.Printf("Error starting CrossChainCode: %s", err)
	}

}
