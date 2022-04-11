package policy

import (
	"fmt"
	"path"

	"github.com/securisec/enumerate/logger"
)

const (
	aws_BASE_URL          = "amazonaws.com"
	aws_JSON_CONTENT_TYPE = "application/x-amz-json-1.1"
	aws_X_AMZ_TARGET      = "X-Amz-Target"
)

const (
	ResourceAPIGateway        = "apigateway"
	ResourceCloudtrail        = "cloudtrail"
	ResourceCodebuild         = "codebuild"
	ResourceCodecommit        = "codecommit"
	ResourceDynamodb          = "dynamodb"
	ResourceDynamodbStreams   = "streams.dynamodb"
	ResourceEC2               = "ec2"
	ResourceECR               = "ecr"
	ResourceECS               = "ecs"
	ResourceEKS               = "eks"
	ResourceElasticache       = "elasticache"
	ResourceElasticBeanStalk  = "elasticbeanstalk"
	ResourceElasticFileSystem = "elasticfilesystem"
	ResourceIAM               = "iam"
	ResourceLambda            = "lambda"
	ResourceS3                = "s3"
	ResourceSQS               = "sqs"
	ResourceSTS               = "sts"
	ResourceCloudfront        = "cloudfront"
)

var Services = map[string][]Service{
	ResourceAPIGateway:        APIGatewayPolicies,
	ResourceCloudfront:        CloudfrontPolicies,
	ResourceCloudtrail:        CloudtrailPolicies,
	ResourceCodebuild:         CodebuildPolicies,
	ResourceCodecommit:        CodecommitPolicies,
	ResourceDynamodb:          DynamoDBPolicies,
	ResourceDynamodbStreams:   DynamodbStreamsPolicies,
	ResourceEC2:               EC2Policies,
	ResourceECR:               ECRPolicies,
	ResourceECS:               ECSPolicies,
	ResourceEKS:               EKSPolicies,
	ResourceElasticache:       ElasticachePolicies,
	ResourceElasticBeanStalk:  ElasticBeanStalkPolicies,
	ResourceElasticFileSystem: ElasticFileSystemPolicies,
	ResourceIAM:               IAMPolicies,
	ResourceLambda:            LambdaPolicies,
	ResourceS3:                S3Policies,
	ResourceSQS:               SQSPolicies,
	ResourceSTS:               STSPolicies,
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
			"%s%s.%s", s.ServicePrefix, service, aws_BASE_URL,
		), s.ServiceSuffix)
		logger.LogDebugVerbose("url", url)
		return url
	}
	url = "https://" + path.Join(fmt.Sprintf(
		"%s%s.%s.%s", s.ServicePrefix, service, region, aws_BASE_URL,
	), s.ServiceSuffix)
	logger.LogDebugVerbose("url", url)
	return url
}
