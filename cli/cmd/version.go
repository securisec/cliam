package cmd

import (
	"fmt"

	"github.com/securisec/cliam/cli/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version and build info",
	Run: func(cmd *cobra.Command, args []string) {
		v := fmt.Sprintf(
			"Version: %s\nBuildDate: %s\nGitCommit: %s\nGitBranch: %s",
			version.Version,
			version.BuildDate,
			version.GitCommit,
			version.GitBranch,
		)
		fmt.Println(v)
	},
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
