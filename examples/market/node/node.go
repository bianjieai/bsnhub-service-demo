package node

import (
	"fmt"
	"log"

	"github.com/bianjieai/bsnhub-service-demo/examples/market/service"
	sdk "github.com/bianjieai/irita-sdk-go"
	"github.com/bianjieai/irita-sdk-go/types"
)

func Start(config types.ClientConfig, baseTx types.BaseTx) {
	client := sdk.NewIRITAClient(config)
	// client.SetOutput(os.Stdout)
	serviceName := service.PriceServiceName
	baseTx.Memo = fmt.Sprintf("service invocation response for %s", serviceName)
	_, err := client.Service.SubscribeServiceRequest(
		serviceName, service.GetServiceCallBack(serviceName), baseTx)
	if err != nil {
		log.Printf("failed to register invocation listener, err: %s", err.Error())
		return
	}

	select {}
}
