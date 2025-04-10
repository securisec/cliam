package cmd

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
	"github.com/spf13/cobra"
)

var awsSnipeCmd = &cobra.Command{
	Use:     "snipe resource api [flags]",
	Example: "snipe s3 lambda list-functions",
	Short:   "Check a specfic api permissions for an AWS resource.",
	Run:     awsSnipeCmdFunc,
	Args:    cobra.MinimumNArgs(2),
	ValidArgsFunction: func(_ *cobra.Command, args []string, _ string) ([]string, cobra.ShellCompDirective) {
		switch len(args) {
		case 0:
			return aws.GetAWSResources(), cobra.ShellCompDirectiveNoFileComp
		case 1:
			validActions := []string{}
			test := scanner.GetSingleServiceMap(args[0])
			for _, v := range test {
				validActions = append(validActions, shared.CamelToKebabCase(v.Policy.Permission))
			}
			return validActions, cobra.ShellCompDirectiveNoFileComp
		case 2:
			var keys []string
			policies := scanner.GetSingleServiceMap(args[0])
			action := shared.KebabToCamelCase(args[1], true)
			for _, p := range policies {
				if p.Policy.Permission == action {
					if len(p.Policy.ExtraCommandLineFlag) > 0 {
						keys = append(keys, p.Policy.ExtraCommandLineFlag+"=")
					}
					break
				}
			}
			return keys, cobra.ShellCompDirectiveNoSpace
		default:
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
	},
	// PreRun:  awsLoadEnvVarsFirst,
	PostRun: PostRunStatsFunc,
}

func init() {
	awsCmd.AddCommand(awsSnipeCmd)
}

func awsSnipeCmdFunc(_ *cobra.Command, args []string) {
	if awsEndpoint != "" {
		// TODO: remove
		logger.LogWarning("🚧 Custom endpoint enumeration may not be consistent")
	}

	if len(args) < 2 {
		printValidArgs(aws.GetAWSResources)
		os.Exit(1)
	}

	extras := []string{}
	if len(args) == 3 {
		extras = append(extras, args[2])
	}

	key, secret, token, regions := getCredsAndRegion()
	regions = getRegions(regions)
	// cliLogRegion(awsRegion)
	// resources := shared.RemoveDuplicates(args)

	creds := signer.SetCredentials(key, secret, token, awsProfile)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan scanner.ServiceMap)
	max := make(chan struct{}, MaxThreads)

	go func() {
		// defer wg.Done() // 🔥 this seems to resolve the panic errors

		for ser := range ch {

			options := scanner.Options{
				Endpoint:   awsEndpoint,
				Creds:      creds,
				Region:     ser.Region,
				SaveOutput: SaveOutput,
			}
			wg.Add(1)

			go func(wg *sync.WaitGroup, service scanner.ServiceMap, options scanner.Options) {
				max <- struct{}{}
				options.ServiceMap = service
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(RequestTimeout)*time.Second)

				defer func() {
					cancel()
					<-max
				}()

				statusCode, _, err := scanner.EnumerateSpecificResource(ctx, options)
				if err != nil {
					cliErrorLogger(service, err)
					failureCounter++
					wg.Done()
					return
				}
				cliResponseLoggerAWS(service, statusCode, mapToArray(service.Policy.ExtraValueMap), options.Region)

				wg.Done()

			}(wg, ser, options)

		}
	}()

	awsSnipeSendToChannel(ch, args[0], args[1], extras, regions)

	// TODO 🔥 this is blocking if defered, or panicing because closed
	wg.Done() // refacor this because this is poor code and anti pattern
	wg.Wait()

	// TODO 🔥 refactor this and move to aws persistant run. right now it will panic otherwise because of waitgroup
	if saveResults != "" {
		o, _ := json.Marshal(Results)
		if err := os.WriteFile(saveResults, o, os.ModePerm); err != nil {
			logger.LogError(err)
		}
	}
	close(ch)

}

func awsSnipeSendToChannel(ch chan scanner.ServiceMap, resource, permission string, extrasArray, regions []string) {
	permission = shared.KebabToCamelCase(permission, true)
	enumerate := scanner.GetSingleServiceMap(resource)
	sm := scanner.ServiceMap{}
	for _, e := range enumerate {
		if e.Policy.Permission != permission {
			continue
		}
		sm = e
		break
	}

	for _, region := range regions {
		sm.Region = region
		// if a known resource name is set, we will enumerate only the extra permissions
		if len(extrasArray) > 0 {
			sm.Policy.ExtraValueMap = awsModifyExtraMap(ModifyExtraMap(extrasArray))
		}
		ch <- sm
	}
}
