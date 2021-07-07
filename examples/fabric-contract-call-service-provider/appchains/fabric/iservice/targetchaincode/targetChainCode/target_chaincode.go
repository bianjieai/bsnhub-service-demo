package targetChainCode

import (
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (c *TargetChainCode) callService(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	isReceived, err := stub.GetState(args[0])
	if(len(isReceived)>0){
		return shim.Error(fmt.Sprintf("the request has been received"))
	}
	var argAry []string
	err = json.Unmarshal([]byte(args[2]),&argAry)
	if err!= nil{
		return shim.Error("the callData cannot be Unmarshal")
	}
	chainCodeArgs := util.ToChaincodeArgs(argAry...)
	response := stub.InvokeChaincode(args[1],chainCodeArgs,args[3])
	if response.Status != shim.OK {
		return shim.Error(response.Message)
	}

	if err := stub.PutState(args[0], []byte("received")); err != nil {
		return shim.Error(fmt.Sprintf("put requestID info error；%s", err))
	}
	return response
}