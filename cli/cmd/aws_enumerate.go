package cmd

import (
	"context"
	"os"
	"sync"
	"time"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
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

	// spinner things
	// var (
	// 	total, current int
	// )
	// spinner, err := createSpinner()
	// if err != nil {
	// 	logger.Logger.Fatal().Err(err).Msg("failed to create spinner")
	// }
	// defer cleanupSpinner(spinner)

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	key, secret, token, region := getCredsAndRegion()
	cliLogRegion(awsRegion)
	resources := removeDuplicates(args)

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
					// if spinner.Status() == yacspin.SpinnerStopped {
					// 	spinner.Start()
					// }
					// current += 1
					// spinner.Message(fmt.Sprintf("%d/%d", current, total))
					cancel()
					<-max
				}()

				// if spinner.Status() == yacspin.SpinnerRunning {
				// 	spinner.Stop()
				// }

				statusCode, err := scanner.EnumerateSpecificResource(ctx, region, s, creds, saveOutput, awsKnownResourceName)
				if err != nil {
					cliErrorLogger(s, err)
					wg.Done()
					return
				}
				cliCompletionLogger(s, statusCode)

				wg.Done()

			}(wg, s)

		}
	}()

	awsSendToChannel(ch, resources, awsKnownResourceName)

	close(ch)
	wg.Wait()
}
