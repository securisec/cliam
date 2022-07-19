package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/securisec/cliam/azure"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var azureSubscriptionCmd = &cobra.Command{
	Use:    "subscriptions",
	Short:  "List all subscriptions",
	PreRun: azureValidateRequiredFlags,
	Run:    azureSubscriptionCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return azure.GetPolicyKeys(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	azureCmd.AddCommand(azureSubscriptionCmd)
}

func azureSubscriptionCmdFunc(_ *cobra.Command, _ []string) {
	token := azureGetOauthToken()
	_, subs, err := azure.GetFirstSubscriptions(token)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get subscription ID")
	}
	o, err := json.MarshalIndent(subs, "", "  ")
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to marshal subscriptions")
	}

	fmt.Println(string(o))
}
