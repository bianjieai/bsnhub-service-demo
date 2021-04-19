package main

import (
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/mysql"
	"github.com/spf13/cobra"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/app"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/common"
	contractservice "github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/contract-service"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/iservice"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/store"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-call-service-provider/server"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `fisco-contract-call-sp start [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 0 {
				configFileName = common.DefaultConfigFileName
			} else {
				configFileName = args[0]
			}

			config, err := common.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			logger := common.Logger

			store, err := store.NewStore(config.GetString(common.ConfigKeyStorePath))
			if err != nil {
				return err
			}
			chainManager := server.NewChainManager(store)

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			contractService, err := contractservice.BuildContractService(config, chainManager)
			if err != nil {
				return err
			}

			contractService.Logger = logger
			appInstance := app.NewApp(iserviceClient, contractService, logger)

			mysqlConfig := mysql.NewConfig(config)
			mysql.NewDB(mysqlConfig)
			defer mysql.DB.Close()

			go server.StartWebServer(chainManager)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
