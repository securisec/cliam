package cmd

import (
	"context"
	"encoding/json"

	"github.com/securisec/cliam/aws/policy"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
)

var awsUtilsStsCallerIdentityCmd = &cobra.Command{
	Use:   "sts-get-caller-identity",
	Short: "Enumerate AWS EC2 snapshots across specified regions",
	Run:   awsUtilsStsCallerIdentityCmdFunc,
	// PreRun: awsLoadEnvVarsFirst,
}

func init() {
	awsUtilsCmd.AddCommand(awsUtilsStsCallerIdentityCmd)
}

func awsUtilsStsCallerIdentityCmdFunc(cmd *cobra.Command, args []string) {
	// get credentials
	key, secret, token, region := getCredsAndRegion()
	creds := signer.SetCredentials(key, secret, token, awsProfile)

	pol := policy.STSPolicies["GetCallerIdentity"]
	url, err := pol.GetRequestURL(region, "sts", awsEndpoint)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get request URL")
	}
	pol.ReqURL = url

	_, res, body, err := signer.MakeScannerRequest(context.Background(), region, "sts", &pol, creds)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to make request")
	}

	if res.StatusCode != 200 {
		logger.LoggerStdErr.Fatal().Msgf("Failed to get caller identity: %d", res.StatusCode)
	}

	var r map[string]interface{}
	json.Unmarshal(body, &r)

	rr, ok := r["GetCallerIdentityResponse"].(map[string]interface{})
	if !ok {
		logger.LoggerStdErr.Fatal().Msg("Failed to parse response")
	}

	l := logger.Logger.Info()
	for k, v := range rr["GetCallerIdentityResult"].(map[string]interface{}) {
		l.Str(k, v.(string))
	}
	l.Msg(shared.GetMessageColor("success"))
}
