package awscli

import (
	"fmt"

	"github.com/securisec/iam-enumerate/aws/policy"
	"github.com/spf13/cobra"
)

var awsSingleCmd = &cobra.Command{
	Use:   "service",
	Short: "Enumerate permissions for a single AWS service.",
	Run:   awsSingleCmdFunc,
	// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 	return nil, cobra.ShellCompDirectiveNoFileComp
	// },
}

func init() {
	awsCmd.AddCommand(awsSingleCmd)
}

func awsSingleCmdFunc(cmd *cobra.Command, args []string) {
	// TODO: implement
	policy.Services
	fmt.Println("aws called")
}
