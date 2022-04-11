package scanner

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/securisec/enumerate/policy"
	"github.com/securisec/enumerate/signer"
)

var (
	MaxThreads = 10
)

// EnumerateAll scan all services and permissions
func EnumerateAll(ctx context.Context, region string, creds *credentials.Credentials) {
	ch := make(chan serviceMap)

	wg := &sync.WaitGroup{}
	wg.Add(MaxThreads)

	for i := 0; i < MaxThreads; i++ {
		go func() {
			for {
				s, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				signer.MakeRequest(ctx, region, s.Service, &s.Policy, creds)
			}
		}()
	}

	for service, policies := range policy.Services {
		for _, policy := range policies {
			ch <- serviceMap{Service: service, Policy: policy}
		}
	}

	close(ch)
	wg.Wait()
}

func ScanSpecificService(ctx context.Context, region, service string, creds *credentials.Credentials) error {
	ch := make(chan serviceMap)

	wg := &sync.WaitGroup{}
	wg.Add(MaxThreads)

	for i := 0; i < MaxThreads; i++ {
		go func() {
			for {
				s, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				signer.MakeRequest(ctx, region, s.Service, &s.Policy, creds)
			}
		}()
	}

	policies, ok := policy.Services[service]
	if !ok {
		return fmt.Errorf("service %s not found", service)
	}

	for _, policy := range policies {
		ch <- serviceMap{Service: service, Policy: policy}
	}

	close(ch)
	wg.Wait()

	return nil
}

type serviceMap struct {
	Service string
	Policy  policy.Service
}
