package cmd

import (
	"github.com/securisec/cliam/aws"
	"github.com/spf13/cobra"
)

var awsStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Enumerate permissions for common storage AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate s3, efs, 
snowball and storage gateway.`,
	Run:               awsStorageCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsStorageCmd)
	awsStorageCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsStorageCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	resources := []string{
		aws.S3,
		aws.Snowball,
		aws.StorageGateway,
		aws.ElasticFileSystem,
	}

	awsSharedEnumerate(resources, saveOutput)
}
