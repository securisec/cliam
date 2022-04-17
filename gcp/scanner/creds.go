package scanner

import (
	"context"
	"io/ioutil"

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
		option.WithTelemetryDisabled(),
	)
}
