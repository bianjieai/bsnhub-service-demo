package main

import (
	"io/ioutil"

	"github.com/spf13/cobra"

	"github.com/bianjieai/bsnhub-service-demo/common"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/app"
	contractsservice "github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/contracts-service"
	"github.com/bianjieai/bsnhub-service-demo/examples/bcos-contracts-service-provider/iservice"
)

func DeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy iservice according to the metadata",
		Example: `bcos-contracts-service-provider deploy [config-file]`,
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

			schemas, err := ioutil.ReadFile("./metadata/service-definition.json")
			if err != nil {
				return err
			}

			pricing, err := ioutil.ReadFile("./metadata/service-binding.json")
			if err != nil {
				return err
			}

			logger := common.Logger

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			contractsService := contractsservice.MakeContractsService(config)
			contractsService.Logger = logger

			appInstance := app.NewApp(iserviceClient, contractsService, logger)

			err = appInstance.DeployIService(string(schemas), string(pricing))
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
