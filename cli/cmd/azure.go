package cmd

import (
	"fmt"
	"os"

	"github.com/securisec/cliam/azure"
	"github.com/securisec/cliam/azure/policy"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
)

var (
	azureSubscriptionID    string
	azureTenantID          string
	azureClientID          string
	azureClientSecret      string
	azureResourceGroupName string
	azureOauthToken        string
	azureCertificatePath   string
	azureKnownResourceMap  map[string]string
	// TODO support other auth types
)

var azureCmd = &cobra.Command{
	Use:   "azure",
	Short: "Enumerate Azure credentials for their permissions.",
	Long: `Enumerate Azure credentials for their permissions.
In most cases, a valid Azure Subscription ID is required. If one 
is not provided, the CLI will attempt to lookup available subscriptions 
and use the first one.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.LoggerStdErr.Error().Msg("azure not implemented yet")
	},
}

func init() {
	RootCmd.AddCommand(azureCmd)

	azureCmd.PersistentFlags().StringVar(&azureSubscriptionID, "subscription-id", "", "Azure Subscription ID")
	azureCmd.PersistentFlags().StringVar(&azureTenantID, "tenant-id", "", "Azure Tenant ID")
	azureCmd.PersistentFlags().StringVar(&azureClientID, "client-id", "", "Azure Client ID / Username")
	azureCmd.PersistentFlags().StringVar(&azureClientSecret, "client-secret", "", "Azure Client Secret / Password")
	azureCmd.PersistentFlags().StringVar(&azureResourceGroupName, "resource-group-name", "", "Azure Resource Group")
	azureCmd.PersistentFlags().StringVar(&azureOauthToken, "oauth-token", "", "Optionall use a valid Azure OAuth Token. Can also use CLIAM_AZURE_OAUTH_TOKEN envvar")
	azureCmd.PersistentFlags().StringVar(&azureCertificatePath, "certificate-path", "", "Path to Certificate for certificate based authentication")
	azureCmd.PersistentFlags().StringToStringVar(&azureKnownResourceMap, "known-value", map[string]string{}, "Azure cli flags. When known-resource-name is set, additional permissions where a resource needs to be specified is enumerated.")
}

// 🚧 TODO: Implement envar support for Azure. flags override envars.
func azureValidateRequiredFlags(cmd *cobra.Command, args []string) {
	// https://github.com/Azure/azure-sdk/blob/main/_includes/tables/environment_variables.md
	if azureClientID == "" {
		azureClientID = os.Getenv("AZURE_CLIENT_ID")
	}
	if azureClientSecret == "" {
		azureClientSecret = os.Getenv("AZURE_CLIENT_SECRET")
	}
	if azureTenantID == "" {
		azureTenantID = os.Getenv("AZURE_TENANT_ID")
	}

	if azureSubscriptionID == "" {
		azureSubscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	}
	if azureResourceGroupName == "" {
		azureResourceGroupName = os.Getenv("AZURE_RESOURCE_GROUP")
	}
	if azureOauthToken == "" {
		azureOauthToken = os.Getenv("CLIAM_AZURE_OAUTH_TOKEN")
	}
	if azureCertificatePath == "" {
		azureCertificatePath = os.Getenv("AZURE_CLIENT_CERTIFICATE_PATH")
	}

	if azureCertificatePath == "" && azureClientSecret == "" {
		logger.LoggerStdErr.Fatal().Msg("azure requires a client secret or certificate path")
	}

	for _, v := range []string{
		// azureSubscriptionID,
		azureTenantID,
		azureClientID,
	} {
		if v == "" {
			logger.LoggerStdErr.Fatal().Msg(fmt.Sprintf("azure required flags are missing: %s", v))
		}
	}
}

func azureSendToChannel(ch chan policy.Policy, resources []string) {
	var extras []policy.Policy
	hold := make([]policy.Policy, 0)

	rm := shared.RemoveDuplicates(resources)
	for _, resource := range rm {
		policies, ok := azure.Policies[resource]
		if !ok {
			continue
		}
		for _, policy := range policies {
			if policy.IsExtra {
				extras = append(extras, policy)
			} else {
				hold = append(hold, policy)
			}
		}
	}

	for _, p := range hold {
		ch <- p
	}

	// if a known resource name is set, we will enumerate only the extra permissions
	if len(azureKnownResourceMap) > 0 && len(extras) > 0 {
		for _, ee := range extras {
			ee.ExtraValue = azureKnownResourceMap
			// ee.ExtraValue = awsModifyExtraMap(awsKnownResourceMap)
			ch <- ee
		}
	}
}

func azureLogSuccessMessage(policy policy.Policy, body map[string]interface{}, extras ...map[string]string) {
	l := logger.Logger.Info().Str(policy.Resource, policy.OperationID)
	if len(extras) > 0 && logger.VERBOSE {
		for k, v := range extras[0] {
			l = l.Str(k, v)
		}
	}
	if _, hasValue := body["value"]; hasValue {
		v, ok := body["value"].([]interface{})
		if ok {
			if len(v) == 0 && logger.VERBOSE {
				logger.LoggerStdErr.Debug().Str(policy.Resource, policy.OperationID).Interface("body", body).Msg(shared.GetMessageColor("maybe"))
				return
			}
		}
	}
	l.Msg(shared.GetMessageColor("success"))
}

func azureGetOauthToken() string {
	if azureOauthToken != "" {
		return azureOauthToken
	}
	// client secret is provided so we will use that
	if azureClientSecret != "" {
		token, err := azure.GetTokenFromUsernameAndPassword(azureTenantID, azureClientID, azureClientSecret)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("failed to get azure oauth token")
		}
		return token
		// use certificate to obtain token
	} else if azureCertificatePath != "" {
		token, err := azure.GetTokenFromCertificate(azureTenantID, azureClientID, azureCertificatePath)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("failed to get azure oauth token")
		}
		return token
	}
	logger.Logger.Fatal().Msg("azure oauth token or client secret is required")
	return ""
}
