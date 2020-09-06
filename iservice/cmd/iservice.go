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
		Example: `iservice start [key_name] [market]`,
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			// set market
			if len(args) > 1 {
				market.MarketType = args[1]
			}

			keyDao := store.NewFileDAO(keysPath)

			options := []types.Option{
				types.KeyDAOOption(keyDao),
				types.TimeoutOption(10),
			}

			cfg, err := types.NewClientConfig(NodeURI, ChainID, options...)
			cfg.Level = "debug"
			if err != nil {
				panic(err)
			}

			baseTx := types.BaseTx{
				From:     args[0],
				Password: "1234567890",
			}
			node.Start(cfg, baseTx)
			return nil
		},
	}
	return cmd
}

const (
	NodeURI = "tcp://localhost:26657"
	ChainID = "test"
)

var (
	keysPath = os.ExpandEnv(filepath.Join("$HOME", ".iritacli"))
)
