package mysql

import (
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
)

// OnServiceRequestReceived is the hook which is called when the service request is received
func OnServiceRequestReceived(requestID, toChainID string) {
	err := update("to_chainid", toChainID, requestID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}

// OnContractTxSend is the hook which is called when the request is sent to target chain
func OnContractTxSend(requestID, toTx string) {
	err := update("to_tx", toTx, requestID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}

// OnInterchainResponseSent is the hook which is called when the response is sent to hub
func OnInterchainResponseSent(icResID, hubResTx string) {
	err := update("hub_res_tx", hubResTx, icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}

func TxErrCollection(requestID, errStr string){
	err := update("error", errStr, requestID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}
