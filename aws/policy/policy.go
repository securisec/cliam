package policy

import (
	"fmt"
	"path"

	"github.com/securisec/enumerate/logger"
)

const (
	BASE_AWS_URL          = "amazonaws.com"
	AWS_JSON_CONTENT_TYPE = "application/x-amz-json-1.1"
	AWS_X_AMZ_TARGET      = "X-Amz-Target"
)

var Services = map[string][]Service{
	"dynamodb":         DynamoDB,
	"streams.dynamodb": DynamodbStreams,
	"ec2":              EC2,
	"ecr":              ECR,
	"iam":              IAM,
	"lambda":           Lambda,
	"s3":               S3,
	"sqs":              SQS,
	"sts":              STS,
}

// Service base service struct for all services
type Service struct {
	// ServiceSuffix is the suffix of the service
	ServiceSuffix string `json:"serviceSuffix"`
	// ServicePrefix is the prefix of the service
	ServicePrefix string `json:"servicePrefix"`
	// Permission camel case permission name
	Permission string `json:"permission"`
	// OptionalQuery for future needs
	OptionalQueryParams map[string]string `json:"optionalQueryParams,omitempty"`
	// Method request method. Default is GET
	Method string `json:"method"`
	// FormData for form based requests
	FormData map[string]string `json:"formData,omitempty"`
	// JsonData for json based requests
	JsonData string `json:"jsonData,omitempty"`
	// Headers additional headers
	Headers map[string]string `json:"headers,omitempty"`
	// IgnoreRegion some services are global
	IgnoreRegion bool `json:"ignoreRegion,omitempty"`
}

// GetRequestURL returns the request url for the service
// This will combine the service, region and service suffix
// to return a valid request url for the permission
func (s *Service) GetRequestURL(region, service string) string {
	var url string
	if s.IgnoreRegion {
		url = "https://" + path.Join(fmt.Sprintf(
			"%s%s.%s", s.ServicePrefix, service, BASE_AWS_URL,
		), s.ServiceSuffix)
		logger.LogDebug("url", url)
		return url
	}
	url = "https://" + path.Join(fmt.Sprintf(
		"%s%s.%s.%s", s.ServicePrefix, service, region, BASE_AWS_URL,
	), s.ServiceSuffix)
	logger.LogDebug("url", url)
	return url
}
