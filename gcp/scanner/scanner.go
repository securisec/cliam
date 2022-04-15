package scanner

import (
	"context"

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
