package aws

import (
	"sort"

	"github.com/securisec/cliam/aws/policy"
)

const (
	AccessAnalyzers       = "access-analyzer"
	ACM                   = "acm"
	APIGateway            = "apigateway"
	AppFlow               = "appflow"
	AppRunner             = "apprunner"
	AppStream             = "appstream"
	AppSync               = "appsync"
	Athena                = "athena"
	AuditManager          = "auditmanager"
	AutoScaling           = "autoscaling"
	Backup                = "backup"
	BackupGateway         = "backup-gateway"
	Batch                 = "batch"
	Chime                 = "chime"
	CloudDirectory        = "clouddirectory"
	CloudFormation        = "cloudformation"
	CloudFront            = "cloudfront"
	CloudSearch           = "cloudsearch"
	CloudTrail            = "cloudtrail"
	CodeArtifact          = "codeartifact"
	CodeBuild             = "codebuild"
	CodeCommit            = "codecommit"
	CodeDeploy            = "codedeploy"
	CognitoIdentity       = "cognito-identity"
	CognitoIDP            = "cognito-idp"
	CognitoSync           = "cognito-sync"
	Comprehend            = "comprehend"
	ComputeOptimizer      = "compute-optimizer"
	ConfigService         = "config"
	DataBrew              = "databrew"
	DataExchange          = "dataexchange"
	DeviceFarm            = "devicefarm"
	DevopsGuru            = "devops-guru"
	DirectConnect         = "directconnect"
	DMS                   = "dms"
	DS                    = "ds"
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
	FSX                   = "fsx"
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
	IOT                   = "iot"
	IotWireless           = "api.iotwireless"
	Kafka                 = "kafka"
	KafkaConnect          = "kafkaconnect"
	Kinesis               = "kinesis"
	KinesisVideo          = "kinesisvideo"
	KMS                   = "kms"
	LakeFormation         = "lakeformation"
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
	Opensearch            = "opensearch"
	Organizations         = "organizations"
	Outposts              = "outposts"
	Panorama              = "panorama"
	Personalize           = "personalize"
	Polly                 = "polly"
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
	ServiceQuotas         = "servicequotas"
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
)

var Services = map[string]map[string]policy.Service{
	AccessAnalyzers:       policy.AccessAnalyzersPolicies,
	ACM:                   policy.ACMPolicies,
	APIGateway:            policy.APIGatewayPolicies,
	AppFlow:               policy.AppFlowPolicies,
	AppRunner:             policy.AppRunnerPolicies,
	AppStream:             policy.AppStreamPolicies,
	AppSync:               policy.AppSyncPolicies,
	Athena:                policy.AthenaPolicies,
	AuditManager:          policy.AuditManagerPolicies,
	AutoScaling:           policy.AutoScalingPolicies,
	Backup:                policy.BackupPolicies,
	BackupGateway:         policy.BackupGatewayPolicies,
	Batch:                 policy.BatchPolicies,
	Chime:                 policy.ChimePolicies,
	CloudDirectory:        policy.CloudDirectoryPolicies,
	CloudFormation:        policy.CloudFormationPolicies,
	CloudFront:            policy.CloudFrontPolicies,
	CloudSearch:           policy.CloudSearchPolicies,
	CloudTrail:            policy.CloudTrailPolicies,
	CodeArtifact:          policy.CodeArtifactPolicies,
	CodeBuild:             policy.CodeBuildPolicies,
	CodeCommit:            policy.CodeCommitPolicies,
	CodeDeploy:            policy.CodeDeployPolicies,
	CognitoIdentity:       policy.CognitoIdentityPolicies,
	CognitoIDP:            policy.CognitoIDPPolicies,
	CognitoSync:           policy.CognitoSyncPolicies,
	Comprehend:            policy.ComprehendPolicies,
	ComputeOptimizer:      policy.ComputeOptimizerPolicies,
	ConfigService:         policy.ConfigServicePolicies,
	DataBrew:              policy.DataBrewPolicies,
	DataExchange:          policy.DataExchangePolicies,
	DeviceFarm:            policy.DeviceFarmPolicies,
	DevopsGuru:            policy.DevopsGuruPolicies,
	DirectConnect:         policy.DirectConnectPolicies,
	DMS:                   policy.DMSPolicies,
	DS:                    policy.DSPolicies,
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
	FSX:                   policy.FSXPolicies,
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
	IOT:                   policy.IOTPolicies,
	IotWireless:           policy.IotWirelessPolicies,
	Kafka:                 policy.KafkaPolicies,
	KafkaConnect:          policy.KafkaConnectPolicies,
	Kinesis:               policy.KinesisPolicies,
	KinesisVideo:          policy.KinesisVideoPolicies,
	KMS:                   policy.KMSPolicies,
	LakeFormation:         policy.LakeFormationPolicies,
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
	Opensearch:            policy.OpensearchPolicies,
	Organizations:         policy.OrganizationsPolicies,
	Outposts:              policy.OutpostsPolicies,
	Panorama:              policy.PanoramaPolicies,
	Personalize:           policy.PersonalizePolicies,
	Polly:                 policy.PollyPolicies,
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
	ServiceQuotas:         policy.ServiceQuotasPolicies,
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
	sort.Strings(keys)
	return keys
}
