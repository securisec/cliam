package gcp

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/securisec/cliam/gcp/policy"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

// GetAccessToken returns a valid access token for the provided
// service account.
func GetAccessToken(ctx context.Context, serviceAccountPath string) (string, error) {
	j, err := ioutil.ReadFile(serviceAccountPath)
	if err != nil {
		return "", err
	}
	creds, err := google.CredentialsFromJSON(ctx, j, compute.CloudPlatformScope)
	// creds, err := transport.Creds(ctx)
	if err != nil {
		return "", err
	}
	token, err := creds.TokenSource.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

// GetGCPResources get all GCP resources
func GetGCPResources() []string {
	keys := make([]string, 0, len(policy.Resources))
	for k := range policy.Resources {
		keys = append(keys, k)
	}
	return keys
}

// GetPolicyForResource returns the policies for the given resource
func GetPolicyForResource(key string) []string {
	p, ok := policy.Resources[key]
	if !ok {
		return []string{}
	}
	hold := make([]string, 0, len(p.Actions))
	for _, v := range p.Actions {
		hold = append(hold, fmt.Sprintf("%s.%s", p.Method, v))
	}
	return hold
}
