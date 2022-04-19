package aws

import (
	"sort"

	"github.com/securisec/cliam/aws/policy"
)

const (
	APIGateway            = "apigateway"
	Cloudtrail            = "cloudtrail"
	Codebuild             = "codebuild"
	Codecommit            = "codecommit"
	Dynamodb              = "dynamodb"
	DynamodbStreams       = "streams.dynamodb"
	EC2                   = "ec2"
	ECR                   = "ecr"
	ECS                   = "ecs"
	EKS                   = "eks"
	Elasticache           = "elasticache"
	ElasticBeanStalk      = "elasticbeanstalk"
	ElasticFileSystem     = "elasticfilesystem"
	IAM                   = "iam"
	Lambda                = "lambda"
	S3                    = "s3"
	SQS                   = "sqs"
	STS                   = "sts"
	Cloudfront            = "cloudfront"
	SNS                   = "sns"
	SSM                   = "ssm"
	Route53               = "route53"
	RDS                   = "rds"
	ELB                   = "elasticloadbalancing"
	Cloudformation        = "cloudformation"
	ECRPublic             = "ecr-public"
	ES                    = "es"
	Cloudsearch           = "cloudsearch"
	SageMaker             = "sagemaker"
	SecretsManager        = "secretsmanager"
	StorageGateway        = "storagegateway"
	WAF                   = "waf"
	WAFRegional           = "waf-regional"
	Inspector             = "inspector"
	Email                 = "email"
	Events                = "events"
	DMS                   = "dms"
	Glue                  = "glue"
	Grafana               = "grafana"
	Autoscaling           = "autoscaling"
	Athena                = "athena"
	Lightsail             = "lightsail"
	Logs                  = "logs"
	Mediaconnect          = "mediaconnect"
	Appstream2            = "appstream2"
	AppRunner             = "apprunner"
	Workspaces            = "workspaces"
	ConfigService         = "config"
	KMS                   = "kms"
	Kinesis               = "kinesis"
	Kafka                 = "kafka"
	Route53Resolver       = "route53resolver"
	Shield                = "shield"
	Redshift              = "redshift"
	Directconnect         = "directconnect"
	DeviceFarm            = "devicefarm"
	SMS                   = "sms"
	VoiceID               = "voiceid"
	Xray                  = "xray"
	Personalize           = "personalize"
	Organizations         = "organizations"
	MemoryDB              = "memory-db"
	MediaStore            = "mediastore"
	KafkaConnect          = "kafkaconnect"
	Inspector2            = "inspector2"
	Macie2                = "macie2"
	SecurityHub           = "securityhub"
	Backup                = "backup"
	Outposts              = "outposts"
	NetworkFirewall       = "network-firewall"
	DevopsGuru            = "devops-guru"
	Forecase              = "forecast"
	Greengrass            = "greengrass"
	GroundStation         = "groundstation"
	GuardDuty             = "guardduty"
	ImageBuilder          = "imagebuilder"
	LicenseManager        = "license-manager"
	MediaConvert          = "mediaconvert"
	MediaLive             = "medialive"
	MigrationHubStratergy = "migrationhub-strategy"
	MQ                    = "mq"
	Panorama              = "panorama"
	Proton                = "proton"
	Robomaker             = "robomaker"
	ServiceCatalog        = "servicecatalog"
	Transfer              = "transfer"
	Transcribe            = "transcribe"
	Translate             = "translate"
	Wellarchitected       = "wellarchitected"
	WorkspacesWeb         = "workspaces-web"
	KinesisVideo          = "kinesisvideo"
	IotWireless           = "api.iotwireless"
	GameLift              = "gamelift"
	FraudDetector         = "frauddetector"
	Comprehend            = "comprehend"
	CodeDeploy            = "codedeploy"
	CloudDirectory        = "clouddirectory"
	Chime                 = "chime"
	WafV2                 = "wafv2"
	AccessAnalyzers       = "access-analyzer"
	ACM                   = "acm"
	AppFlow               = "appflow"
	AuditManager          = "auditmanager"
	DataBrew              = "databrew"
	Snowball              = "snowball"
)

var Services = map[string][]policy.Service{
	Outposts:              policy.OutpostsPolicies,
	Snowball:              policy.SnowballPolicies,
	WafV2:                 policy.WafV2Policies,
	AccessAnalyzers:       policy.AccessAnalyzersPolicies,
	ACM:                   policy.ACMPolicies,
	AppFlow:               policy.AppFlowPolicies,
	AuditManager:          policy.AuditManagerPolicies,
	DataBrew:              policy.DataBrewPolicies,
	GameLift:              policy.GameLiftPolicies,
	Chime:                 policy.ChimePolicies,
	CloudDirectory:        policy.CloudDirectoryPolicies,
	CodeDeploy:            policy.CodeDeployPolicies,
	IotWireless:           policy.IotWirelessPolicies,
	Comprehend:            policy.ComprehendPolicies,
	FraudDetector:         policy.FraudDetectorPolicies,
	MQ:                    policy.MQPolicies,
	KinesisVideo:          policy.KinesisVideoPolicies,
	Wellarchitected:       policy.WellarchitectedPolicies,
	WorkspacesWeb:         policy.WorkspacesWebPolicies,
	Proton:                policy.ProtonPolicies,
	Transfer:              policy.TransferPolicies,
	Translate:             policy.TranslatePolicies,
	Transcribe:            policy.TranscribePolicies,
	ServiceCatalog:        policy.ServiceCatalogPolicies,
	Robomaker:             policy.RobomakerPolicies,
	Panorama:              policy.PanoramaPolicies,
	MediaConvert:          policy.MediaConvertPolicies,
	MigrationHubStratergy: policy.MigrationHubStratergyPolicies,
	MediaLive:             policy.MediaLivePolicies,
	LicenseManager:        policy.LicenseManagerPolicies,
	NetworkFirewall:       policy.NetworkFirewallPolicies,
	ImageBuilder:          policy.ImageBuilderPolicies,
	GuardDuty:             policy.GuardDutyPolicies,
	GroundStation:         policy.GroundStationPolicies,
	Greengrass:            policy.GreengrassPolicies,
	Forecase:              policy.ForecastPolicies,
	DevopsGuru:            policy.DevopsGuruPolicies,
	APIGateway:            policy.APIGatewayPolicies,
	Cloudfront:            policy.CloudfrontPolicies,
	Cloudtrail:            policy.CloudtrailPolicies,
	Codebuild:             policy.CodebuildPolicies,
	Codecommit:            policy.CodecommitPolicies,
	Dynamodb:              policy.DynamoDBPolicies,
	DynamodbStreams:       policy.DynamodbStreamsPolicies,
	EC2:                   policy.EC2Policies,
	ECR:                   policy.ECRPolicies,
	ECS:                   policy.ECSPolicies,
	EKS:                   policy.EKSPolicies,
	Elasticache:           policy.ElasticachePolicies,
	ElasticBeanStalk:      policy.ElasticBeanStalkPolicies,
	ElasticFileSystem:     policy.ElasticFileSystemPolicies,
	IAM:                   policy.IAMPolicies,
	Lambda:                policy.LambdaPolicies,
	S3:                    policy.S3Policies,
	SQS:                   policy.SQSPolicies,
	STS:                   policy.STSPolicies,
	SNS:                   policy.SNSPolicies,
	SSM:                   policy.SSMPolicies,
	Route53:               policy.Route53Policies,
	RDS:                   policy.RDSPolicies,
	ELB:                   policy.ELBPolicies,
	Cloudformation:        policy.CloudformationPolicies,
	ECRPublic:             policy.ECRPublicPolicies,
	ES:                    policy.ESPolicies,
	Cloudsearch:           policy.CloudsearchPolicies,
	SageMaker:             policy.SageMakerPolicies,
	SecretsManager:        policy.SecretsManagerPolicies,
	StorageGateway:        policy.StorageGatewayPolicies,
	WAF:                   policy.WAFPolicies,
	WAFRegional:           policy.WAFRegionalPolicies,
	Inspector:             policy.InspectorPolicies,
	Email:                 policy.EmailPolicies,
	Events:                policy.EventsPolicies,
	DMS:                   policy.DMSPolicies,
	Glue:                  policy.GluePolicies,
	Grafana:               policy.GrafanaPolicies,
	Autoscaling:           policy.AutoscalingPolicies,
	Athena:                policy.AthenaPolicies,
	Lightsail:             policy.LightsailPolicies,
	Logs:                  policy.LogsPolicies,
	Mediaconnect:          policy.MediaconnectPolicies,
	Appstream2:            policy.Appstream2Policies,
	AppRunner:             policy.AppRunnerPolicies,
	Workspaces:            policy.WorkspacesPolicies,
	ConfigService:         policy.ConfigServicePolicies,
	KMS:                   policy.KMSPolicies,
	Kinesis:               policy.KinesisPolicies,
	Kafka:                 policy.KafkaPolicies,
	Route53Resolver:       policy.Route53ResolverPolicies,
	Shield:                policy.ShieldPolicies,
	Redshift:              policy.RedshiftPolicies,
	Directconnect:         policy.DirectconnectPolicies,
	DeviceFarm:            policy.DeviceFarmPolicies,
	SMS:                   policy.SMSPolicies,
	VoiceID:               policy.VoiceIdPolicies,
	Xray:                  policy.XrayPolicies,
	Personalize:           policy.PersonalizePolicies,
	Organizations:         policy.OrganizationsPolicies,
	MemoryDB:              policy.MemoryDBPolicies,
	MediaStore:            policy.MediaStorePolicies,
	KafkaConnect:          policy.KafkaConnectPolicies,
	Inspector2:            policy.Inspector2Policies,
	Macie2:                policy.Macie2Policies,
	SecurityHub:           policy.SecurityHubPolicies,
	Backup:                policy.BackupPolicies,
}

// GetAWSResources returns a list of AWS services
func GetAWSResources() []string {
	keys := make([]string, 0, len(Services))
	for k := range Services {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	return keys
}
