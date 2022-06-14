package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	azp "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

// GetTokenFromUsernameAndPassword returns a valid bearer token from the given
// username, password and tenant. The token can be used to make REST api calls.
func GetTokenFromUsernameAndPassword(tenantID, username, password string) (string, error) {
	az, err := azidentity.NewClientSecretCredential(tenantID, username, password, nil)
	if err != nil {
		return "", err
	}

	token, err := az.GetToken(context.Background(), azp.TokenRequestOptions{
		Scopes: []string{"https://management.core.windows.net//.default"}})
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

// GetFirstSubscriptionID returns the first subscription ID from the given.
// The token can be used to make REST api calls.
func GetFirstSubscriptionID(bearerToken string) (string, error) {
	req, err := http.NewRequest("GET", "https://management.azure.com/subscriptions?api-version=2020-01-01", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	res, err := http.DefaultClient.Do(req)
	// read body
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	var subs subscriptionResponse
	if err := json.NewDecoder(res.Body).Decode(&subs); err != nil {
		return "", err
	}
	if subs.Count.Value <= 0 {
		return "", fmt.Errorf("no subscriptions found")
	}

	return subs.Value[0].SubscriptionID, nil
}

type subscriptionResponse struct {
	Value []struct {
		SubscriptionID string `json:"subscriptionId"`
	} `json:"value"`
	Count struct {
		Value int `json:"value"`
	} `json:"count"`
}
