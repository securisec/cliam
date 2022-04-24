package cmd

import (
	"github.com/securisec/cliam/aws"
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

	resources := []string{
		aws.Lambda,
		aws.IAM,
		aws.S3,
		aws.APIGateway,
		aws.EC2,
	}

	awsSharedEnumerate(resources, saveOutput)
}
