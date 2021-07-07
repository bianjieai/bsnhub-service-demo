package main

import (
	"bsn-irita-fabric-provider/appchains"
	"bsn-irita-fabric-provider/iservice"
	"bsn-irita-fabric-provider/server"
	"github.com/spf13/cobra"

	"bsn-irita-fabric-provider/common"
)

const (
	_Service_name = "service.service_name"
	_HttpPort     = "base.http_port"
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

			//todo 测试
			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			appChain := appchains.NewAppChainHandler(config.GetString(_Service_name), config, &iserviceClient, common.Logger)

			err = appChain.Start()

			if err != nil {
				common.Logger.Errorf("the appChain start failed %v", err)
				return err
			}

			httpPort := config.GetInt(_HttpPort)
			if httpPort == 0 {
				httpPort = 18051
			}

			server.StartWebServer(appChain, httpPort)

			return nil
		},
	}

	return cmd
}
