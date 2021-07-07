package mysql

import (
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
)

// OnServiceRequestReceived is the hook which is called when the service request is received
func OnServiceRequestReceived(icResID, toChainID string) {
	err := insert("ic_request_id", icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
	err = update("to_chainid", toChainID, icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}

	err = updateTime("tx_createtime", icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}

// OnContractTxSend is the hook which is called when the request is sent to target chain
func OnContractTxSend(icResID, toTx string) {
	err := update("to_tx", toTx, icResID)
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
	err = update("tx_status", "1", icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}

func TxErrCollection(icResID, errStr string) {
	err := update("error", errStr, icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
	err = update("tx_status", "2", icResID)
	if err != nil {
		common.Logger.Errorf(err.Error())
		return
	}
}
