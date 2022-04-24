package cmd

import (
	"github.com/securisec/cliam/aws"
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

	resources := []string{
		aws.APIGateway,
		aws.Lambda,
		aws.S3,
		aws.SNS,
		aws.SQS,
		aws.Dynamodb,
	}

	awsSharedEnumerate(resources, saveOutput)
}
