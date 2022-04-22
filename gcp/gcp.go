package gcp

import (
	"context"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/securisec/cliam/gcp/policy"
	"github.com/securisec/cliam/gcp/rest"
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
	sort.Sort(sort.StringSlice(keys))
	return keys
}

// GetPoliciesForResource returns the policies for the given resource
func GetPoliciesForResource(key string) []string {
	policies, ok := policy.Resources[key]
	if !ok {
		return []string{}
	}
	permissions := make([]string, 0)
	for _, policy := range policies {
		for _, action := range policy.Actions {
			permissions = append(permissions, fmt.Sprintf("%s.%s", policy.Method, action))
		}
	}
	return permissions
}

// GetAvailableRestKeys returns the policies for the given resource
func GetAvailableRestKeys() []string {
	hold := make([]string, 0)
	for k := range rest.RestApiCalls {
		hold = append(hold, k)
	}
	sort.Sort(sort.StringSlice(hold))
	return hold
}
