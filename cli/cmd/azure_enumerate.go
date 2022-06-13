package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/securisec/cliam/azure"
	"github.com/securisec/cliam/azure/policy"
	"github.com/securisec/cliam/azure/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var azureEnumerateCmd = &cobra.Command{
	Use:     "enumerate [resource...]",
	Example: "enumerate Microsoft.Web.WebApps Microsoft.Web.Sites...",
	Short:   "Enumerate permissions for specified azure resources.",
	PreRun:  azureValidateRequiredFlags,
	Run:     azureEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return azure.GetPolicyKeys(), cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	azureCmd.AddCommand(azureEnumerateCmd)
	azureEnumerateCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func azureEnumerateCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		printValidArgs(azure.GetPolicyKeys)
		os.Exit(1)
	}

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	token, err := azure.GetTokenFromUsernameAndPassword(azureTenantID, azureClientID, azureClientSecret)
	if err != nil {
		logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get access token")
	}
	if azureSubscriptionID == "" {
		azureSubscriptionID, err = azure.GetFirstSubscriptionID(token)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get subscription ID")
		}
	}

	known := policy.CommonPathProperties{
		SubscriptionID:    azureSubscriptionID,
		ResourceGroupName: azureResourceGroupName,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan policy.Policy, 0)
	max := make(chan struct{}, MaxThreads)

	go func() {
		defer wg.Done()

		for s := range ch {
			wg.Add(1)

			go func(wg *sync.WaitGroup, p policy.Policy) {
				max <- struct{}{}
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(RequestTimeout)*time.Second)

				defer func() {
					cancel()
					<-max
					wg.Done()
				}()

				req, err := p.BuildRequest(ctx, token, known)
				if err != nil {
					logger.LogError(err)
					return
				}

				status, body, err := scanner.MakeRequest(req)
				if err != nil {
					logger.LogError(err)
					return
				}

				if status > 399 {
					logger.LogError(fmt.Errorf("%d %s", status, body))
				}

				azureLogSuccessMessage(p, azureKnownResourceMap)

				if saveOutput {
					ioutil.WriteFile(fmt.Sprintf("%s.json", p.OperationID), []byte(body), 0644)
				}

			}(wg, s)

		}
	}()

	azureSendToChannel(ch, args)

	close(ch)
	wg.Wait()
}
