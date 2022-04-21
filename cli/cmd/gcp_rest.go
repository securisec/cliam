package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var gcpRestCmd = &cobra.Command{
	Use:   "rest [--parent project=someproject] [resource...]",
	Short: "GCP permissions using the REST API",
	Long:  `Unlike enumerate, this command uses the REST API to enumerate permissions.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	gcpCmd.AddCommand(gcpRestCmd)
}
