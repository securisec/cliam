package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/securisec/cliam/azure"
	"github.com/securisec/cliam/azure/policy"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
	"moul.io/http2curl/v2"
)

var azureRequestBuilderCmd = &cobra.Command{
	Use:   "curl-builder",
	Short: "Build the curl command to test an azure policy.",
	Long:  "Some requests requires known values. For these, use the -k command to supply them",
	PreRun: func(cmd *cobra.Command, args []string) {
		if azureSubscriptionID == "" {
			azureSubscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
		}
		if azureResourceGroupName == "" {
			azureResourceGroupName = os.Getenv("AZURE_RESOURCE_GROUP")
		}
		if azureOauthToken == "" {
			azureOauthToken = os.Getenv("CLIAM_AZURE_OAUTH_TOKEN")
		}
	},
	Run:               azureRequestBuilderCmdFunc,
	Args:              cobra.NoArgs,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	azureCmd.AddCommand(azureRequestBuilderCmd)
	azureRequestBuilderCmd.Flags().StringP("policy", "p", "", "The policy to build")
	azureRequestBuilderCmd.Flags().StringP("operation", "o", "", "The operation to build.")
	azureRequestBuilderCmd.Flags().StringSliceP("values", "n", []string{}, "The values to use for known values.")

	// completers
	azureRequestBuilderCmd.RegisterFlagCompletionFunc("policy", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return azure.GetPolicyKeys(), cobra.ShellCompDirectiveNoFileComp
	})
	azureRequestBuilderCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		hold := []string{}
		policy, _ := cmd.Flags().GetString("policy")
		if policy == "" {
			return hold, cobra.ShellCompDirectiveNoSpace
		}
		policies, ok := azure.Policies[policy]
		if !ok {
			return hold, cobra.ShellCompDirectiveNoFileComp
		}
		for k := range policies {
			hold = append(hold, k)
		}
		return hold, cobra.ShellCompDirectiveNoFileComp
	})

	// know value completer
	azureRequestBuilderCmd.RegisterFlagCompletionFunc("values", func(cmd *cobra.Command, _ []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		p, _ := cmd.Flags().GetString("policy")
		o, _ := cmd.Flags().GetString("operation")
		pol := azureGetSpecificOperation(p, o)
		hold := []string{}
		matches := shared.TemplatePropertyRegex.FindAllStringSubmatch(pol.Path, -1)
		for _, m := range matches {
			hold = append(hold, m[1]+"=")
		}
		return hold, cobra.ShellCompDirectiveNoSpace
	})

	// mark required
	azureRequestBuilderCmd.MarkFlagRequired("policy")
	azureRequestBuilderCmd.MarkFlagRequired("operation")
}

func azureRequestBuilderCmdFunc(cmd *cobra.Command, _ []string) {
	pCli, _ := cmd.Flags().GetString("policy")
	operation, _ := cmd.Flags().GetString("operation")
	values, _ := cmd.Flags().GetStringSlice("values")

	p := azureGetSpecificOperation(pCli, operation)
	p.PathValues = ModifyExtraMap(values)

	if azureOauthToken == "" {
		azureOauthToken = "CHANGE_ME"
	}

	req, err := p.BuildRequest(context.Background(), azureOauthToken, policy.CommonPathProperties{
		SubscriptionID:    azureSubscriptionID,
		ResourceGroupName: azureResourceGroupName,
	})
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("failed to build request")
	}

	curl, err := http2curl.GetCurlCommand(req)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("failed to get curl command")
	}
	fmt.Println(curl.String())
}

func azureGetSpecificOperation(policy, operationId string) policy.Policy {
	p, ok := azure.Policies[policy]
	if !ok {
		logger.LoggerStdErr.Fatal().Msg("policy not found")
	}
	o, ok := p[operationId]
	if !ok {
		logger.LoggerStdErr.Fatal().Msg("operation not found")
	}
	return o
}
