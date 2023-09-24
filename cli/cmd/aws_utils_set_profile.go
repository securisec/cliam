package cmd

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

var awsUtilsSetProfileCmd = &cobra.Command{
	Use:   "set-profile",
	Short: "Set creds as aws credentials profile",
	Run:   awsUtilsSetProfileCmdFunc,
	// PreRun: awsLoadEnvVarsFirst,
}

func init() {
	awsUtilsCmd.AddCommand(awsUtilsSetProfileCmd)
	awsUtilsSetProfileCmd.Flags().StringP("profile-name", "n", "cliam", "Set a named profile")
}

func awsUtilsSetProfileCmdFunc(cmd *cobra.Command, _ []string) {
	name, _ := cmd.Flags().GetString("profile-name")
	// get credentials
	key, secret, token, _ := getCredsAndRegion()

	if err := awsSaveCliamProfile(name, key, secret, token); err != nil {
		logger.LogError(err)
	}

	logger.LoggerStdErr.Info().Msg("Saved profile " + name)
}

func awsReadAWSCredentialsFile() (*ini.File, string, error) {
	cc := credentials.SharedCredentialsProvider{}
	if _, err := cc.Retrieve(); err != nil {
		return nil, "", err
	}
	config, err := ini.Load(cc.Filename)
	return config, cc.Filename, err
}

func awsSaveCliamProfile(name, key, secret, token string) error {
	config, configFile, err := awsReadAWSCredentialsFile()
	if err != nil {
		return err
	}
	ns, err := config.NewSection(name)
	if err != nil {
		return err
	}
	ns.Key("aws_access_key_id").SetValue(key)
	ns.Key("aws_secret_access_key").SetValue(secret)
	if token != "" {
		ns.Key("aws_session_token").SetValue(token)
	}
	return config.SaveTo(configFile)
}
