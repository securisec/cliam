package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var gcpEnumerateCmd = &cobra.Command{
	Use:   "enumerate [resource...]",
	Short: "Enumerate specified GCP permissions",
	Long: `Enumerate uses cloudresource manager to check if the service account has the specified permissions. 
If resource manager is disabled, this command will end with a 403 error.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			printValidArgs(gcp.GetGCPResources)
			os.Exit(1)
		}
	},
	Run: gcpEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return gcp.GetGCPResources(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	gcpCmd.AddCommand(gcpEnumerateCmd)
}

func gcpEnumerateCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	}

	accessToken := gcpAccessToken

	services := removeDuplicates(args)

	sa, project, _, _ := getSaAndRegion()

	ctx := context.Background()

	if accessToken == "" {
		at, err := gcp.GetAccessToken(ctx, sa)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get credentials from service account")
			return
		}
		accessToken = at
	}

	ch := make(chan string)
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

				ps, err := scanner.GetPermissionsFromResourceManager(ctx, accessToken, project, ser)
				if err != nil {
					logger.LoggerStdErr.Error().Err(err).Msg("")
					wg.Done()
					return
				}

				for _, p := range ps {
					x := strings.Split(p, ".")
					logger.LogSuccess(fmt.Sprintf("%s.%s", x[0], x[1]), x[2])
				}

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
