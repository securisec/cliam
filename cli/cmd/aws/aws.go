package awscli

import (
	"fmt"

	"github.com/securisec/iam-enumerate-cli/cmd"
	"github.com/spf13/cobra"
)

var awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Enumerate AWS credentials for their permissions.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("aws called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(awsCmd)
}
