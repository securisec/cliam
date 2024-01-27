package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var gcpResourceManagerCmd = &cobra.Command{
	Use:   "resourcemanager",
	Short: "GCP permissions using the resourcemanager API",
	Long: `Enumerate uses cloudresource manager to check if the service account has the specified permissions. 
If resource manager is disabled, this command will end with a 403 error.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

func init() {
	gcpCmd.AddCommand(gcpResourceManagerCmd)
}
