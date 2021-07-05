package testChainCode

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


var successMsg = []byte("success")
var err_NoFunc = shim.Error("function not found")

type TestChainCode struct {
}

func (c *TestChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Init")

	return shim.Success(successMsg)
}

func (c *TestChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Invoke")
	function, args := stub.GetFunctionAndParameters()

	if strings.ToLower(function) == "getvalue" {
		return c.getValue(stub, args)
	}
	return err_NoFunc
}
