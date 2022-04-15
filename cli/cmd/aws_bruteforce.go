package cmd

import (
	"context"
	"sync"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsBruteforceCmd = &cobra.Command{
	Use:               "bruteforce",
	Short:             "Bruteforce permissions for all AWS resources.",
	Run:               awsBruteforceCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsBruteforceCmd)
	// awsBruteforceCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsBruteforceCmdFunc(cmd *cobra.Command, args []string) {
	// saveOutput, _ := cmd.Flags().GetBool("save-output")
	ch := make(chan string)

	key, secret, token, region := getCredsAndRegion()
	creds := signer.SetCredentials(key, secret, token)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	max := make(chan struct{}, MaxThreads)

	go func() {
		defer wg.Done()
		for s := range ch {

			wg.Add(1)

			go func(wg *sync.WaitGroup, ser string) {
				max <- struct{}{}
				defer func() {
					<-max
				}()

				if err := scanner.EnumerateSpecificResource(context.Background(), region, ser, creds); err != nil {
					logger.LoggerStdErr.Err(err).Msg("")
					wg.Done()
					return
				}

				wg.Done()
			}(wg, s)

		}
	}()

	for _, service := range aws.GetAWSResources() {
		ch <- service
	}

	close(ch)
	wg.Wait()
}
