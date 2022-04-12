package policy

import (
	"fmt"
	"path"
)

const (
	aws_BASE_URL          = "amazonaws.com"
	aws_JSON_CONTENT_TYPE = "application/x-amz-json-1.1"
	aws_X_AMZ_TARGET      = "X-Amz-Target"
)

const (
	APIGateway        = "apigateway"
	Cloudtrail        = "cloudtrail"
	Codebuild         = "codebuild"
	Codecommit        = "codecommit"
	Dynamodb          = "dynamodb"
	DynamodbStreams   = "streams.dynamodb"
	EC2               = "ec2"
	ECR               = "ecr"
	ECS               = "ecs"
	EKS               = "eks"
	Elasticache       = "elasticache"
	ElasticBeanStalk  = "elasticbeanstalk"
	ElasticFileSystem = "elasticfilesystem"
	IAM               = "iam"
	Lambda            = "lambda"
	S3                = "s3"
	SQS               = "sqs"
	STS               = "sts"
	Cloudfront        = "cloudfront"
	SNS               = "sns"
	SSM               = "ssm"
	Route53           = "route53"
	RDS               = "rds"
	ELB               = "elasticloadbalancing"
)

var Services = map[string][]Service{
	APIGateway:        APIGatewayPolicies,
	Cloudfront:        CloudfrontPolicies,
	Cloudtrail:        CloudtrailPolicies,
	Codebuild:         CodebuildPolicies,
	Codecommit:        CodecommitPolicies,
	Dynamodb:          DynamoDBPolicies,
	DynamodbStreams:   DynamodbStreamsPolicies,
	EC2:               EC2Policies,
	ECR:               ECRPolicies,
	ECS:               ECSPolicies,
	EKS:               EKSPolicies,
	Elasticache:       ElasticachePolicies,
	ElasticBeanStalk:  ElasticBeanStalkPolicies,
	ElasticFileSystem: ElasticFileSystemPolicies,
	IAM:               IAMPolicies,
	Lambda:            LambdaPolicies,
	S3:                S3Policies,
	SQS:               SQSPolicies,
	STS:               STSPolicies,
	SNS:               SNSPolicies,
	SSM:               SSMPolicies,
	Route53:           Route53Policies,
	RDS:               RDSPolicies,
	ELB:               ELBPolicies,
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
		return url
	}
	url = "https://" + path.Join(fmt.Sprintf(
		"%s%s.%s.%s", s.ServicePrefix, service, region, aws_BASE_URL,
	), s.ServiceSuffix)
	return url
}
