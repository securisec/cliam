package scanner

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/securisec/cliam/gcp/policy"
	"google.golang.org/api/cloudresourcemanager/v1"
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

	op := policy.Service{
		Method:  ps.Method,
		Actions: p.Permissions,
	}

	return op, err
}
