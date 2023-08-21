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
	"github.com/securisec/cliam/shared"
)

// Options config for enumeration scanner
type Options struct {
	Region, Endpoint string
	ServiceMap       ServiceMap
	Creds            *credentials.Credentials
	SaveOutput       bool
}

// EnumerateSpecificResource enumerate
func EnumerateSpecificResource(
	ctx context.Context,
	config Options,
) (int, []byte, error) {
	u, err := config.ServiceMap.Policy.GetRequestURL(config.Region, config.ServiceMap.Resource, config.Endpoint)
	if err != nil {
		return 0, nil, err
	}
	config.ServiceMap.Policy.ReqURL = u

	if config.ServiceMap.Policy.IsExtra {
		s, err := config.ServiceMap.Policy.UpdateForExtra()
		if err != nil {
			return 0, nil, err
		}
		config.ServiceMap.Policy = s
	}

	_, res, body, err := signer.MakeScannerRequest(ctx, config.Region, config.ServiceMap.Resource, &config.ServiceMap.Policy, config.Creds)
	if err != nil {
		return 0, nil, err
	}

	if res.StatusCode == http.StatusOK {
		if config.SaveOutput {
			saveOutputToFile(config.ServiceMap, body)
		}
	}

	return res.StatusCode, body, nil
}

type ServiceMap struct {
	Resource string
	Policy   policy.Service
}

func GetServiceMap(resources []string) []ServiceMap {
	rm := shared.RemoveDuplicates(resources)
	hold := make([]ServiceMap, 0)
	for _, resource := range rm {
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
