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

func EnumerateSpecificResource(
	ctx context.Context,
	region string,
	ser ServiceMap,
	creds *credentials.Credentials,
	saveOutput bool,
) (int, error) {
	ser.Policy.ReqURL = ser.Policy.GetRequestURL(region, ser.Resource)

	if ser.Policy.IsExtra {
		s, err := ser.Policy.UpdateForExtra()
		if err != nil {
			return 0, err
		}
		ser.Policy = s
	}

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
