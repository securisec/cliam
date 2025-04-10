package cmd

import (
	"fmt"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/policy"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
	"moul.io/http2curl/v2"
)

var awsRequestBuilderCmd = &cobra.Command{
	Use:     "curl-builder",
	Short:   "Build the curl command to test an aws policy. Only the first specified AWS region is used",
	Example: "aws curl-builder --policy somepolicy --operation someoperation --values somevalue=somevalue",
	Long:    "Some requests requires known values. For these, use the -n command to supply them",
	PreRun: func(_ *cobra.Command, _ []string) {
		if len(awsRegions) == 0 {
			logger.LoggerStdErr.Fatal().Msg("At least one region is required")
		}
	},
	Run:               awsRequestBuilderCmdFunc,
	Args:              cobra.NoArgs,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsRequestBuilderCmd)
	awsRequestBuilderCmd.Flags().StringP("resource", "p", "", "The resource to build")
	awsRequestBuilderCmd.Flags().StringP("operation", "o", "", "The operation to build.")
	awsRequestBuilderCmd.Flags().StringSliceP("values", "n", []string{}, "The values to use for known values.")

	// completers
	awsRequestBuilderCmd.RegisterFlagCompletionFunc("resource", func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return aws.GetAWSResources(), cobra.ShellCompDirectiveNoFileComp
	})
	awsRequestBuilderCmd.RegisterFlagCompletionFunc("operation", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		hold := []string{}
		resource, _ := cmd.Flags().GetString("resource")
		if resource == "" {
			return hold, cobra.ShellCompDirectiveNoSpace
		}
		policies, ok := aws.Services[resource]
		if !ok {
			return hold, cobra.ShellCompDirectiveNoFileComp
		}
		for k := range policies {
			hold = append(hold, k)
		}
		return hold, cobra.ShellCompDirectiveNoSpace
	})

	// know value completer
	awsRequestBuilderCmd.RegisterFlagCompletionFunc("values", func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		p, _ := cmd.Flags().GetString("resource")
		o, _ := cmd.Flags().GetString("operation")
		pol := awsGetSpecificOperation(p, o)
		hold := []string{}

		if pol.ExtraCommandLineFlag != "" {
			hold = append(hold, pol.ExtraCommandLineFlag)
		}
		return hold, cobra.ShellCompDirectiveNoSpace
	})

	// mark required
	awsRequestBuilderCmd.MarkFlagRequired("resource")
	awsRequestBuilderCmd.MarkFlagRequired("operation")
}

func awsRequestBuilderCmdFunc(cmd *cobra.Command, _ []string) {
	pCli, _ := cmd.Flags().GetString("resource")
	operation, _ := cmd.Flags().GetString("operation")
	values, _ := cmd.Flags().GetStringSlice("values")

	key, secret, token, region := getCredsAndRegion()

	creds := signer.SetCredentials(key, secret, token, awsProfile)

	p := awsGetSpecificOperation(pCli, operation)

	extraMap := ModifyExtraMap(values)
	if len(extraMap) > 0 {
		p.ExtraValueMap = extraMap
	}

	u, err := p.GetRequestURL(region[0], pCli, awsEndpoint)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("error getting request url")
	}
	p.ReqURL = u
	if p.IsExtra {
		s, err := p.UpdateForExtra()
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("error updating policy for extra")
		}
		p = s
	}

	req, err := signer.BuildRequest(region[0], pCli, &p, creds)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("error building request")
	}

	c, err := http2curl.GetCurlCommand(req)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("error getting curl command")
	}
	fmt.Println(c.String())
}

func awsGetSpecificOperation(policy, operationID string) policy.Service {
	p, ok := aws.Services[policy]
	if !ok {
		logger.LoggerStdErr.Fatal().Msg("policy not found")
	}
	o, ok := p[operationID]
	if !ok {
		logger.LoggerStdErr.Fatal().Msg("operation not found")
	}
	return o
}
