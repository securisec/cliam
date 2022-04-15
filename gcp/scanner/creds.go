package scanner

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/securisec/cliam/gcp/policy"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

func GetCredsFromServiceAccount(ctx context.Context, saPath string) (*cloudresourcemanager.Service, error) {
	d, err := ioutil.ReadFile(saPath)
	if err != nil {
		return nil, err
	}
	return cloudresourcemanager.NewService(
		ctx,
		option.WithCredentialsJSON(d),
		option.WithScopes(cloudresourcemanager.CloudPlatformScope),
		option.WithTelemetryDisabled(),
	)
}

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
