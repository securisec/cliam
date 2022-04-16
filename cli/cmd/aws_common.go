package cmd

import (
	"context"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsCommonCmd = &cobra.Command{
	Use:   "common",
	Short: "Enumerate permissions for common AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate apigateway, lambda, 
s3, iam, and ec2`,
	Run:               awsCommonCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsCommonCmd)
	awsCommonCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsCommonCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	key, secret, token, region := getCredsAndRegion()
	services := []string{
		aws.Lambda,
		aws.IAM,
		aws.S3,
		aws.APIGateway,
		aws.EC2,
	}

	creds := signer.SetCredentials(key, secret, token)

	if err := scanner.EnumerateMultipleResources(context.Background(), region, services, creds, MaxThreads, &saveOutput); err != nil {
		logger.LoggerStdErr.Err(err).Msg("")
	}
}
