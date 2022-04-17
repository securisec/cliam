package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Enumerate GCP service accounts for their permissions.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var (
	gcpServiceAccountPath string
	gcpProjectId          string
	gcpRegion             string
	gcpZone               string
)

func init() {
	RootCmd.AddCommand(gcpCmd)
	gcpCmd.PersistentFlags().StringVar(&gcpServiceAccountPath, "service-account", "", "GCP service account path")
	gcpCmd.PersistentFlags().StringVar(&gcpProjectId, "project-id", "", "GCP project id")
	gcpCmd.PersistentFlags().StringVar(&gcpRegion, "region", "us-central1", "GCP Region")
	gcpCmd.PersistentFlags().StringVar(&gcpZone, "zone", "us-central1-a", "GCP Zone")
}

func getSaAndRegion() (string, string, string, string) {
	return getSaPath(), getProjectId(), gcpRegion, gcpZone
}

func getSaPath() string {
	if gcpServiceAccountPath != "" {
		return expandPath(gcpServiceAccountPath)
	}
	if k, ok := os.LookupEnv("GOOGLE_APPLICATION_CREDENTIALS"); ok {
		return k
	}
	return expandPath(promptInput("GCP service account path: "))
}

func getProjectId() string {
	if gcpProjectId != "" {
		return gcpProjectId
	}
	if k, ok := os.LookupEnv("CLOUDSDK_CORE_PROJECT"); ok {
		return k
	}
	return promptInput("GCP project id: ")
}
