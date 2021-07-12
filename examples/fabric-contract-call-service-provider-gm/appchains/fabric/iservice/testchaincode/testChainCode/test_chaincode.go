package testChainCode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (c *TestChainCode) getValue(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	result := args[0]
	return shim.Success([]byte(result))
}