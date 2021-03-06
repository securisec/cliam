package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var gcpCmd = &cobra.Command{
	Use:   "gcp",
	Short: "Enumerate GCP service accounts for their permissions.",
	Long: `The following environment variables can be used to override the default behavior:
- GOOGLE_APPLICATION_CREDENTIALS
- CLOUDSDK_CORE_PROJECT`,
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
	gcpAccessToken        string
)

func init() {
	RootCmd.AddCommand(gcpCmd)
	gcpCmd.PersistentFlags().StringVar(&gcpServiceAccountPath, "service-account", "", "GCP service account path")
	gcpCmd.PersistentFlags().StringVar(&gcpProjectId, "project-id", "", "GCP project id")
	gcpCmd.PersistentFlags().StringVar(&gcpRegion, "region", "us-central1", "GCP Region")
	gcpCmd.PersistentFlags().StringVar(&gcpZone, "zone", "us-central1-a", "GCP Zone")
	gcpCmd.PersistentFlags().StringVar(&gcpAccessToken, "access-token", "", "GCP token")
	gcpCmd.RegisterFlagCompletionFunc("region", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp_regions, cobra.ShellCompDirectiveNoFileComp
	})
	gcpCmd.RegisterFlagCompletionFunc("zone", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp_zones, cobra.ShellCompDirectiveNoFileComp
	})
}

func getSaAndRegion() (string, string, string, string) {
	return getSaPath(), getProjectId(), gcpRegion, gcpZone
}

func getSaPath() string {
	if gcpAccessToken != "" {
		return ""
	}
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
