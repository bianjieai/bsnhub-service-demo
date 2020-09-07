package main

import (
	"os"
	"path/filepath"

	"github.com/bianjieai/bsnhub-service-demo/iservice/market"
	"github.com/bianjieai/bsnhub-service-demo/iservice/node"
	"github.com/bianjieai/irita-sdk-go/types/store"

	"github.com/bianjieai/irita-sdk-go/types"
	"github.com/spf13/cobra"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(startCmd())
}

var rootCmd = &cobra.Command{
	Use:          "iservice",
	Short:        "iservice daemon",
	SilenceUsage: true,
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "start daemon",
		Example: `iservice start [chain-id] [node-uri] [key_name] [password] [market]`,
		Args:    cobra.RangeArgs(4, 5),
		RunE: func(cmd *cobra.Command, args []string) error {
			// set market
			if len(args) > 4 {
				market.MarketType = args[4]
			}

			keyDao := store.NewFileDAO(keysPath)

			options := []types.Option{
				types.KeyDAOOption(keyDao),
				types.TimeoutOption(10),
			}

			cfg, err := types.NewClientConfig(args[1], args[0], options...)
			if err != nil {
				panic(err)
			}

			baseTx := types.BaseTx{
				From:     args[2],
				Password: args[3],
			}
			node.Start(cfg, baseTx)
			return nil
		},
	}
	return cmd
}

var (
	keysPath = os.ExpandEnv(filepath.Join("$HOME", ".iritacli"))
)
