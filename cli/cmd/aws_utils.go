package cmd

import (
	"github.com/spf13/cobra"
)

var awsUtilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "AWS utilities",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	awsCmd.AddCommand(awsUtilsCmd)
}
