package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cliam",
	Short: "Cloud Enumerate is a tool to enumerate cloud credentials for their permissions.",
	// Long: `Enumerate is a tool to enumerate AWS credentials for their permissions.`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
