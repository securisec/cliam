package cmd

import (
	"context"
	"sync"
	"time"

	"github.com/securisec/cliam/aws/scanner"
	"github.com/securisec/cliam/aws/signer"
)

func awsSharedEnumerate(resources []string, saveOutput bool) {
	key, secret, token, region := getCredsAndRegion()
	cliLogRegion(awsRegion)

	creds := signer.SetCredentials(key, secret, token, awsProfile)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan scanner.ServiceMap)
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

				statusCode, _, err := scanner.EnumerateSpecificResource(ctx, region, awsEndpoint, s, creds, saveOutput)
				if err != nil {
					cliErrorLogger(s, err)
					failureCounter++
					wg.Done()
					return
				}
				cliResponseLoggerAWS(s, statusCode)

				wg.Done()

			}(wg, s)

		}
	}()

	awsSendToChannel(ch, resources)

	close(ch)
	wg.Wait()
}
