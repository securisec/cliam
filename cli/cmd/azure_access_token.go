package cmd

import (
	"fmt"

	"github.com/securisec/cliam/azure"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var azureAccessTokenCmd = &cobra.Command{
	Use:               "access-token",
	Short:             "Get access token to use with REST apis",
	PreRun:            azureValidateRequiredFlags,
	Run:               azureAccessTokenCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	azureCmd.AddCommand(azureAccessTokenCmd)
}

func azureAccessTokenCmdFunc(_ *cobra.Command, _ []string) {

	token, err := azure.GetTokenFromUsernameAndPassword(azureTenantID, azureClientID, azureClientSecret)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get access token")
	}
	fmt.Println(token)
}
