package policy

import (
	"fmt"
	"path"
)

const BASE_AWS_URL = "amazonaws.com"

var Services = map[string][]Service{
	"ec2":    EC2,
	"lambda": Lambda,
	"s3":     S3,
}

// Service base service struct for all services
type Service struct {
	// ServiceSuffix is the suffix of the service
	ServiceSuffix string `json:"serviceSuffix"`
	// Permission camel case permission name
	Permission string `json:"permission"`
	// OptionalQuery for future needs
	OptionalQueryParams map[string]string `json:"optionalQueryParams,omitempty"`
}

// GetRequestURL returns the request url for the service
// This will combine the service, region and service suffix
// to return a valid request url for the permission
func (s *Service) GetRequestURL(region, service string) string {
	return "https://" + path.Join(fmt.Sprintf(
		"%s.%s.%s", service, region, BASE_AWS_URL,
	), s.ServiceSuffix)
}
