package main

import (
	"github.com/spf13/cobra"

	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/app"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/common"
	contractsservice "github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/contracts-service"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/iservice"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `bcos-contracts-service-provider start [config-file]`,
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

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			contractsService := contractsservice.MakeContractsService(config)
			contractsService.Logger = logger

			appInstance := app.NewApp(iserviceClient, contractsService, logger)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
