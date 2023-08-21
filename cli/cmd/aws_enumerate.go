package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
)

var awsEnumerateCmd = &cobra.Command{
	Use:     "enumerate [resource...]",
	Example: "enumerate s3 lambda iam",
	Short:   "Enumerate permissions for specified AWS resources.",
	Run:     awsEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetAWSResources(), cobra.ShellCompDirectiveNoFileComp
	},
	// PreRun:  awsLoadEnvVarsFirst,
	PostRun: PostRunStatsFunc,
}

func init() {
	awsCmd.AddCommand(awsEnumerateCmd)
}

func awsEnumerateCmdFunc(_ *cobra.Command, args []string) {
	if awsEndpoint != "" {
		// TODO: remove
		logger.LogWarning("🚧 Custom endpoint enumeration may not be consistent")
	}

	if len(args) == 0 {
		printValidArgs(aws.GetAWSResources)
		os.Exit(1)
	}

	key, secret, token, region := getCredsAndRegion()
	cliLogRegion(awsRegion)
	resources := shared.RemoveDuplicates(args)

	creds := signer.SetCredentials(key, secret, token, awsProfile)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan scanner.ServiceMap)
	max := make(chan struct{}, MaxThreads)

	go func() {
		defer wg.Done()

		for s := range ch {
			wg.Add(1)

			go func(wg *sync.WaitGroup, s scanner.ServiceMap) {
				max <- struct{}{}
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(RequestTimeout)*time.Second)

				defer func() {
					cancel()
					<-max
				}()

				options := scanner.Options{
					Endpoint:   awsEndpoint,
					Creds:      creds,
					ServiceMap: s,
					Region:     region,
					SaveOutput: SaveOutput,
				}

				statusCode, body, err := scanner.EnumerateSpecificResource(ctx, options)
				if err != nil {
					cliErrorLogger(s, err)
					failureCounter++
					wg.Done()
					return
				}
				cliResponseLoggerAWS(s, statusCode, mapToArray(s.Policy.ExtraValueMap))

				// if deep scanning is enabled, parse response
				if awsDeepScan && s.Policy.ResponseParser != nil && statusCode == 200 {
					extras, err := s.Policy.ResponseParser.ExtraExtractor(body)
					if err != nil {
						wg.Done()
						return
					}
					if len(extras) > 0 {
						// update the extras with extra name
						for _, extra := range extras {
							awsKnownResourceMap = []string{fmt.Sprintf("%s=%s", extra.Flag, extra.ResponseKey)}
							awsSendToChannel(ch, resources, []string{fmt.Sprintf("%s=%s", extra.Flag, extra.ResponseKey)})
						}
					}
				}

				wg.Done()

			}(wg, s)

		}
	}()

	awsSendToChannel(ch, resources, []string{})

	// TODO 🔥 this is blocking if defered, or panicing because closed
	wg.Done() // refacor this because this is poor code and anti pattern
	wg.Wait()

	// TODO 🔥 refactor this and move to aws persistant run. right now it will panic otherwise because of waitgroup
	if saveResults != "" {
		o, _ := json.Marshal(Results)
		if err := os.WriteFile(saveResults, o, os.ModePerm); err != nil {
			logger.LogError(err)
		}
	}
	close(ch)

}

func mapToArray(data map[string]string) []string {
	var hold []string
	for k, v := range data {
		hold = append(hold, fmt.Sprintf("%s=%s", k, v))
	}
	return hold
}
