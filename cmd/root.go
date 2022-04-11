package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "enumerate",
	Short: "Enumerate is a tool to enumerate AWS credentials for their permissions.",
	// Long: `Enumerate is a tool to enumerate AWS credentials for their permissions.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
