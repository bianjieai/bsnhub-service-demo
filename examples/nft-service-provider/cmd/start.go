package main

import (
	"github.com/spf13/cobra"

	"github.com/bianjieai/bsnhub-service-demo/common"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/app"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/iservice"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/service"
)

func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start provider daemon",
		Example: `nft-service-provider start [config-file]`,
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

			service := service.MakeService(config)
			service.Logger = logger

			appInstance := app.NewApp(iserviceClient, service, logger)
			appInstance.Start()

			return nil
		},
	}

	return cmd
}
