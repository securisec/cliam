package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Enumerate AWS credentials for their permissions.",
	Long: `The following environment variables can be used to override the default values:
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_SESSION_TOKEN`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		// awsLoadEnvVarsFirst(cmd, args)
	},
}

var (
	awsAccessKeyID      string
	awsSecretAccessKey  string
	awsSessionToken     string
	awsRegion           string
	awsProfile          string
	awsSessionJson      string
	awsEndpoint         string
	awsKnownResourceMap []string
	// awsKnownOnly         bool
)

func init() {
	RootCmd.AddCommand(awsCmd)
	awsCmd.PersistentFlags().StringVar(&awsAccessKeyID, "access-key-id", os.Getenv("AWS_ACCESS_KEY_ID"), "AWS Access Key ID")
	awsCmd.PersistentFlags().StringVar(&awsSecretAccessKey, "secret-access-key", os.Getenv("AWS_SECRET_ACCESS_KEY"), "AWS Secret Access Key")
	awsCmd.PersistentFlags().StringVar(&awsSessionToken, "session-token", os.Getenv("AWS_SESSION_TOKEN"), "AWS Session Token")
	awsCmd.PersistentFlags().StringVar(&awsRegion, "region", "us-east-1", "AWS Region")
	awsCmd.PersistentFlags().StringVar(&awsProfile, "profile", "", "AWS Profile. When profile is set, access-key-id, secret-access-key, and session-token are ignored.")
	awsCmd.PersistentFlags().StringVar(&awsSessionJson, "session-json", "", "AWS Session JSON file. This flag attempts to read session information from the specified file. Helpful with temporary credentials.")
	awsCmd.PersistentFlags().StringVar(&awsEndpoint, "endpoint", "", "AWS Endpoint. Custom AWS endpoint.")
	awsCmd.PersistentFlags().StringSliceVarP(&awsKnownResourceMap, "known-value", "k", []string{}, "AWS Resource Name. Maps directly with aws cli flags. This flag can be used multiple times.")
	awsCmd.RegisterFlagCompletionFunc("region", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws_Regions, cobra.ShellCompDirectiveNoFileComp
	})
	// awsCmd.RegisterFlagCompletionFunc("known-value", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 	return []string{}, cobra.ShellCompDirectiveNoFileComp
	// })
	awsCmd.RegisterFlagCompletionFunc("known-value", func(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
		var keys []string
		policies := scanner.GetServiceMap(args)
		for _, p := range policies {
			if len(p.Policy.ExtraCommandLineFlag) > 0 {
				keys = append(keys, p.Policy.ExtraCommandLineFlag+"=")
			}
		}
		return keys, cobra.ShellCompDirectiveNoSpace
	})
}

// return the key, secret, token and region
func getCredsAndRegion() (string, string, string, string) {
	if awsSessionJson != "" {
		s, err := awsReadSessionJsonFile()
		if err != nil {
			logger.LoggerStdErr.Fatal().Msg("Failed to read session json file")
		}
		return s.Credentials.AccessKeyId, s.Credentials.SecretAccessKey, s.Credentials.Token, awsRegion
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
	if envar == "AWS_ACCESS_KEY_ID" && awsAccessKeyID != "" {
		return awsAccessKeyID
	}
	if envar == "AWS_SECRET_ACCESS_KEY" && awsSecretAccessKey != "" {
		return awsSecretAccessKey
	}
	if k, ok := os.LookupEnv(envar); ok {
		return k
	}
	return promptInput(message)
}

func awsReadSessionJsonFile() (awsSessionJsonStruct, error) {
	var s awsSessionJsonStruct
	o, err := os.ReadFile(awsSessionJson)
	if err != nil {
		return s, err
	}
	err = json.Unmarshal(o, &s)
	return s, err
}

func awsSendToChannel(ch chan scanner.ServiceMap, resources []string) {
	var extras []scanner.ServiceMap
	enumerate := scanner.GetServiceMap(resources)
	for _, e := range enumerate {
		// for now, move the permissions that require a resource to the extras
		// this will then be looped later if known-resource-name is set
		if e.Policy.IsExtra {
			extras = append(extras, e)
		} else {
			// if the known only flag is set, ignore general permissions that doesnt require a resource name
			if len(awsKnownResourceMap) == 0 {
				ch <- e
			}
		}
	}

	// if a known resource name is set, we will enumerate only the extra permissions
	extraFlags := ModifyExtraMap(awsKnownResourceMap)
	if len(extraFlags) > 0 && len(extras) > 0 {
		for _, ee := range extras {
			ee.Policy.ExtraValueMap = awsModifyExtraMap(extraFlags)
			ch <- ee
		}
	}
}

type awsSessionJsonStruct struct {
	Credentials struct {
		AccessKeyId     string `json:"AccessKeyId"`
		SecretAccessKey string `json:"SecretAccessKey"`
		Token           string `json:"SessionToken"`
		Expiration      string `json:"Expiration"`
	} `json:"Credentials"`
}

func awsModifyExtraMap(m map[string]string) map[string]string {
	h := map[string]string{}
	for k, v := range m {
		h[strings.ReplaceAll(k, "-", "_")] = v
	}
	return h
}

// func awsLoadEnvVarsFirst(_ *cobra.Command, _ []string) {
// 	awsSessionToken = os.Getenv("AWS_SESSION_TOKEN")
// 	awsAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
// 	awsSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
// }
