package cmd

import (
	"github.com/securisec/cliam/aws"
	"github.com/spf13/cobra"
)

var awsComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Enumerate permissions for common compute AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate efs, ec2, 
elasticbeanstalk, elb and lambda.`,
	Run:               awsComputeCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsComputeCmd)
	awsComputeCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsComputeCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	resources := []string{
		aws.Lambda,
		aws.EC2,
		aws.ElasticBeanStalk,
		aws.ElasticFileSystem,
		aws.ELB,
	}

	awsSharedEnumerate(resources, saveOutput)
}
