package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Enumerate AWS credentials for their permissions.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

var (
	awsAccessKeyID     string
	awsSecretAccessKey string
	awsSessionToken    string
	awsRegion          string
)

func init() {
	RootCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVar(&awsAccessKeyID, "access-key-id", "", "AWS Access Key ID")
	awsCmd.PersistentFlags().StringVar(&awsSecretAccessKey, "secret-access-key", "", "AWS Secret Access Key")
	awsCmd.PersistentFlags().StringVar(&awsSessionToken, "session-token", "", "AWS Session Token")
	awsCmd.PersistentFlags().StringVar(&awsRegion, "region", "us-east-1", "AWS Region")
}

func getCredsAndRegion() (string, string, string, string) {
	return getKeyID(), getSecretKey(), getToken(), awsRegion
}

func getKeyID() string {
	if awsAccessKeyID != "" {
		return awsAccessKeyID
	}
	if k, ok := os.LookupEnv("AWS_ACCESS_KEY_ID"); ok {
		return k
	}
	return promptInput("AWS Access Key ID: ")
}

func getSecretKey() string {
	if awsSecretAccessKey != "" {
		return awsSecretAccessKey
	}
	if k, ok := os.LookupEnv("AWS_SECRET_ACCESS_KEY"); ok {
		return k
	}
	return promptInput("AWS Secret Access Key: ")
}

func getToken() string {
	if awsSessionToken != "" {
		return awsSessionToken
	}
	if k, ok := os.LookupEnv("AWS_SESSION_TOKEN"); ok {
		return k
	}
	return ""
}
