package cmd

import (
	"github.com/spf13/cobra"
)

var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Enumerate Azure credentials for their permissions.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(azureCmd)
}

// 🚧 TODO: Implement
