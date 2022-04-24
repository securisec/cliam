package cmd

import (
	"github.com/securisec/cliam/aws"
	"github.com/spf13/cobra"
)

var awsDatabasesCmd = &cobra.Command{
	Use:   "databases",
	Short: "Enumerate permissions for common AWS database resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate redshift, rds, 
and dynamodb.`,
	Run:               awsDatabasesCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsDatabasesCmd)
	awsDatabasesCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsDatabasesCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	resources := []string{
		aws.RDS,
		aws.Redshift,
		aws.Dynamodb,
	}

	awsSharedEnumerate(resources, saveOutput)
}
