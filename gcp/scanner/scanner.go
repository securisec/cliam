package scanner

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/securisec/cliam/gcp"
	"github.com/securisec/cliam/gcp/rest"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/googleapi"
)

func GetPermissionsForResource(
	ctx context.Context,
	service *cloudresourcemanager.Service,
	projectId, resource string,
) ([]string, error) {
	var holdSuccess []string
	permissions := gcp.GetPoliciesForResource(resource)

	chunkedPolicies := chunkPolicies(permissions, 100)

	for _, chunk := range chunkedPolicies {
		p, err := service.Projects.TestIamPermissions(projectId, &cloudresourcemanager.TestIamPermissionsRequest{
			Permissions: chunk,
		}).Do()
		if err != nil {
			e, ok := err.(*googleapi.Error)
			if !ok {
				return holdSuccess, err
			}
			return holdSuccess, errors.New(e.Message)
		}
		holdSuccess = append(holdSuccess, p.Permissions...)
	}

	return holdSuccess, nil
}

func chunkPolicies(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for {
		if len(slice) == 0 {
			break
		}

		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}

// func EnumerateSpecificResource(
// 	ctx context.Context,
// 	creds *cloudresourcemanager.Service,
// 	projectId, resource string,
// ) (policy.Service, error) {
// 	return GetPermissionsForResource(ctx, creds, projectId, policy.Resources[resource])
// }

// func EnumerateMultipleResources(
// 	ctx context.Context,
// 	options *GCPEnumOptions,
// 	resources ...string,
// ) ([]policy.Service, error) {
// 	var permissions []policy.Service

// 	for _, resource := range resources {
// 		r, err := GetPermissionsForResource(ctx, options.Creds, options.ProjectId, policy.Resources[resource])
// 		if err != nil {
// 			return nil, err
// 		}
// 		permissions = append(permissions, r)
// 	}
// 	return permissions, nil
// }

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
		req, err = http.NewRequestWithContext(ctx, r.ReqMethod, url, nil)
	} else {
		o, err := json.Marshal(r.ReqBody)
		if err != nil {
			return r, nil, err
		}
		req, err = http.NewRequestWithContext(ctx, r.ReqMethod, url, bytes.NewBuffer(o))
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
	res.Body.Close()
	return r, body, nil
}
