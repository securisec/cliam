package cmd

import (
	"context"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsMultipleCmd = &cobra.Command{
	Use:   "enumerate [resource...]",
	Short: "Enumerate permissions for specified AWS resources.",
	Run:   awsMultipleCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetAWSResources(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	awsCmd.AddCommand(awsMultipleCmd)
}

func awsMultipleCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}

	key, secret, token, region := getCredsAndRegion()
	services := removeDuplicates(args)

	creds := signer.SetCredentials(key, secret, token)

	if err := scanner.EnumerateMultipleResources(context.Background(), region, services, creds); err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("")
	}
}
