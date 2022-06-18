package aws

import (
	"sort"

	"github.com/securisec/cliam/aws/policy"
)

const (
	LakeFormation         = "lakeformation"
	AccessAnalyzers       = "access-analyzer"
	ACM                   = "acm"
	APIGateway            = "apigateway"
	AppFlow               = "appflow"
	AppRunner             = "apprunner"
	Appstream2            = "appstream2"
	Athena                = "athena"
	AuditManager          = "auditmanager"
	Autoscaling           = "autoscaling"
	Backup                = "backup"
	Chime                 = "chime"
	CloudDirectory        = "clouddirectory"
	Cloudformation        = "cloudformation"
	Cloudfront            = "cloudfront"
	Cloudsearch           = "cloudsearch"
	Cloudtrail            = "cloudtrail"
	Codebuild             = "codebuild"
	Codecommit            = "codecommit"
	CodeDeploy            = "codedeploy"
	Comprehend            = "comprehend"
	ConfigService         = "config"
	DataBrew              = "databrew"
	DeviceFarm            = "devicefarm"
	DevopsGuru            = "devops-guru"
	Directconnect         = "directconnect"
	DMS                   = "dms"
	Dynamodb              = "dynamodb"
	DynamodbStreams       = "streams.dynamodb"
	EC2                   = "ec2"
	ECR                   = "ecr"
	ECRPublic             = "ecr-public"
	ECS                   = "ecs"
	EKS                   = "eks"
	Elasticache           = "elasticache"
	ElasticBeanStalk      = "elasticbeanstalk"
	ElasticFileSystem     = "elasticfilesystem"
	ELB                   = "elasticloadbalancing"
	Email                 = "email"
	ES                    = "es"
	Events                = "events"
	Forecase              = "forecast"
	FraudDetector         = "frauddetector"
	GameLift              = "gamelift"
	Glue                  = "glue"
	Grafana               = "grafana"
	Greengrass            = "greengrass"
	GroundStation         = "groundstation"
	GuardDuty             = "guardduty"
	IAM                   = "iam"
	ImageBuilder          = "imagebuilder"
	Inspector             = "inspector"
	Inspector2            = "inspector2"
	IotWireless           = "api.iotwireless"
	Kafka                 = "kafka"
	KafkaConnect          = "kafkaconnect"
	Kinesis               = "kinesis"
	KinesisVideo          = "kinesisvideo"
	KMS                   = "kms"
	Lambda                = "lambda"
	LicenseManager        = "license-manager"
	Lightsail             = "lightsail"
	Logs                  = "logs"
	Macie2                = "macie2"
	Mediaconnect          = "mediaconnect"
	MediaConvert          = "mediaconvert"
	MediaLive             = "medialive"
	MediaStore            = "mediastore"
	MemoryDB              = "memory-db"
	MigrationHubStratergy = "migrationhub-strategy"
	MQ                    = "mq"
	NetworkFirewall       = "network-firewall"
	Organizations         = "organizations"
	Outposts              = "outposts"
	Panorama              = "panorama"
	Personalize           = "personalize"
	Proton                = "proton"
	RDS                   = "rds"
	Redshift              = "redshift"
	Robomaker             = "robomaker"
	Route53               = "route53"
	Route53Resolver       = "route53resolver"
	S3                    = "s3"
	SageMaker             = "sagemaker"
	SecretsManager        = "secretsmanager"
	SecurityHub           = "securityhub"
	ServiceCatalog        = "servicecatalog"
	Shield                = "shield"
	SMS                   = "sms"
	Snowball              = "snowball"
	SNS                   = "sns"
	SQS                   = "sqs"
	SSM                   = "ssm"
	StorageGateway        = "storagegateway"
	STS                   = "sts"
	Transcribe            = "transcribe"
	Transfer              = "transfer"
	Translate             = "translate"
	VoiceID               = "voiceid"
	WAF                   = "waf"
	WAFRegional           = "waf-regional"
	WafV2                 = "wafv2"
	Wellarchitected       = "wellarchitected"
	Workspaces            = "workspaces"
	WorkspacesWeb         = "workspaces-web"
	Xray                  = "xray"
	ServiceQuotas         = "servicequotas"
	IOT                   = "iot"
	BackupGateway         = "backup-gateway"
	ComputeOptimizer      = "compute-optimizer"
	Batch                 = "batch"
	DataExchange          = "dataexchange"
	DS                    = "ds"
	FSX                   = "fsx"
)

var Services = map[string]map[string]policy.Service{
	AccessAnalyzers:       policy.AccessAnalyzersPolicies,
	FSX:                   policy.FSXPolicies,
	DataExchange:          policy.DataExchangePolicies,
	DS:                    policy.DSPolicies,
	ComputeOptimizer:      policy.ComputeOptimizerPolicies,
	IOT:                   policy.IOTPolicies,
	BackupGateway:         policy.BackupGatewayPolicies,
	ServiceQuotas:         policy.ServiceQuotasPolicies,
	Batch:                 policy.BatchPolicies,
	LakeFormation:         policy.LakeFormationPolicies,
	ACM:                   policy.ACMPolicies,
	APIGateway:            policy.APIGatewayPolicies,
	AppFlow:               policy.AppFlowPolicies,
	AppRunner:             policy.AppRunnerPolicies,
	Appstream2:            policy.Appstream2Policies,
	Athena:                policy.AthenaPolicies,
	AuditManager:          policy.AuditManagerPolicies,
	Autoscaling:           policy.AutoscalingPolicies,
	Backup:                policy.BackupPolicies,
	Chime:                 policy.ChimePolicies,
	CloudDirectory:        policy.CloudDirectoryPolicies,
	Cloudformation:        policy.CloudformationPolicies,
	Cloudfront:            policy.CloudfrontPolicies,
	Cloudsearch:           policy.CloudsearchPolicies,
	Cloudtrail:            policy.CloudtrailPolicies,
	Codebuild:             policy.CodebuildPolicies,
	Codecommit:            policy.CodecommitPolicies,
	CodeDeploy:            policy.CodeDeployPolicies,
	Comprehend:            policy.ComprehendPolicies,
	ConfigService:         policy.ConfigServicePolicies,
	DataBrew:              policy.DataBrewPolicies,
	DeviceFarm:            policy.DeviceFarmPolicies,
	DevopsGuru:            policy.DevopsGuruPolicies,
	Directconnect:         policy.DirectconnectPolicies,
	DMS:                   policy.DMSPolicies,
	Dynamodb:              policy.DynamoDBPolicies,
	DynamodbStreams:       policy.DynamodbStreamsPolicies,
	EC2:                   policy.EC2Policies,
	ECR:                   policy.ECRPolicies,
	ECRPublic:             policy.ECRPublicPolicies,
	ECS:                   policy.ECSPolicies,
	EKS:                   policy.EKSPolicies,
	Elasticache:           policy.ElasticachePolicies,
	ElasticBeanStalk:      policy.ElasticBeanStalkPolicies,
	ElasticFileSystem:     policy.ElasticFileSystemPolicies,
	ELB:                   policy.ELBPolicies,
	Email:                 policy.EmailPolicies,
	ES:                    policy.ESPolicies,
	Events:                policy.EventsPolicies,
	Forecase:              policy.ForecastPolicies,
	FraudDetector:         policy.FraudDetectorPolicies,
	GameLift:              policy.GameLiftPolicies,
	Glue:                  policy.GluePolicies,
	Grafana:               policy.GrafanaPolicies,
	Greengrass:            policy.GreengrassPolicies,
	GroundStation:         policy.GroundStationPolicies,
	GuardDuty:             policy.GuardDutyPolicies,
	IAM:                   policy.IAMPolicies,
	ImageBuilder:          policy.ImageBuilderPolicies,
	Inspector:             policy.InspectorPolicies,
	Inspector2:            policy.Inspector2Policies,
	IotWireless:           policy.IotWirelessPolicies,
	Kafka:                 policy.KafkaPolicies,
	KafkaConnect:          policy.KafkaConnectPolicies,
	Kinesis:               policy.KinesisPolicies,
	KinesisVideo:          policy.KinesisVideoPolicies,
	KMS:                   policy.KMSPolicies,
	Lambda:                policy.LambdaPolicies,
	LicenseManager:        policy.LicenseManagerPolicies,
	Lightsail:             policy.LightsailPolicies,
	Logs:                  policy.LogsPolicies,
	Macie2:                policy.Macie2Policies,
	Mediaconnect:          policy.MediaconnectPolicies,
	MediaConvert:          policy.MediaConvertPolicies,
	MediaLive:             policy.MediaLivePolicies,
	MediaStore:            policy.MediaStorePolicies,
	MemoryDB:              policy.MemoryDBPolicies,
	MigrationHubStratergy: policy.MigrationHubStratergyPolicies,
	MQ:                    policy.MQPolicies,
	NetworkFirewall:       policy.NetworkFirewallPolicies,
	Organizations:         policy.OrganizationsPolicies,
	Outposts:              policy.OutpostsPolicies,
	Panorama:              policy.PanoramaPolicies,
	Personalize:           policy.PersonalizePolicies,
	Proton:                policy.ProtonPolicies,
	RDS:                   policy.RDSPolicies,
	Redshift:              policy.RedshiftPolicies,
	Robomaker:             policy.RobomakerPolicies,
	Route53:               policy.Route53Policies,
	Route53Resolver:       policy.Route53ResolverPolicies,
	S3:                    policy.S3Policies,
	SageMaker:             policy.SageMakerPolicies,
	SecretsManager:        policy.SecretsManagerPolicies,
	SecurityHub:           policy.SecurityHubPolicies,
	ServiceCatalog:        policy.ServiceCatalogPolicies,
	Shield:                policy.ShieldPolicies,
	SMS:                   policy.SMSPolicies,
	Snowball:              policy.SnowballPolicies,
	SNS:                   policy.SNSPolicies,
	SQS:                   policy.SQSPolicies,
	SSM:                   policy.SSMPolicies,
	StorageGateway:        policy.StorageGatewayPolicies,
	STS:                   policy.STSPolicies,
	Transcribe:            policy.TranscribePolicies,
	Transfer:              policy.TransferPolicies,
	Translate:             policy.TranslatePolicies,
	VoiceID:               policy.VoiceIdPolicies,
	WAF:                   policy.WAFPolicies,
	WAFRegional:           policy.WAFRegionalPolicies,
	WafV2:                 policy.WafV2Policies,
	Wellarchitected:       policy.WellarchitectedPolicies,
	Workspaces:            policy.WorkspacesPolicies,
	WorkspacesWeb:         policy.WorkspacesWebPolicies,
	Xray:                  policy.XrayPolicies,
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
