package cmd

import (
	"encoding/json"
	"os"

	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
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
	gcpCmd.PersistentFlags().StringVar(&gcpServiceAccountPath, "service-account", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"), "GCP service account path. Env GOOGLE_APPLICATION_CREDENTIALS")
	gcpCmd.PersistentFlags().StringVar(&gcpProjectId, "project-id", os.Getenv("CLOUDSDK_CORE_PROJECT"), "GCP project id. Env CLOUDSDK_CORE_PROJECT")
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
		d, err := gcpReadServiceAccount(gcpServiceAccountPath)
		if err == nil {
			gcpProjectId = d.ProjectID
		}
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

type gcpServiceAccountType struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CERTURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CERTURL       string `json:"client_x509_cert_url"`
}

func gcpReadServiceAccount(path string) (*gcpServiceAccountType, error) {
	sa := &gcpServiceAccountType{}
	o, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(o, sa); err != nil {
		return nil, err
	}
	if CLIVerbose {
		logger.LoggerStdErr.Debug().Str("email", sa.ClientEmail).Msg(shared.GetMessageColor("info"))
	}
	gcpProjectId = sa.ProjectID
	return sa, nil
}
