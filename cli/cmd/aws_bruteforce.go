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
	awsBruteforceCmd.Flags().Int("max-threads", 5, "Maximum number of threads to use")
	// awsBruteforceCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsBruteforceCmdFunc(cmd *cobra.Command, args []string) {
	// saveOutput, _ := cmd.Flags().GetBool("save-output")
	maxThreads, _ := cmd.Flags().GetInt("max-threads")
	ch := make(chan string)

	key, secret, token, region := getCredsAndRegion()
	creds := signer.SetCredentials(key, secret, token)

	wg := &sync.WaitGroup{}
	wg.Add(maxThreads)

	for i := 0; i < maxThreads; i++ {

		go func() {
			defer wg.Done()
			for {
				s, ok := <-ch
				if !ok {
					return
				}

				if err := scanner.EnumerateSpecificResource(context.Background(), region, s, creds); err != nil {
					logger.LoggerStdErr.Err(err).Msg("")
					return
				}

			}
		}()

	}

	for _, service := range aws.GetAWSResources() {
		ch <- service
	}

	close(ch)
	wg.Wait()
}
