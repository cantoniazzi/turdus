// Package cmd is the package of the CLI resources
package cmd

import (
	"github.com/cantoniazzi/turdus/api"
	"github.com/spf13/cobra"
)

var apiCommand = &cobra.Command{
	Use:   "api",
	Short: "Starting Turdus API",
	Long:  "Starting Turdus API",
	RunE:  apiExecute,
}

func init() {
	RootCmd.AddCommand(apiCommand)
}

func apiExecute(cmd *cobra.Command, args []string) error {
	api.Start()
	return nil
}
