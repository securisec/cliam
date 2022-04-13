package cmd

import (
	"context"

	"github.com/securisec/iam-enumerate/aws"
	"github.com/securisec/iam-enumerate/aws/scanner"
	"github.com/securisec/iam-enumerate/aws/signer"
	"github.com/securisec/iam-enumerate/logger"
	"github.com/spf13/cobra"
)

var awsSingleCmd = &cobra.Command{
	Use:   "service [service]",
	Short: "Enumerate permissions for a single AWS service.",
	Run:   awsSingleCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetAWSServices(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	awsCmd.AddCommand(awsSingleCmd)
}

func awsSingleCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}

	key, secret, token, region := getCredsAndRegion()
	service := args[0]

	creds := signer.SetCredentials(key, secret, token)

	if err := scanner.EnumerateSpecificService(context.Background(), region, service, creds); err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("")
	}
}
