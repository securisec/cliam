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
	"github.com/spf13/cobra"
)

var awsEnumerateCmd = &cobra.Command{
	Use:   "enumerate [resource...]",
	Short: "Enumerate permissions for specified AWS resources.",
	Run:   awsEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return aws.GetAWSResources(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	awsCmd.AddCommand(awsEnumerateCmd)
	awsEnumerateCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsEnumerateCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		printValidArgs(aws.GetAWSResources)
		os.Exit(1)
	}

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	key, secret, token, region := getCredsAndRegion()
	resources := removeDuplicates(args)

	creds := signer.SetCredentials(key, secret, token)

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

				if err := scanner.EnumerateSpecificResource(ctx, region, s, creds, saveOutput); err != nil {
					logger.LoggerStdErr.Err(err).Msg("")
					wg.Done()
					return
				}

				wg.Done()

			}(wg, s)

		}
	}()

	for _, resource := range resources {
		policies, ok := aws.Services[resource]
		if !ok {
			continue
		}
		for _, policy := range policies {
			ch <- scanner.ServiceMap{
				Resource: resource,
				Policy:   policy,
			}
		}
	}

	close(ch)
	wg.Wait()
}
