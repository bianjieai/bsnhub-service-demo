package main

import (
	"github.com/spf13/cobra"
)

func DeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deploy",
		Short:   "Deploy iservice according to the metadata",
		Example: `nft-service-provider deploy [config-file]`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO
			return nil
		},
	}

	return cmd
}
