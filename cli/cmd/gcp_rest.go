package cmd

import (
	"context"
	"os"
	"sync"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/policy"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var gcpRestCmd = &cobra.Command{
	Use:     "rest [resource...]",
	Short:   "Enumerate specified GCP permissions using the REST API",
	Long:    `Unlike enumerate, this command uses the REST API to enumerate permissions.`,
	Example: "gcp rest --parent project=my-project compute.instances",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
	},
	Run: gcpRestCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp.GetGCPResources(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	gcpCmd.AddCommand(gcpRestCmd)
	gcpRestCmd.Flags().StringToString("parent", nil, "Specify the parent. i.e. project=my-project or organization=my-org")
	gcpRestCmd.MarkFlagRequired("parent")
}

func gcpRestCmdFunc(cmd *cobra.Command, args []string) {
	parent, _ := cmd.Flags().GetStringToString("parent")

	parentType, parentID := processParent(parent)
	if parentType == "" || parentID == "" {
		logger.LoggerStdErr.Fatal().Msg("parent must be specified as project=my-project or organization=my-org")
	}

	resources := removeDuplicates(args)

	sa, _, region := getSaAndRegion()

	ctx := context.Background()
	accessToken, err := gcp.GetAccessToken(ctx, sa)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get access token")
	}

	// get array of permissions to check
	permissions := make([]policy.RestCall, 0)
	for _, resource := range resources {
		if r, ok := policy.Resources[resource]; ok {
			for _, p := range r.RESTCalls {
				p.ParentType = parentType
				p.ParentID = parentID
				p.Zone = region
				// pol, err := p.GetURL()
				// if err != nil {
				// 	logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get policy URL")
				// }
				permissions = append(permissions, p)
			}
		}
	}

	if len(permissions) == 0 {
		logger.LoggerStdErr.Fatal().Msg("No permissions found")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan policy.RestCall, 0)
	// the is the maximum concurrent goroutines
	max := make(chan struct{}, MaxThreads)

	go func() {
		defer wg.Done()

		for s := range ch {
			wg.Add(1)

			go func(wg *sync.WaitGroup, ser policy.RestCall) {

				max <- struct{}{}
				defer func() {
					<-max
				}()

				res, err := scanner.EnumerateRestApiRequest(ctx, accessToken, ser)
				if err != nil && logger.DEBUG {
					logger.LoggerStdErr.Debug().Err(err).Msg("Failed to enumerate")
				}
				if res.Response.StatusCode == 200 {
					logger.LogSuccess(res.PermissionMethod, res.Action)
				}

				wg.Done()

			}(wg, s)

		}
	}()

	for _, v := range permissions {
		ch <- v
	}

	close(ch)
	wg.Wait()
}

func processParent(parent map[string]string) (string, string) {
	if parent == nil {
		logger.LoggerStdErr.Error().Msg("Parent is required")
		os.Exit(1)
	}
	if v, ok := parent["project"]; ok {
		return "projects", v
	}
	if v, ok := parent["organization"]; ok {
		return "organizations", v
	}
	if v, ok := parent["folder"]; ok {
		return "folders", v
	}
	if v, ok := parent["billingAccount"]; ok {
		return "billingAccounts", v
	}
	logger.LoggerStdErr.Error().Msg("Invalid parent")
	return "", ""
}
