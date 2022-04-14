package cmd

import (
	"context"
	"strings"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var gcpSingleCmd = &cobra.Command{
	Use:   "enumerate [resource...]",
	Short: "Enumerate specified GCP permissions",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	Run: gcpServiceCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp.GetGCPResources(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	gcpCmd.AddCommand(gcpSingleCmd)
}

func gcpServiceCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}

	ctx := context.Background()
	creds, err := scanner.GetCredsFromServiceAccount(ctx, getSaPath())
	if err != nil {
		panic(err)
	}
	ps, err := scanner.EnumerateMultipleResources(ctx, creds, getProjectId(), args...)
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		for _, a := range p.Actions {
			a = strings.Split(a, ".")[2]
			logger.LogSuccess(p.Method, a)
		}
	}
}
