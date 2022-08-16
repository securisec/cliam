package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	gcpRestParent     map[string]string
	gcpRestBody       map[string]string
	gcpRestResourceID string
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
	gcpRestCmd.PersistentFlags().StringToStringVar(&gcpRestParent, "parent", map[string]string{}, "Specify the parent. i.e. project=my-project or organization=my-org. Valid keys are project, origanization, folder, billingAccount")
	gcpRestCmd.PersistentFlags().StringToStringVar(&gcpRestBody, "body", map[string]string{}, "Rest API body to use when not a GET request")
	gcpRestCmd.PersistentFlags().StringVar(&gcpRestResourceID, "resource-id", "", `Resource ID to use when not a GET request. Resource id is quite 
versitile and is used as a generic term for any resource. For example, if you are
enumerating pubsub, resource-id could be the subscription name; but also if you are 
enumerating compute, resource-id could be the instance name.`)

	// Completers
	gcpRestCmd.RegisterFlagCompletionFunc("parent", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{"project=", "organization=", "folder="}, cobra.ShellCompDirectiveNoSpace
	})
}
