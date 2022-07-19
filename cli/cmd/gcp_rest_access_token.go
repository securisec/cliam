package cmd

import (
	"context"
	"fmt"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var gcpRestAccessTokenCmd = &cobra.Command{
	Use:   "access-token",
	Short: "Get access token for the specified service account",
	Long: `This command outputs the access token for the specified service account which can then 
be used to manually GCP REST API's.`,
	Run:               gcpRestAccessTokenCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	gcpCmd.AddCommand(gcpRestAccessTokenCmd)
}

func gcpRestAccessTokenCmdFunc(_ *cobra.Command, _ []string) {

	sa, _, _, _ := getSaAndRegion()

	ctx := context.Background()
	accessToken, err := gcp.GetAccessToken(ctx, sa)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get access token")
	}
	fmt.Println(accessToken)
}
