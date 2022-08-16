package cmd

import (
	"context"
	"fmt"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/rest"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
	"moul.io/http2curl/v2"
)

var gcpRestCurlBuilderCmd = &cobra.Command{
	Use:     "curl-builder",
	Short:   "Build the curl command to test a gcp policy.",
	Example: "TODO",
	Run:     gcpRestCurlBuilderFunc,
}

func init() {
	gcpRestCmd.AddCommand(gcpRestCurlBuilderCmd)
	gcpRestCurlBuilderCmd.Flags().StringP("policy", "p", "", "The policy to build")
	gcpRestCurlBuilderCmd.Flags().StringP("operation", "o", "", "The operation to build.")
	// gcpRestCurlBuilderCmd.Flags().StringSliceP("values", "n", []string{}, "The values to use for known values.")

	gcpRestCurlBuilderCmd.RegisterFlagCompletionFunc("policy", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return gcp.GetGCPResources(), cobra.ShellCompDirectiveNoFileComp
	})
	gcpRestCurlBuilderCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		hold := []string{}
		pCli, _ := cmd.Flags().GetString("policy")
		if pCli == "" {
			return hold, cobra.ShellCompDirectiveNoSpace
		}
		policies, ok := rest.RestApiCalls[pCli]
		if !ok {
			return hold, cobra.ShellCompDirectiveNoFileComp
		}
		for k := range policies {
			hold = append(hold, k)
		}
		return hold, cobra.ShellCompDirectiveNoFileComp
	})
	// gcpRestCurlBuilderCmd.RegisterFlagCompletionFunc("values", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// 	hold := []string{}
	// 	pCli, _ := cmd.Flags().GetString("policy")
	// 	oCli, _ := cmd.Flags().GetString("operation")

	// 	p := gcpGetSpecificOperation(pCli, oCli)
	// 	matches := shared.AzureTemplatePropertyRegex.FindAllStringSubmatch(p.URL, -1)
	// 	for _, m := range matches {
	// 		hold = append(hold, m[1]+"=")
	// 	}
	// 	return hold, cobra.ShellCompDirectiveNoSpace
	// })

	gcpRestCurlBuilderCmd.MarkFlagRequired("policy")
	gcpRestCurlBuilderCmd.MarkFlagRequired("operation")
}

func gcpRestCurlBuilderFunc(cmd *cobra.Command, _ []string) {
	var err error
	ctx := context.Background()
	sa, _, region, zone := getSaAndRegion()

	pCli, _ := cmd.Flags().GetString("policy")
	oCli, _ := cmd.Flags().GetString("operation")
	// values, _ := cmd.Flags().GetStringSlice("values")
	p := gcpGetSpecificOperation(pCli, oCli)
	parentType, parentID := processParent(gcpRestParent)

	if parentType == "" || parentID == "" {
		if gcpProjectId != "" {
			parentType = "projects"
			parentID = gcpProjectId
		} else {
			logger.LoggerStdErr.Fatal().Msg("Parent type and id must be set")
		}
	}

	accessToken := gcpAccessToken
	if accessToken == "" {
		accessToken, err = gcp.GetAccessToken(ctx, sa)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get access token")
		}
	}

	p.ParentType = parentType
	p.ParentID = parentID
	p.ResourceID = gcpRestResourceID
	p.Zone = zone
	p.Region = region
	p.ReqBody = gcpRestBody
	req, err := scanner.RequestBuilder(ctx, p, accessToken)
	if err != nil {
		logger.LoggerStdErr.Fatal().Str("url", p.URL).Err(err).Msg("Failed to build request")
	}

	c, err := http2curl.GetCurlCommand(req)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get curl command")
	}
	fmt.Println(c.String())
}

func gcpGetSpecificOperation(policy, operationId string) rest.RestCall {
	policies, ok := rest.RestApiCalls[policy]
	if !ok {
		logger.LoggerStdErr.Fatal().Msgf("Policy %s not found", policy)
	}
	o, ok := policies[operationId]
	if !ok {
		logger.LoggerStdErr.Fatal().Msgf("Operation %s not found", operationId)
	}

	return o
}
