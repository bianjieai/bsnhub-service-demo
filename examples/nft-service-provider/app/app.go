package app

import (
	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/iservice"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/service"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/types"
)

// App represents the provider application
type App struct {
	IServiceClient iservice.ServiceClientWrapper
	Service        service.Service
	Logger         *log.Logger
}

// NewApp constructs a new App instance
func NewApp(
	iserviceClient iservice.ServiceClientWrapper,
	service service.Service,
	logger *log.Logger,
) App {
	return App{
		IServiceClient: iserviceClient,
		Service:        service,
		Logger:         logger,
	}
}

// Start starts the provider process
func (app App) Start() {
	app.Logger.Infof("app started")

	err := app.IServiceClient.SubscribeServiceRequest(
		types.ServiceName,
		app.Service.Callback,
	)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}

	select {}
}

// DeployIService deploys the iservice according to the given metadata
func (app App) DeployIService(definition string, binding string) error {
	// TODO
	return nil
}
