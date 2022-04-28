package cmd

import (
	"github.com/spf13/cobra"
)

var firebaseCmd = &cobra.Command{
	Use:   "firebase",
	Short: "Enumerate Firebase permissions.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var (
	firebaseProjectId string
	firebaseExtras    map[string]string
)

func init() {
	RootCmd.AddCommand(firebaseCmd)
	firebaseCmd.PersistentFlags().StringVar(&firebaseProjectId, "project-id", "", "Firebase project id")
	firebaseCmd.PersistentFlags().StringToStringVar(&firebaseExtras, "extra", map[string]string{}, "Extra parameters like collection name")
	firebaseCmd.MarkPersistentFlagRequired("project-id")
}
