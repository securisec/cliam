package scanner

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/policy"
	"github.com/securisec/cliam/aws/signer"
	"github.com/securisec/cliam/logger"
)

var (
	MaxThreads = 5
)

// EnumerateAll scan all services and permissions
func EnumerateAll(ctx context.Context, region string, creds interface{}) error {
	c, ok := creds.(*credentials.Credentials)
	if !ok {
		return fmt.Errorf("creds must be of type *credentials.Credentials")
	}

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
				_, res, err := signer.MakeRequest(ctx, region, s.Service, &s.Policy, c)
				if err != nil {
					logger.LogError(err)
					wg.Done()
					return
				}
				if res.StatusCode == http.StatusOK {
					logger.LogSuccess(s.Service, s.Policy.Permission)
				}
			}
		}()
	}

	for service, policies := range aws.Services {
		for _, policy := range policies {
			ch <- serviceMap{Service: service, Policy: policy}
		}
	}

	close(ch)
	wg.Wait()

	return nil
}

func EnumerateSpecificResource(ctx context.Context, region, resource string, creds interface{}) error {
	c, ok := creds.(*credentials.Credentials)
	if !ok {
		return fmt.Errorf("creds must be of type *credentials.Credentials")
	}

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
				_, res, err := signer.MakeRequest(ctx, region, s.Service, &s.Policy, c)
				if err != nil {
					logger.LogError(err)
					wg.Done()
					return
				}
				if res.StatusCode == http.StatusOK {
					logger.LogSuccess(s.Service, s.Policy.Permission)
				}
				if logger.DEBUG {
					logger.LogDenied(res.StatusCode, s.Service, s.Policy.Permission)
				}
			}
		}()
	}

	policies, ok := aws.Services[resource]
	if !ok {
		return fmt.Errorf("service %s not found", resource)
	}

	for _, policy := range policies {
		ch <- serviceMap{Service: resource, Policy: policy}
	}

	close(ch)
	wg.Wait()

	return nil
}

func EnumerateMultipleResources(
	ctx context.Context,
	region string,
	resources []string,
	creds interface{},
	saveOutput *bool,
) error {
	c, ok := creds.(*credentials.Credentials)
	if !ok {
		return fmt.Errorf("creds must be of type *credentials.Credentials")
	}

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
				_, res, err := signer.MakeRequest(ctx, region, s.Service, &s.Policy, c)
				if err != nil {
					logger.LogError(err)
					wg.Done()
					return
				}
				if res.StatusCode == http.StatusOK {
					logger.LogSuccess(s.Service, s.Policy.Permission)
					if *saveOutput {
						saveOutputToFile(s, res)
					}
				}
			}
		}()
	}

	policies := make(map[string][]policy.Service, 0)
	for _, resource := range resources {
		if p, ok := aws.Services[resource]; ok {
			policies[resource] = p
		}
	}

	for s, ps := range policies {
		for _, p := range ps {
			ch <- serviceMap{Service: s, Policy: p}
		}
	}

	close(ch)
	wg.Wait()

	return nil
}

// CheckSpecificPermission check a specific permission based on the provided policy
// This applies to AWS
func CheckSpecificPermission(
	region, service string,
	policy policy.Service,
	creds *credentials.Credentials,
) (*http.Request, *http.Response, error) {
	return signer.MakeRequest(context.Background(), region, service, &policy, creds)
}

type serviceMap struct {
	Service string
	Policy  policy.Service
}

func saveOutputToFile(s serviceMap, res *http.Response) error {
	fileName := fmt.Sprintf("%s.%s.json", s.Service, s.Policy.Permission)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, body, os.ModePerm)
}
