package cmd

import (
	"github.com/spf13/cobra"
)

var digitaloceanCmd = &cobra.Command{
	Use:   "digitalocean",
	Short: "Enumerate Azure credentials for their permissions.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	RootCmd.AddCommand(digitaloceanCmd)
}

// 🚧 TODO: Implement
