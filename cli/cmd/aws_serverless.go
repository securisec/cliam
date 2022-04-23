package cmd

import (
	"context"
	"sync"
	"time"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/spf13/cobra"
)

var awsServerlessCmd = &cobra.Command{
	Use:   "serverless",
	Short: "Enumerate permissions for common serverless AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate apigateway, lambda, 
s3, sns, sqs and dynamodb.`,
	Run:               awsServerlessCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsServerlessCmd)
	awsServerlessCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsServerlessCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	key, secret, token, region := getCredsAndRegion()
	cliLogRegion(awsRegion)

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

	resources := []string{
		aws.APIGateway,
		aws.Lambda,
		aws.S3,
		aws.SNS,
		aws.SQS,
		aws.Dynamodb,
	}

	awsSendToChannel(ch, resources, awsKnownResourceName)

	close(ch)
	wg.Wait()
}
