package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/rest"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
)

var gcpRestEnumerateCmd = &cobra.Command{
	Use:     "enumerate [resource...]",
	Short:   "Enumerate specified GCP permissions using the REST API",
	Long:    `Unlike enumerate, this command uses the REST API to enumerate permissions.`,
	Example: "gcp rest enumerate --parent project=my-project compute.instances",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printValidArgs(gcp.GetAvailableRestKeys)
			os.Exit(1)
		}
	},
	Run: gcpRestEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp.GetAvailableRestKeys(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	gcpRestCmd.AddCommand(gcpRestEnumerateCmd)
	gcpRestEnumerateCmd.Flags().StringToString("parent", nil, "Specify the parent. i.e. project=my-project or organization=my-org. Valid keys are project, origanization, folder, billingAccount")
	gcpRestEnumerateCmd.Flags().StringToString("body", map[string]string{}, "Rest API body to use when not a GET request")
	gcpRestEnumerateCmd.Flags().String("resource-id", "", `Resource ID to use when not a GET request. Resource id is quite 
versitile and is used as a generic term for any resource. For example, if you are
enumerating pubsub, resource-id could be the subscription name; but also if you are 
enumerating compute, resource-id could be the instance name.`)
	// gcpRestEnumerateCmd.MarkFlagRequired("parent")
	gcpRestEnumerateCmd.Flags().Bool("save-output", false, "Save output to file")
}

func gcpRestEnumerateCmdFunc(cmd *cobra.Command, args []string) {
	accessToken := gcpAccessToken

	parent, err := cmd.Flags().GetStringToString("parent")
	body, err := cmd.Flags().GetStringToString("body")
	resourceID, err := cmd.Flags().GetString("resource-id")
	saveOutput, err := cmd.Flags().GetBool("save-output")

	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Invalid parent")
		os.Exit(1)
	}

	sa, projectID, region, zone := getSaAndRegion()
	cliGcpLogRegion(region, zone)
	if projectID != "" {
		parent["project"] = projectID
	}

	parentType, parentID := processParent(parent)

	if parentType == "" || parentID == "" {
		logger.LoggerStdErr.Fatal().Msg("parent must be specified as project=my-project or organization=my-org")
	}

	resources := shared.RemoveDuplicates(args)

	ctx := context.Background()
	if gcpAccessToken == "" {
		accessToken, err = gcp.GetAccessToken(ctx, sa)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get access token")
		}
	}

	// get array of permissions to check
	permissions := make([]rest.RestCall, 0)
	for _, resource := range resources {
		if r, ok := rest.RestApiCalls[resource]; ok {
			for _, p := range r {
				p.ParentType = parentType
				p.ParentID = parentID
				p.ResourceID = resourceID
				p.Zone = zone
				p.Region = region
				p.ReqBody = body
				permissions = append(permissions, p)
			}
		}
	}

	if len(permissions) == 0 {
		logger.LoggerStdErr.Fatal().Msg("No permissions found")
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan rest.RestCall, 0)
	// the is the maximum concurrent goroutines
	max := make(chan struct{}, MaxThreads)

	go func() {
		defer wg.Done()

		for s := range ch {
			wg.Add(1)

			go func(wg *sync.WaitGroup, ser rest.RestCall) {

				max <- struct{}{}
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(RequestTimeout)*time.Second)
				defer func() {
					cancel()
					<-max
				}()

				res, body, err := scanner.EnumerateRestApiRequest(ctx, accessToken, ser)
				if err != nil {
					logger.LoggerStdErr.Debug().Str(ser.PermissionMethod, ser.Action).Err(err).Msg("Failed to enumerate")
					wg.Done()
					return
				}
				if res.Response.StatusCode == 200 {
					logger.LogSuccess(res.PermissionMethod, res.Action)
					// optionall save output
					if saveOutput {
						path := fmt.Sprintf("%s.%s.json", ser.PermissionMethod, ser.Action)
						if err := ioutil.WriteFile(path, body, os.ModePerm); err != nil {
							logger.LoggerStdErr.Debug().Err(err).Msg("Failed to save output")
						}
					}
				}

				if logger.DEBUG && res.Response.StatusCode != 200 {
					logger.LogDenied(res.Response.StatusCode, res.PermissionMethod, res.Action)
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
