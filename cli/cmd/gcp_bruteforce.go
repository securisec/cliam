package cmd

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var gcpBruteforceCmd = &cobra.Command{
	Use:               "bruteforce",
	Short:             "Enumerate all GCP permissions",
	Run:               gcpBruteforceCmdFunc,
	ValidArgsFunction: cobra.NoFileCompletions,
}

func init() {
	gcpCmd.AddCommand(gcpBruteforceCmd)
}

func gcpBruteforceCmdFunc(cmd *cobra.Command, _ []string) {

	ctx := context.Background()
	sa, project, region, zone := getSaAndRegion()
	creds, err := scanner.GetCredsFromServiceAccount(ctx, sa)
	cliGcpLogRegion(region, zone)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get credentials from service account")
		return
	}
	services := gcp.GetGCPResources()

	ch := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	// the is the maximum concurrent goroutines
	max := make(chan struct{}, MaxThreads)

	go func() {
		defer wg.Done()
		for s := range ch {

			wg.Add(1)

			go func(wg *sync.WaitGroup, ser string) {

				// block the channel until the maximum number of goroutines is reached
				max <- struct{}{}
				defer func() {
					<-max
				}()

				ps, err := scanner.GetPermissionsForResource(ctx, creds, project, ser)
				if err != nil {
					logger.LoggerStdErr.Error().Err(err).Msg("")
					wg.Done()
					return
				}

				for _, p := range ps {
					x := strings.Split(p, ".")
					logger.LogSuccess(fmt.Sprintf("%s.%s", x[0], x[1]), x[2])
				}

				// done with this goroutine
				wg.Done()
			}(wg, s)
		}

	}()

	for _, s := range services {
		ch <- s
	}

	close(ch)
	wg.Wait()

}
