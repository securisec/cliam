package cmd

import (
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsCommonCmd = &cobra.Command{
	Use:       "service-group",
	Short:     "Enumerate permissions for groups of AWS resources.",
	Long:      "For example, serverless group enumerates permissions for lambda, sqs, s3, etc",
	Run:       awsCommonCmdFunc,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: getAwsServiceGroups(),
	// PreRun:    awsLoadEnvVarsFirst,
	PostRun: PostRunStatsFunc,
}

func init() {
	awsCmd.AddCommand(awsCommonCmd)
}

func awsCommonCmdFunc(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		logger.LoggerStdErr.Fatal().Msg("Invalid number of arguments")
	}

	resources, ok := AWS_SERVICE_GROUPING[args[0]]
	if !ok {
		logger.LoggerStdErr.Fatal().Msg("Invalid service group")
	}

	awsSharedEnumerate(resources, SaveOutput)
}
