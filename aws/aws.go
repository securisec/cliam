package aws

import "github.com/securisec/cliam/aws/policy"

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
	ES                = "es"
	Cloudsearch       = "cloudsearch"
	SageMaker         = "sagemaker"
	SecretsManager    = "secretsmanager"
	StorageGateway    = "storagegateway"
	WAF               = "waf"
	WAFRegional       = "waf-regional"
	Inspector         = "inspector"
	Email             = "email"
	Events            = "events"
	DMS               = "dms"
	Glue              = "glue"
	Grafana           = "grafana"
	Autoscaling       = "autoscaling"
	Athena            = "athena"
	Lightsail         = "lightsail"
	Logs              = "logs"
	Mediaconnect      = "mediaconnect"
	Appstream2        = "appstream2"
	AppRunner         = "apprunner"
	Workspaces        = "workspaces"
	ConfigService     = "config"
	KMS               = "kms"
	Kinesis           = "kinesis"
	Kafka             = "kafka"
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
	ES:                policy.ESPolicies,
	Cloudsearch:       policy.CloudsearchPolicies,
	SageMaker:         policy.SageMakerPolicies,
	SecretsManager:    policy.SecretsManagerPolicies,
	StorageGateway:    policy.StorageGatewayPolicies,
	WAF:               policy.WAFPolicies,
	WAFRegional:       policy.WAFRegionalPolicies,
	Inspector:         policy.InspectorPolicies,
	Email:             policy.EmailPolicies,
	Events:            policy.EventsPolicies,
	DMS:               policy.DMSPolicies,
	Glue:              policy.GluePolicies,
	Grafana:           policy.GrafanaPolicies,
	Autoscaling:       policy.AutoscalingPolicies,
	Athena:            policy.AthenaPolicies,
	Lightsail:         policy.LightsailPolicies,
	Logs:              policy.LogsPolicies,
	Mediaconnect:      policy.MediaconnectPolicies,
	Appstream2:        policy.Appstream2Policies,
	AppRunner:         policy.AppRunnerPolicies,
	Workspaces:        policy.WorkspacesPolicies,
	ConfigService:     policy.ConfigServicePolicies,
	KMS:               policy.KMSPolicies,
	Kinesis:           policy.KinesisPolicies,
	Kafka:             policy.KafkaPolicies,
}

// GetAWSResources returns a list of AWS services
func GetAWSResources() []string {
	keys := make([]string, 0, len(Services))
	for k := range Services {
		keys = append(keys, k)
	}
	return keys
}
