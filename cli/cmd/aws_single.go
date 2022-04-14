package cmd

import (
	"context"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsServiceCmd = &cobra.Command{
	Use:   "service [service]",
	Short: "Enumerate permissions for a single AWS service.",
	Run:   awsServiceCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetAWSServices(), cobra.ShellCompDirectiveNoFileComp
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	awsCmd.AddCommand(awsServiceCmd)
}

func awsServiceCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}

	key, secret, token, region := getCredsAndRegion()
	service := args[0]

	creds := signer.SetCredentials(key, secret, token)

	if err := scanner.EnumerateSpecificResource(context.Background(), region, service, creds); err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("")
	}
}
