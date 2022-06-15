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
	successCounter = 0
	failureCounter = 0
	maybeCounter   = 0
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
}

func init() {
	RootCmd.AddCommand(azureCmd)

	azureCmd.PersistentFlags().StringVar(&azureSubscriptionID, "subscription-id", "", "Azure Subscription ID")
	azureCmd.PersistentFlags().StringVarP(&azureTenantID, "tenant-id", "t", "", "Azure Tenant ID")
	azureCmd.PersistentFlags().StringVar(&azureClientID, "client-id", "", "Azure Client ID / Username")
	azureCmd.PersistentFlags().StringVar(&azureClientSecret, "client-secret", "", "Azure Client Secret / Password")
	azureCmd.PersistentFlags().StringVar(&azureResourceGroupName, "resource-group-name", "", "Azure Resource Group")
	azureCmd.PersistentFlags().StringVar(&azureOauthToken, "oauth-token", "", "Optionall use a valid Azure OAuth Token. Can also use CLIAM_AZURE_OAUTH_TOKEN envvar")
	azureCmd.PersistentFlags().StringVar(&azureCertificatePath, "certificate-path", "", "Path to Certificate for certificate based authentication")
	azureCmd.PersistentFlags().StringToStringVar(&azureKnownResourceMap, "known-value", map[string]string{}, "Azure cli flags. When known-resource-name is set, additional permissions where a resource needs to be specified is enumerated.")

	// know value completer
	azureCmd.RegisterFlagCompletionFunc("known-value", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return azureKnownValueCompleter(args), cobra.ShellCompDirectiveNoSpace
	})
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
	hold := make([]policy.Policy, 0)

	rm := shared.RemoveDuplicates(resources)
	for _, resource := range rm {
		policies, ok := azure.Policies[resource]
		if !ok {
			continue
		}
		for _, policy := range policies {
			policy.PathValues = azureModifyExtraMap(azureKnownResourceMap)
			hold = append(hold, policy)
		}
	}

	for _, p := range hold {
		ch <- p
	}
}

func azureLogSuccessMessage(policy policy.Policy, body map[string]interface{}, extras ...map[string]string) {
	l := logger.Logger.Info().Str(policy.Resource, policy.OperationID)
	if len(extras) > 0 && CLIVerbose {
		for k, v := range extras[0] {
			l = l.Str(k, v)
		}
	}
	if _, hasValue := body["value"]; hasValue {
		v, ok := body["value"].([]interface{})
		if ok {
			if len(v) == 0 && CLIVerbose {
				maybeCounter += 1
				logger.LoggerStdErr.Debug().Str(policy.Resource, policy.OperationID).Interface("body", body).Msg(shared.GetMessageColor("maybe"))
				return
			}
		}
	} else {
		successCounter += 1
		l.Msg(shared.GetMessageColor("success"))
	}
}

func azureGetOauthToken() string {
	if azureOauthToken != "" {
		if !ValidateJwtExpiration(azureOauthToken) {
			logger.LoggerStdErr.Fatal().Msg("azure oauth token is expired")
		}
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

func azureModifyExtraMap(m map[string]string) map[string]string {
	h := map[string]string{}
	for k, v := range m {
		h[shared.KebabToCamelCase(k)] = v
	}
	return h
}

func azureKnownValueCompleter(resources []string) (validExtras []string) {
	hold := []string{}
	for _, resource := range shared.RemoveDuplicates(resources) {
		policies, ok := azure.Policies[resource]
		if !ok {
			continue
		}
		for _, policy := range policies {
			matches := shared.TemplatePropertyRegex.FindAllStringSubmatch(policy.Path, -1)
			for _, m := range matches {
				hold = append(hold, m[1]+"=")
			}
		}
	}
	return shared.RemoveDuplicates(hold)
}

func azurePostRunFunc(cmd *cobra.Command, args []string) {
	if CLIVerbose {
		logger.LoggerStdErr.Debug().Msgf(
			"azure: %s, %s, %s",
			shared.Green(fmt.Sprintf("%d success", successCounter)),
			shared.Yellow(fmt.Sprintf("%d maybe", maybeCounter)),
			shared.Red(fmt.Sprintf("%d failure", failureCounter)),
		)
	}
}
