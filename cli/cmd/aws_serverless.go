package cmd

import (
	"context"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsServerlessCmd = &cobra.Command{
	Use:   "serverless",
	Short: "Enumerate permissions for common serverless AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate apigateway, lambda, 
s3, sns, sqs and dynamodb.`,
	Run:               awsServerlessCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsServerlessCmd)
	awsServerlessCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsServerlessCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	key, secret, token, region := getCredsAndRegion()
	services := []string{
		aws.APIGateway,
		aws.Lambda,
		aws.S3,
		aws.SNS,
		aws.SQS,
		aws.Dynamodb,
	}

	creds := signer.SetCredentials(key, secret, token)

	if err := scanner.EnumerateMultipleResources(context.Background(), region, services, creds, MaxThreads, saveOutput); err != nil {
		logger.LoggerStdErr.Err(err).Msg("")
	}
}
