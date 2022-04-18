package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var gcpEnumerateCmd = &cobra.Command{
	Use:   "enumerate [resource...]",
	Short: "Enumerate specified GCP permissions",
	Long: `Enumerate uses cloudresource manager to check if the service account has the specified permissions. 
If resource manager is disabled, this command will end with a 403 error.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printValidArgs(gcp.GetGCPResources)
			os.Exit(1)
		}
	},
	Run: gcpEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp.GetGCPResources(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	gcpCmd.AddCommand(gcpEnumerateCmd)
}

func gcpEnumerateCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}

	services := removeDuplicates(args)

	sa, project, _, _ := getSaAndRegion()

	ctx := context.Background()
	creds, err := scanner.GetCredsFromServiceAccount(ctx, sa)
	if err != nil {
		logger.LoggerStdErr.Err(err).Msg("Failed to get credentials from service account")
		return
	}
	options := scanner.GCPEnumOptions{
		Creds:     creds,
		ProjectId: project,
	}
	ps, err := scanner.EnumerateMultipleResources(ctx, &options, services...)
	if err != nil {
		logger.LoggerStdErr.Err(err).Msg("")
		return
	}
	for _, p := range ps {
		for _, a := range p.Actions {
			a = strings.Split(a, ".")[2]
			logger.LogSuccess(p.Method, a)
		}
	}
}
