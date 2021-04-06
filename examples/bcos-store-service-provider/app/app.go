package app

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/iservice"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/store"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-store-service-provider/types"
)

// App represents the provider application
type App struct {
	IServiceClient iservice.ServiceClientWrapper
	StoreService   store.StoreService
	Logger         *log.Logger
}

// NewApp constructs a new App instance
func NewApp(
	iserviceClient iservice.ServiceClientWrapper,
	storeService store.StoreService,
	logger *log.Logger,
) App {
	return App{
		IServiceClient: iserviceClient,
		StoreService:   storeService,
		Logger:         logger,
	}
}

// Start starts the provider process
func (app App) Start() {
	app.Logger.Infof("app started")

	err := app.IServiceClient.SubscribeServiceRequest(
		types.ServiceName,
		app.StoreService.Callback,
	)
	if err != nil {
		app.Logger.Errorf("failed to register service request listener, err: %s", err.Error())
		return
	}

	select {}
}

// DeployIService deploys the iservice according to the given metadata
func (app App) DeployIService(schemas string, pricing string) error {
	app.Logger.Infof("starting to deploy %s service", types.ServiceName)

	_, err := app.IServiceClient.ServiceClient.QueryServiceDefinition(types.ServiceName)
	if err != nil {
		app.Logger.Infof("defining service")

		err := app.IServiceClient.DefineService(types.ServiceName, "", "", nil, schemas)
		if err != nil {
			return fmt.Errorf("failed to define service: %s", err.Error())
		}
	} else {
		app.Logger.Infof("service definition already exists")
	}

	_, provider, err2 := app.IServiceClient.ServiceClient.Find(app.IServiceClient.KeyName, app.IServiceClient.Passphrase)
	if err2 != nil {
		return err2
	}

	_, err = app.IServiceClient.ServiceClient.QueryServiceBinding(types.ServiceName, provider.String())
	if err != nil {
		app.Logger.Infof("binding service")

		err := app.IServiceClient.BindService(types.ServiceName, "100000point", pricing, "{}", 100)
		if err != nil {
			return fmt.Errorf("failed to bind service: %s", err.Error())
		}
	} else {
		app.Logger.Infof("service binding already exists")
	}

	app.Logger.Infof("%s service deployment completed", types.ServiceName)

	return nil
}
