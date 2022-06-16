package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	azp "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
)

// GetTokenFromDefault get token from currently logged in user
func GetTokenFromDefault() (string, error) {
	az, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return "", err
	}
	token, err := az.GetToken(context.Background(), azp.TokenRequestOptions{
		Scopes: []string{"https://management.core.windows.net//.default"},
	})
	if err != nil {
		return "", err
	}
	return token.Token, nil
}

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

func GetTokenFromCertificate(tenantID, clientID, certPath string) (string, error) {
	data, err := ioutil.ReadFile(certPath)
	if err != nil {
		return "", err
	}
	certs, key, err := azidentity.ParseCertificates(data, nil)
	if err != nil {
		return "", err
	}
	az, err := azidentity.NewClientCertificateCredential(tenantID, clientID, certs, key, nil)
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

// GetFirstSubscriptions returns the first subscription ID from the given.
// The token can be used to make REST api calls. This function returns a subscription ID
// and a list of all subscriptions.
func GetFirstSubscriptions(bearerToken string) (string, []SubscriptionValue, error) {
	req, err := http.NewRequest("GET", "https://management.azure.com/subscriptions?api-version=2020-01-01", nil)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Authorization", "Bearer "+bearerToken)

	res, err := http.DefaultClient.Do(req)
	// read body
	if err != nil {
		return "", nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	var subs SubscriptionResponse
	if err := json.NewDecoder(res.Body).Decode(&subs); err != nil {
		return "", nil, err
	}
	if subs.Count.Value <= 0 {
		return "", nil, fmt.Errorf("no subscriptions found")
	}

	id := subs.Value[0].SubscriptionID

	if logger.DEBUG {
		s := make([]string, 0, len(subs.Value))
		for _, sub := range subs.Value {
			s = append(s, sub.SubscriptionID)
		}
		logger.LoggerStdErr.Warn().Strs("subscriptions", s).Msg(shared.GetMessageColor("debug"))
	}

	return id, subs.Value, nil
}

type SubscriptionResponse struct {
	Value []SubscriptionValue `json:"value"`
	Count Count               `json:"count"`
}

type Count struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type SubscriptionValue struct {
	ID                   string               `json:"id"`
	AuthorizationSource  string               `json:"authorizationSource"`
	ManagedByTenants     []interface{}        `json:"managedByTenants"`
	SubscriptionID       string               `json:"subscriptionId"`
	TenantID             string               `json:"tenantId"`
	DisplayName          string               `json:"displayName"`
	State                string               `json:"state"`
	SubscriptionPolicies SubscriptionPolicies `json:"subscriptionPolicies"`
}

type SubscriptionPolicies struct {
	LocationPlacementID string `json:"locationPlacementId"`
	QuotaID             string `json:"quotaId"`
	SpendingLimit       string `json:"spendingLimit"`
}
