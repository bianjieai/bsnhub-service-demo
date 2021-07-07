package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the entry point
var (
	rootCmd = &cobra.Command{
		Use:   "nft-service-provider",
		Short: "NFT service provider daemon command line interface",
	}
)

func main() {
	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(StartCmd())
	rootCmd.AddCommand(DeployCmd())
	rootCmd.AddCommand(KeysCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
