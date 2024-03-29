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
	PostRun:           PostRunStatsFunc,
}

func init() {
	gcpCmd.AddCommand(gcpBruteforceCmd)
}

func gcpBruteforceCmdFunc(_ *cobra.Command, _ []string) {
	accessToken := gcpAccessToken

	ctx := context.Background()
	sa, project, region, zone := getSaAndRegion()
	cliGcpLogRegion(map[string]string{
		"region":  region,
		"zone":    zone,
		"project": project,
	})

	if accessToken == "" {
		at, err := gcp.GetAccessToken(ctx, sa)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get credentials from service account")
			return
		}
		accessToken = at
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

				ps, err := scanner.GetPermissionsFromResourceManager(ctx, accessToken, project, ser)
				if err != nil {
					logger.LoggerStdErr.Error().Err(err).Msg("")
					failureCounter++
					wg.Done()
					return
				}

				for _, p := range ps {
					x := strings.Split(p, ".")
					successCounter++
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
