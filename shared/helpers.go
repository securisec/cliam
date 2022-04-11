package shared

import "context"

const (
	CONTENT_TYPE_JSON        = "application/json"
	CONTENT_TYPE_URL_ENCODED = "application/x-www-form-urlencoded"
	USER_AGENT               = "insomnia"
	CONTENT_TYPE_HEADER      = "Content-Type"
)

type Scanner interface {
	EnumerateAll(ctx context.Context, region string, creds interface{}) error
	EnumerateSpecificService(ctx context.Context, region, service string, creds interface{}) error
}
