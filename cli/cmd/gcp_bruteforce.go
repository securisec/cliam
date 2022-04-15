package cmd

import (
	"context"
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
	gcpBruteforceCmd.Flags().Int("max-threads", 5, "Maximum number of threads to use")
}

func gcpBruteforceCmdFunc(cmd *cobra.Command, _ []string) {
	maxThreads, _ := cmd.Flags().GetInt("max-threads")

	ctx := context.Background()
	sa, project, _ := getSaAndRegion()
	creds, err := scanner.GetCredsFromServiceAccount(ctx, sa)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get credentials from service account")
		return
	}
	services := gcp.GetGCPResources()

	ch := make(chan string)

	wg := &sync.WaitGroup{}
	wg.Add(maxThreads)

	options := scanner.GCPEnumOptions{
		Creds:     creds,
		ProjectId: project,
	}

	for i := 0; i < maxThreads; i++ {
		go func() {

			for {
				s, ok := <-ch
				if !ok {
					wg.Done()
					return
				}

				ps, err := scanner.EnumerateMultipleResources(ctx, &options, s)
				if err != nil {
					logger.LoggerStdErr.Fatal().Err(err).Msg("")
					wg.Done()
					return
				}
				for _, p := range ps {
					for _, a := range p.Actions {
						a = strings.Split(a, ".")[2]
						logger.LogSuccess(p.Method, a)
					}
				}

			}
		}()
	}

	for _, s := range services {
		ch <- s
	}

	close(ch)
	wg.Wait()

}
