package scanner

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/securisec/cliam/aws"
	"github.com/securisec/cliam/aws/policy"
	"github.com/securisec/cliam/aws/signer"
)

var (
	MaxThreads = 5
)

// // EnumerateAll scan all services and permissions
// func EnumerateAll(ctx context.Context, region string, creds interface{}) error {
// 	c, ok := creds.(*credentials.Credentials)
// 	if !ok {
// 		return fmt.Errorf("creds must be of type *credentials.Credentials")
// 	}

// 	ch := make(chan ServiceMap)

// 	wg := &sync.WaitGroup{}
// 	wg.Add(MaxThreads)

// 	for i := 0; i < MaxThreads; i++ {
// 		go func() {
// 			for {
// 				s, ok := <-ch
// 				if !ok {
// 					wg.Done()
// 					return
// 				}
// 				_, res, _, err := signer.MakeRequest(ctx, region, s.Resource, &s.Policy, c)
// 				if err != nil {
// 					logger.LogError(err)
// 					wg.Done()
// 					return
// 				}
// 				if res.StatusCode == http.StatusOK {
// 					logger.LogSuccess(s.Resource, s.Policy.Permission)
// 				}
// 			}
// 		}()
// 	}

// 	for service, policies := range aws.Services {
// 		for _, policy := range policies {
// 			ch <- ServiceMap{Resource: service, Policy: policy}
// 		}
// 	}

// 	close(ch)
// 	wg.Wait()

// 	return nil
// }

func EnumerateSpecificResource(
	ctx context.Context,
	region string,
	ser ServiceMap,
	creds *credentials.Credentials,
	saveOutput bool,
) (int, error) {

	_, res, body, err := signer.MakeRequest(ctx, region, ser.Resource, &ser.Policy, creds)
	if err != nil {
		return 0, err
	}

	if res.StatusCode == http.StatusOK {
		if saveOutput {
			saveOutputToFile(ser, body)
		}
	}

	return res.StatusCode, nil
}

// func EnumerateMultipleResources(
// 	ctx context.Context,
// 	region string,
// 	resources []string,
// 	creds interface{},
// 	maxThreads int,
// 	saveOutput bool,
// ) error {
// 	c, ok := creds.(*credentials.Credentials)
// 	if !ok {
// 		return fmt.Errorf("creds must be of type *credentials.Credentials")
// 	}

// 	ch := make(chan ServiceMap)
// 	max := make(chan struct{}, maxThreads)

// 	wg := &sync.WaitGroup{}
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()

// 		for s := range ch {

// 			wg.Add(1)

// 			go func(wg *sync.WaitGroup, s ServiceMap) {

// 				max <- struct{}{}
// 				defer func() {
// 					<-max
// 				}()

// 				_, res, body, err := signer.MakeRequest(ctx, region, s.Resource, &s.Policy, c)
// 				if err != nil {
// 					logger.LogError(err)
// 					wg.Done()
// 					return
// 				}
// 				if res.StatusCode == http.StatusOK {
// 					logger.LogSuccess(s.Resource, s.Policy.Permission)
// 					if saveOutput {
// 						saveOutputToFile(s, body)
// 					}
// 				}
// 				if res.StatusCode != 200 && logger.DEBUG {
// 					logger.LogDenied(res.StatusCode, s.Resource, s.Policy.Permission)
// 				}

// 				wg.Done()
// 			}(wg, s)
// 		}
// 	}()

// 	policies := make(map[string][]policy.Service, 0)
// 	for _, resource := range resources {
// 		if p, ok := aws.Services[resource]; ok {
// 			policies[resource] = p
// 		}
// 	}

// 	for s, ps := range policies {
// 		for _, p := range ps {
// 			ch <- ServiceMap{Resource: s, Policy: p}
// 		}
// 	}

// 	close(ch)
// 	wg.Wait()

// 	return nil
// }

// CheckSpecificPermission check a specific permission based on the provided policy
// This applies to AWS
func CheckSpecificPermission(
	region, service string,
	policy policy.Service,
	creds *credentials.Credentials,
) (*http.Request, *http.Response, []byte, error) {
	return signer.MakeRequest(context.Background(), region, service, &policy, creds)
}

type ServiceMap struct {
	Resource string
	Policy   policy.Service
}

func GetServiceMap(resources []string) []ServiceMap {
	hold := make([]ServiceMap, 0)
	for _, resource := range resources {
		policies, ok := aws.Services[resource]
		if !ok {
			continue
		}
		for _, policy := range policies {
			hold = append(hold, ServiceMap{Resource: resource, Policy: policy})
		}
	}
	return hold
}

func saveOutputToFile(s ServiceMap, body []byte) error {
	fileName := fmt.Sprintf("%s.%s.json", s.Resource, s.Policy.Permission)
	return ioutil.WriteFile(fileName, body, os.ModePerm)
}
