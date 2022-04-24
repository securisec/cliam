package policy

import (
	"bytes"
	"fmt"
	"path"
	"text/template"
)

const (
	aws_BASE_URL          = "amazonaws.com"
	aws_JSON_CONTENT_TYPE = "application/x-amz-json-1.1"
	aws_JSON_1            = "application/x-amz-json-1.0"
	aws_X_AMZ_TARGET      = "X-Amz-Target"
)

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
	JsonData map[string]string `json:"jsonData,omitempty"`
	// Headers additional headers
	Headers map[string]string `json:"headers,omitempty"`
	// IgnoreRegion some services are global
	IgnoreRegion bool `json:"ignoreRegion,omitempty"`
	// QueryParams query params
	QueryParams map[string]string `json:"queryParams,omitempty"`
	// IsExtra allows passing a word to enumerate additional policies
	IsExtra bool
	// ExtraWord the known word to enumerate additional policies
	ExtraWord string
	// ExtraComponentLocation this could be path, query, header, form, json or uri
	ExtraComponentLocation string
	// ExtraComponentBodyKey is used when it ia form or json data
	ExtraComponentBodyKey string
	ReqURL                string
}

func (s Service) UpdateForExtra() (Service, error) {
	var b bytes.Buffer
	// u, err := url.Parse(s.ReqURL)
	// if err != nil {
	// 	return Service{}, err
	// }
	switch s.ExtraComponentLocation {
	case "path":
		temp := template.Must(template.New("").Parse(s.ReqURL))
		if err := temp.Execute(&b, s.ExtraWord); err != nil {
			return Service{}, err
		}
		s.ReqURL = b.String()
		return s, nil
	}
	// TOOD add more cases
	return Service{}, fmt.Errorf("unsupported extra component location")
}

// GetRequestURL returns the request url for the service
// This will combine the service, region and service suffix
// to return a valid request url for the permission
func (s *Service) GetRequestURL(region, service string) string {
	var url string
	region = AwsRegionSafetyNet(service, region)
	if s.IgnoreRegion {
		url = "https://" + path.Join(fmt.Sprintf(
			"%s%s.%s", s.ServicePrefix, service, aws_BASE_URL,
		), s.ServiceSuffix)
		return url
	}
	url = "https://" + path.Join(fmt.Sprintf(
		"%s%s.%s.%s", s.ServicePrefix, service, region, aws_BASE_URL,
	), s.ServiceSuffix)
	return url
}

func AwsRegionSafetyNet(service, region string) string {
	switch service {
	case "iam":
		return "us-east-1"
	case "cloudfront":
		return "us-east-1"
	case "route53":
		return "us-east-1"
	case "ecr-public":
		return "us-east-1"
	default:
		return region
	}
}
