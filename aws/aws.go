package aws

import "github.com/securisec/iam-enumerate/aws/policy"

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
	Cloudformation    = "cloudformation"
	ECRPublic         = "ecr-public"
)

var Services = map[string][]policy.Service{
	APIGateway:        policy.APIGatewayPolicies,
	Cloudfront:        policy.CloudfrontPolicies,
	Cloudtrail:        policy.CloudtrailPolicies,
	Codebuild:         policy.CodebuildPolicies,
	Codecommit:        policy.CodecommitPolicies,
	Dynamodb:          policy.DynamoDBPolicies,
	DynamodbStreams:   policy.DynamodbStreamsPolicies,
	EC2:               policy.EC2Policies,
	ECR:               policy.ECRPolicies,
	ECS:               policy.ECSPolicies,
	EKS:               policy.EKSPolicies,
	Elasticache:       policy.ElasticachePolicies,
	ElasticBeanStalk:  policy.ElasticBeanStalkPolicies,
	ElasticFileSystem: policy.ElasticFileSystemPolicies,
	IAM:               policy.IAMPolicies,
	Lambda:            policy.LambdaPolicies,
	S3:                policy.S3Policies,
	SQS:               policy.SQSPolicies,
	STS:               policy.STSPolicies,
	SNS:               policy.SNSPolicies,
	SSM:               policy.SSMPolicies,
	Route53:           policy.Route53Policies,
	RDS:               policy.RDSPolicies,
	ELB:               policy.ELBPolicies,
	Cloudformation:    policy.CloudformationPolicies,
	ECRPublic:         policy.ECRPublicPolicies,
}

func GetAWSServices() []string {
	keys := make([]string, 0, len(Services))
	for k := range Services {
		keys = append(keys, k)
	}
	return keys
}
