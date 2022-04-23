package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Enumerate AWS credentials for their permissions.",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
}

var (
	awsAccessKeyID     string
	awsSecretAccessKey string
	awsSessionToken    string
	awsRegion          string
	awsProfile         string
	awsSessionJson     string
)

func init() {
	RootCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVar(&awsAccessKeyID, "access-key-id", "", "AWS Access Key ID")
	awsCmd.PersistentFlags().StringVar(&awsSecretAccessKey, "secret-access-key", "", "AWS Secret Access Key")
	awsCmd.PersistentFlags().StringVar(&awsSessionToken, "session-token", "", "AWS Session Token")
	awsCmd.PersistentFlags().StringVar(&awsRegion, "region", "us-east-1", "AWS Region")
	awsCmd.PersistentFlags().StringVar(&awsProfile, "profile", "", "AWS Profile. When profile is set, access-key-id, secret-access-key, and session-token are ignored.")
	awsCmd.PersistentFlags().StringVar(&awsSessionJson, "session-json", "", "AWS Session JSON file. This flag attempts to read session information from the specified file. Helpful with temporary credentials.")
}

// return the key, secret, token and region
func getCredsAndRegion() (string, string, string, string) {
	if awsSessionJson != "" {
		s, err := awsReadSessionJsonFile()
		if err != nil {
			logger.LoggerStdErr.Fatal().Msg("Failed to read session json file")
		}
		return s.AccessKeyId, s.SecretAccessKey, s.Token, awsRegion
	}
	return awsGetEnvarOrPrompt("AWS_ACCESS_KEY_ID", "AWS Access Key ID: "),
		awsGetEnvarOrPrompt("AWS_SECRET_ACCESS_KEY", "AWS Secret Access Key: "),
		awsSessionToken,
		awsRegion
}

func awsGetEnvarOrPrompt(envar, message string) string {
	if awsProfile != "" {
		// profile is set so we will use the profile
		return ""
	}
	if awsAccessKeyID != "" {
		return awsAccessKeyID
	}
	if k, ok := os.LookupEnv(envar); ok {
		return k
	}
	return promptInput(message)
}

func awsReadSessionJsonFile() (awsSessionJsonStruct, error) {
	var s awsSessionJsonStruct
	o, err := ioutil.ReadFile(awsSessionJson)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(o, &s)
	return s, err
}

type awsSessionJsonStruct struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
}
