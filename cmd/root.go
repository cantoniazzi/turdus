// Package cmd is the package of the CLI resources
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the CLI entrypoint
var RootCmd = &cobra.Command{
	Use:   "turdus",
	Short: "Turdus a tweet search app",
	Long:  "Turdus a tweet search app",
}

// Execute CLI commands
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
