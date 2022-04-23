package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/securisec/cliam/aws/scanner"
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
	awsAccessKeyID       string
	awsSecretAccessKey   string
	awsSessionToken      string
	awsRegion            string
	awsProfile           string
	awsSessionJson       string
	awsKnownResourceName string
	awsKnownOnly         bool
)

func init() {
	RootCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVar(&awsAccessKeyID, "access-key-id", "", "AWS Access Key ID")
	awsCmd.PersistentFlags().StringVar(&awsSecretAccessKey, "secret-access-key", "", "AWS Secret Access Key")
	awsCmd.PersistentFlags().StringVar(&awsSessionToken, "session-token", "", "AWS Session Token")
	awsCmd.PersistentFlags().StringVar(&awsRegion, "region", "us-east-1", "AWS Region")
	awsCmd.PersistentFlags().StringVar(&awsProfile, "profile", "", "AWS Profile. When profile is set, access-key-id, secret-access-key, and session-token are ignored.")
	awsCmd.PersistentFlags().StringVar(&awsSessionJson, "session-json", "", "AWS Session JSON file. This flag attempts to read session information from the specified file. Helpful with temporary credentials.")
	awsCmd.PersistentFlags().StringVar(&awsKnownResourceName, "known-resource-name", "", "AWS Resource Name. When known-resource-name is set, additional permissions where a resource needs to be specified is enumerated.")
	awsCmd.PersistentFlags().BoolVar(&awsKnownOnly, "known-only", false, "When set, only permissions where the known-resource-name is specified are enumerated.")
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

func awsSendToChannel(ch chan scanner.ServiceMap, resources []string, word string) {
	var extras []scanner.ServiceMap
	enumerate := scanner.GetServiceMap(resources)
	for _, e := range enumerate {
		// for now, move the permissions that require a resource to the extras
		// this will then be looped later if known-resource-name is set
		if e.Policy.IsExtra {
			extras = append(extras, e)
		} else {
			// if the known only flag is set, ignore general permissions that doesnt require a resource name
			if !awsKnownOnly {
				ch <- e
			}
		}
	}

	// if a known resource name is set, we will enumerate the extra permissions
	if word != "" && len(extras) > 0 {
		for _, ee := range extras {
			ch <- ee
		}
	}
}

type awsSessionJsonStruct struct {
	AccessKeyId     string `json:"AccessKeyId"`
	SecretAccessKey string `json:"SecretAccessKey"`
	Token           string `json:"Token"`
}
