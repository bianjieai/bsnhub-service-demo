package server

import (
	"fmt"
)

// StartWebServer starts the web server with a ChainManager instance
func StartWebServer(
	chainManager *ChainManager,
	port int,
) {
	srv := NewHTTPService(chainManager)

	err := srv.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println(err)
	}
}
