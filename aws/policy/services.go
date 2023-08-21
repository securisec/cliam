package policy

import (
	"bytes"
	"fmt"
	"net/url"
	"path"
	"text/template"

	"github.com/buger/jsonparser"
	"github.com/securisec/cliam/logger"
	"github.com/securisec/cliam/shared"
)

const (
	aws_BASE_URL     = "amazonaws.com"
	aws_JSON_1_1     = "application/x-amz-json-1.1"
	aws_JSON_1_0     = "application/x-amz-json-1.0"
	aws_X_AMZ_TARGET = "X-Amz-Target"
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
	// ExtraValueMap the known word to enumerate additional policies
	ExtraValueMap map[string]string
	// ExtraComponentLocation this could be path, query, header, form, json or uri
	ExtraComponentLocation string
	// ExtraComponentBodyKey is used when it ia form or json data
	ExtraComponentBodyKey string
	// ExtraCommandLineFlag is the cli flag that is used to specifiy the resource in aws.
	// i.e. --bucket-name for s3 buckets
	ExtraCommandLineFlag string
	ReqURL               string
	// for deep scanning
	ResponseParser *ResponseParser
}

// ResponseParser parse response
type ResponseParser struct {
	ResponseFormat string
	// KeysToExtract is the final array key to extract
	KeysToExtract []CommandLineFlagMap
	// ObjectPath is the json keys to the array to extract
	ObjectPath []string
}

// CommandLineFlagMap map command line flags to its response key
type CommandLineFlagMap struct {
	Flag        string
	ResponseKey string
}

// ExtraExtractor extract known values from parsed responses. It works for both xml and json responses.
// `respBytes` is the response body data. `arrayKey` is the final array key to extract. `keys` are the pathway
// to get to array key.
func (r ResponseParser) ExtraExtractor(respBytes []byte) ([]CommandLineFlagMap, error) {
	converted, err := shared.ResponseToJSON(r.ResponseFormat, respBytes)
	if err != nil {
		return nil, err
	}
	// get the array of data following the keys
	dataArray, dataType, _, err := jsonparser.Get(converted, r.ObjectPath...)
	if err != nil {
		return nil, err
	}
	if logger.DEBUG {
		logger.LogDebug("ExtraExtractor", string(dataArray))
	}
	// placeholder to hold extracted data
	var hold []CommandLineFlagMap

	// the value of the inner key can be an object or an array or strings or objects
	switch dataType {

	case jsonparser.Object:
		err = jsonparser.ObjectEach(dataArray, func(key, value []byte, dataType jsonparser.ValueType, offset int) error {
			for _, k := range r.KeysToExtract {
				if string(key) == k.ResponseKey {
					hold = append(hold, CommandLineFlagMap{Flag: k.Flag, ResponseKey: string(value)})
				}
			}
			return err
		})
		if err != nil {
			logger.LogError(err)
		}
		break

	case jsonparser.Array:
		_, err = jsonparser.ArrayEach(dataArray, func(value []byte, arrayDataType jsonparser.ValueType, offset int, err error) {
			// switch to iterate over different datatype
			for _, k := range r.KeysToExtract {
				switch arrayDataType {
				// array of strings
				case jsonparser.String:
					hold = append(hold, CommandLineFlagMap{Flag: k.Flag, ResponseKey: string(value)})
					break

				// array of objects
				case jsonparser.Object:
					data, err := jsonparser.GetString(value, k.ResponseKey)
					if err != nil {
						logger.LogError(err)
					}
					hold = append(hold, CommandLineFlagMap{Flag: k.Flag, ResponseKey: data})
					break
				}
			}
		})
		break
	}

	return hold, err
}

// UpdateForExtra update Service struct to account for extras
func (s Service) UpdateForExtra() (Service, error) {
	if err := s.hasCorrectExtraKey(); err != nil {
		return s, err
	}
	var b bytes.Buffer
	// u, err := url.Parse(s.ReqURL)
	// if err != nil {
	// 	return Service{}, err
	// }
	switch s.ExtraComponentLocation {
	case "path":
		temp := template.Must(template.New("").Parse(s.ReqURL))
		temp.Option("missingkey=error")
		if err := temp.Execute(&b, s.ExtraValueMap); err != nil {
			return Service{}, err
		}
		// check for empty values
		if shared.AwsTemplatePropertyRegex.MatchString(b.String()) {
			return Service{}, fmt.Errorf("empty path")
		}
		s.ReqURL = b.String()
		return s, nil
	case "form":
		if len(s.FormData) == 0 {
			s.FormData = make(map[string]string)
		}
		s.FormData[s.ExtraComponentBodyKey] = s.ExtraValueMap[s.ExtraCommandLineFlag]
		return s, nil
	case "json":
		if len(s.JsonData) == 0 {
			s.JsonData = make(map[string]string)
		}
		s.JsonData[s.ExtraComponentBodyKey] = s.ExtraValueMap[s.ExtraCommandLineFlag]
		return s, nil
	}
	// TOOD add more cases
	return Service{}, fmt.Errorf("unsupported extra component location")
}

// if the passed maps doesnt have the correct flag, return error
func (s Service) hasCorrectExtraKey() error {
	if _, ok := s.ExtraValueMap[s.ExtraCommandLineFlag]; !ok {
		return fmt.Errorf("missing %s", s.ExtraCommandLineFlag)
	}
	return nil
}

// GetRequestURL returns the request url for the service
// This will combine the service, region and service suffix
// to return a valid request url for the permission
func (s *Service) GetRequestURL(region, service, endpoint string) (string, error) {
	var eURL string
	if endpoint != "" {
		pu, err := url.Parse(endpoint)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s://", pu.Scheme) + path.Join(pu.Host, s.ServiceSuffix), nil
	}
	region = AwsRegionSafetyNet(service, region)
	if s.IgnoreRegion {
		eURL = "https://" + path.Join(fmt.Sprintf(
			"%s%s.%s", s.ServicePrefix, service, aws_BASE_URL,
		), s.ServiceSuffix)
		return eURL, nil
	}
	eURL = "https://" + path.Join(fmt.Sprintf(
		"%s%s.%s.%s", s.ServicePrefix, service, region, aws_BASE_URL,
	), s.ServiceSuffix)
	return eURL, nil
}

// AwsRegionSafetyNet some services does not require a region. This function helps fix that.
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
