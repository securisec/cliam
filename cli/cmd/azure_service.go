package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/securisec/cliam/azure"
	"github.com/securisec/cliam/azure/policy"
	"github.com/securisec/cliam/azure/scanner"
	"github.com/securisec/cliam/logger"
	"github.com/spf13/cobra"
)

var azureServiceCmd = &cobra.Command{
	Use:       "service",
	Short:     "Enumerate permissions for a specific group of azure services.",
	PreRun:    azureValidateRequiredFlags,
	Run:       azureServiceCmdFunc,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: azureServiceValidArgs(),
}

func init() {
	azureCmd.AddCommand(azureServiceCmd)
}

func azureServiceCmdFunc(_ *cobra.Command, args []string) {

	var err error
	if len(args) == 0 {
		printValidArgs(azure.GetPolicyKeys)
		os.Exit(1)
	}

	services := azureGetMatchingServices(args[0])

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

				logExtras := ModifyExtraMap(azureKnownResourceMap)
				logExtras["method"] = req.Method
				logExtras["url"] = req.URL.String()
				azureLogSuccessMessage(p, body, logExtras)

				if SaveOutput {
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

	azureSendToChannel(ch, services)

	close(ch)
	wg.Wait()

}

func azureServiceValidArgs() []string {
	services := []string{}
	for _, service := range azure.GetPolicyKeys() {
		s := strings.Split(service, ".")
		if len(s) >= 2 {
			services = append(services, s[1])
		}
	}
	return services
}

func azureGetMatchingServices(service string) []string {
	hold := map[string][]string{}
	for _, service := range azure.GetPolicyKeys() {
		s := strings.Split(service, ".")
		if len(s) >= 2 {
			if _, ok := hold[s[1]]; ok {
				hold[s[1]] = append(hold[s[1]], service)
			} else {
				hold[s[1]] = []string{s[0]}
			}
		}
	}
	return hold[service]
}
