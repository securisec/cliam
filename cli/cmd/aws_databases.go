package cmd

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var awsDatabasesCmd = &cobra.Command{
	Use:   "databases",
	Short: "Enumerate permissions for common AWS database resources.",
	Long: `Enumerate permissions for common AWS resources. It will enumerate redshift, rds, 
and dynamodb.`,
	Run:               awsDatabasesCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	awsCmd.AddCommand(awsDatabasesCmd)
	awsDatabasesCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func awsDatabasesCmdFunc(cmd *cobra.Command, _ []string) {

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

				if _, err := scanner.EnumerateSpecificResource(ctx, region, s, creds, saveOutput); err != nil {
					if errors.Is(err, context.DeadlineExceeded) {
						logger.LoggerStdErr.Error().Str(s.Resource, s.Policy.Permission).Msg("request timed out")
					} else {
						logger.LoggerStdErr.Err(err).Msg("")
					}
					wg.Done()
					return
				}

				wg.Done()

			}(wg, s)

		}
	}()

	resources := []string{
		aws.RDS,
		aws.Redshift,
		aws.Dynamodb,
	}

	enumerate := scanner.GetServiceMap(resources)
	for _, e := range enumerate {
		ch <- e
	}

	close(ch)
	wg.Wait()
}
