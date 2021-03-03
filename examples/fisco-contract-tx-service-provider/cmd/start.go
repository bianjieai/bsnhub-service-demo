package main

import (
	"github.com/spf13/cobra"

	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-tx-service-provider/app"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-tx-service-provider/common"
	contractservice "github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-tx-service-provider/contract-service"
	"github.com/bianjieai/bsnhub-service-demo/examples/fisco-contract-tx-service-provider/iservice"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `fisco-contract-tx-sp start [config-file]`,
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

			contractService, err := contractservice.BuildContractService(config)
			if err != nil {
				return err
			}

			contractService.Logger = logger

			appInstance := app.NewApp(iserviceClient, contractService, logger)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
