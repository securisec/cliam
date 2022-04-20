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

var awsComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Enumerate permissions for common compute AWS resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate efs, ec2, 
elasticbeanstalk, elb and lambda.`,
	Run:               awsComputeCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsComputeCmd)
	awsComputeCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsComputeCmdFunc(cmd *cobra.Command, _ []string) {

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
		aws.Lambda,
		aws.EC2,
		aws.ElasticBeanStalk,
		aws.ElasticFileSystem,
		aws.ELB,
	}

	enumerate := scanner.GetServiceMap(resources)
	for _, e := range enumerate {
		ch <- e
	}

	close(ch)
	wg.Wait()
}
