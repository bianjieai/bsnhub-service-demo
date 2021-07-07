package server

import (
	"bsn-irita-fabric-provider/appchains"
	"fmt"
)

func StartWebServer(
	chainManager appchains.AppChainHandlerI,
	port int,
) {
	srv := NewHTTPService(chainManager)

	err := srv.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
	}
}
