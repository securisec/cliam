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

func EnumerateSpecificResource(
	ctx context.Context,
	region string,
	ser ServiceMap,
	creds *credentials.Credentials,
	saveOutput bool,
	word string,
) (int, error) {

	p := ser.Policy
	requestURL := p.GetRequestURL(region, ser.Resource)

	// TODO update Policy here if known stuff is set

	_, res, body, err := signer.MakeRequest(ctx, region, ser.Resource, &ser.Policy, creds, requestURL)
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
