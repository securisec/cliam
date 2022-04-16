package scanner

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/securisec/cliam/gcp/policy"
	"google.golang.org/api/cloudresourcemanager/v1"
)

func EnumerateSpecificResource(
	ctx context.Context,
	creds *cloudresourcemanager.Service,
	projectId, resource string,
) (policy.Service, error) {
	return GetPermissionsForResource(ctx, creds, projectId, policy.Resources[resource])
}

func EnumerateMultipleResources(
	ctx context.Context,
	options *GCPEnumOptions,
	resources ...string,
) ([]policy.Service, error) {
	var permissions []policy.Service

	for _, resource := range resources {
		r, err := GetPermissionsForResource(ctx, options.Creds, options.ProjectId, policy.Resources[resource])
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, r)
	}
	return permissions, nil
}

func EnumerateAllResources(
	ctx context.Context,
	creds *cloudresourcemanager.Service,
	projectId string,
) ([]policy.Service, error) {
	var permissions []policy.Service
	for _, p := range policy.Resources {
		r, err := GetPermissionsForResource(ctx, creds, projectId, p)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, r)
	}
	return permissions, nil
}

type GCPEnumOptions struct {
	Creds     *cloudresourcemanager.Service
	ProjectId string
}

// EnumerateRestApiRequest is a request to enumerate permissions using the REST API
// care should be paid to r policy.Resources
func EnumerateRestApiRequest(
	ctx context.Context,
	accessToken string,
	r policy.RestCall,
) (policy.RestCall, error) {
	var req *http.Request

	url, err := r.GetURL()
	if err != nil {
		return r, err
	}

	isGet := r.ReqMethod == "GET"

	if isGet {
		req, err = http.NewRequest(r.ReqMethod, url, nil)
	} else {
		o, err := json.Marshal(r.ReqBody)
		if err != nil {
			return r, err
		}
		req, err = http.NewRequest(r.ReqMethod, url, bytes.NewBuffer(o))
	}
	if err != nil {
		return r, err
	}
	if !isGet {
		req.Header.Add("Content-Type", "application/json")
	}

	// add auth bearer token
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("user-agent", "insomnia/2022.2.1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return r, err
	}

	r.Response = res
	defer res.Body.Close()
	return r, nil
}
