package cmd

import (
	"context"
	"encoding/json"
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
	Use:     "enumerate [resource...] [-k key=value...]",
	Example: "enumerate Microsoft.Web.WebApps Microsoft.Web.Sites...",
	Short:   "Enumerate permissions for specified azure resources.",
	PreRun:  azureValidateRequiredFlags,
	Run:     azureEnumerateCmdFunc,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return azure.GetPolicyKeys(), cobra.ShellCompDirectiveNoFileComp
	},
	PostRun: azurePostRunFunc,
}

func init() {
	azureCmd.AddCommand(azureEnumerateCmd)
	azureEnumerateCmd.Flags().Bool("save-output", false, "Save output to file on success")
}

func azureEnumerateCmdFunc(cmd *cobra.Command, args []string) {
	var err error
	if len(args) == 0 {
		printValidArgs(azure.GetPolicyKeys)
		os.Exit(1)
	}

	saveOutput, _ := cmd.Flags().GetBool("save-output")

	azureOauthToken = azureGetOauthToken()

	if azureSubscriptionID == "" {
		azureSubscriptionID, _, err = azure.GetFirstSubscriptions(azureOauthToken)
		if err != nil {
			logger.LoggerStdErr.Fatal().Err(err).Msg("Failed to get subscription ID")
		}
	}

	known := policy.CommonPathProperties{
		SubscriptionID:    azureSubscriptionID,
		ResourceGroupName: azureResourceGroupName,
	}
	if known.ResourceGroupName == "" {
		logger.LoggerStdErr.Warn().Msg("No resource group name specified, results will be limited.")
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

				req, err := p.BuildRequest(ctx, azureOauthToken, known)
				if err != nil {
					failureCounter += 1
					logger.LogError(err, p.OperationID)
					return
				}

				res, body, err := scanner.MakeRequest(req)
				if err != nil {
					failureCounter += 1
					logger.LogError(err, p.OperationID)
					return
				}

				if res.StatusCode > 399 {
					failureCounter += 1
					logger.LogError(fmt.Errorf("%d %s", res.StatusCode, body), p.OperationID)
				}

				logExtras := azureModifyExtraMap(azureKnownResourceMap)
				logExtras["method"] = req.Method
				logExtras["url"] = req.URL.String()
				azureLogSuccessMessage(p, body, logExtras)

				if saveOutput {
					o, err := json.Marshal(body)
					if err != nil {
						logger.LogError(err)
						return
					}
					ioutil.WriteFile(fmt.Sprintf("%s.json", p.OperationID), o, 0644)
				}

			}(wg, s)

		}
	}()

	azureSendToChannel(ch, args)

	close(ch)
	wg.Wait()
}
