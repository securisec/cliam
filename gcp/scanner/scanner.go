package scanner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/securisec/cliam/gcp/policy"
	"github.com/securisec/cliam/gcp/rest"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/googleapi"
)

func GetPermissionsForResource(
	ctx context.Context,
	service *cloudresourcemanager.Service,
	projectId string,
	ps policy.Service,
) (policy.Service, error) {
	var permissions []string
	for _, action := range ps.Actions {
		permissions = append(permissions, fmt.Sprintf("%s.%s", ps.Method, action))
	}

	p, err := service.Projects.TestIamPermissions(projectId, &cloudresourcemanager.TestIamPermissionsRequest{
		Permissions: permissions,
	}).Do()

	if err != nil {
		e, ok := err.(*googleapi.Error)
		if !ok {
			return policy.Service{}, err
		}
		return policy.Service{}, errors.New(e.Message)
	}

	op := policy.Service{
		Method:  ps.Method,
		Actions: p.Permissions,
	}

	return op, nil
}

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

// func EnumerateAllResources(
// 	ctx context.Context,
// 	creds *cloudresourcemanager.Service,
// 	projectId string,
// ) ([]policy.Service, error) {
// 	var permissions []policy.Service
// 	for _, p := range policy.Resources {
// 		r, err := GetPermissionsForResource(ctx, creds, projectId, p)
// 		if err != nil {
// 			return nil, err
// 		}
// 		permissions = append(permissions, r)
// 	}
// 	return permissions, nil
// }

type GCPEnumOptions struct {
	Creds     *cloudresourcemanager.Service
	ProjectId string
}

// EnumerateRestApiRequest is a request to enumerate permissions using the REST API
// care should be paid to r policy.Resources
func EnumerateRestApiRequest(
	ctx context.Context,
	accessToken string,
	r rest.RestCall,
) (rest.RestCall, []byte, error) {
	var req *http.Request

	url, err := r.GetURL()
	if err != nil {
		return r, nil, err
	}

	isGet := r.ReqMethod == "GET"

	if isGet {
		req, err = http.NewRequest(r.ReqMethod, url, nil)
	} else {
		o, err := json.Marshal(r.ReqBody)
		if err != nil {
			return r, nil, err
		}
		req, err = http.NewRequest(r.ReqMethod, url, bytes.NewBuffer(o))
	}
	if err != nil {
		return r, nil, err
	}
	if !isGet {
		req.Header.Add("Content-Type", "application/json")
	}

	// add auth bearer token
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("user-agent", "google-cloud-sdk gcloud/379.0.0")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return r, nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return r, nil, err
	}

	r.Response = res
	defer res.Body.Close()
	return r, body, nil
}
