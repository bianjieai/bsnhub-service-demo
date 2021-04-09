package server

import (
	"fmt"
)

// StartWebServer starts the web server with a ChainManager instance
func StartWebServer(
	chainManager *ChainManager,
) {
	srv := NewHTTPService(chainManager)

	err := srv.Router.Run(":8082")
	if err != nil {
		fmt.Println(err)
	}
}
