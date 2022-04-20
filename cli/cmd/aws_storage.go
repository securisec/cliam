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

var awsStorageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Enumerate permissions for common storage AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate s3, efs, 
snowball and storage gateway.`,
	Run:               awsStorageCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsStorageCmd)
	awsStorageCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsStorageCmdFunc(cmd *cobra.Command, _ []string) {

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	key, secret, token, region := getCredsAndRegion()

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

				statusCode, err := scanner.EnumerateSpecificResource(ctx, region, s, creds, saveOutput)
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
		aws.S3,
		aws.Snowball,
		aws.StorageGateway,
		aws.ElasticFileSystem,
	}

	enumerate := scanner.GetServiceMap(resources)
	for _, e := range enumerate {
		ch <- e
	}

	close(ch)
	wg.Wait()
}
