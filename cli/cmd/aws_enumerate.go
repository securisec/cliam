package cmd

import (
	"context"
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

	ch := make(chan scanner.ServiceMap, 0)
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

				statusCode, err := scanner.EnumerateSpecificResource(ctx, region, awsEndpoint, s, creds, SaveOutput)
				if err != nil {
					cliErrorLogger(s, err)
					failureCounter++
					wg.Done()
					return
				}
				cliResponseLoggerAWS(s, statusCode)

				wg.Done()

			}(wg, s)

		}
	}()

	awsSendToChannel(ch, resources)

	close(ch)
	wg.Wait()
}
