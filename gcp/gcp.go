package gcp

import (
	"context"
	"io/ioutil"

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
