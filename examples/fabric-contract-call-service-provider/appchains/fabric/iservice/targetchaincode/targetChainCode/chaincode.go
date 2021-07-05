package targetChainCode

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)


var successMsg = []byte("success")
var err_NoFunc = shim.Error("function not found")

type TargetChainCode struct {
}

func (c *TargetChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Init")

	return shim.Success(successMsg)
}

func (c *TargetChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Invoke")
	function, args := stub.GetFunctionAndParameters()

	if strings.ToLower(function) == "callservice" {
		return c.callService(stub, args)
	}
	return err_NoFunc
}
