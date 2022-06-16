package cmd

import "github.com/securisec/cliam/aws"

var AWS_SERVICE_GROUPING = map[string][]string{
	"common": {
		aws.Lambda,
		aws.IAM,
		aws.S3,
		aws.APIGateway,
		aws.EC2,
	},
	"compute": {
		aws.Lambda,
		aws.EC2,
		aws.ElasticBeanStalk,
		aws.ElasticFileSystem,
		aws.ELB,
	},
	"databases": {
		aws.RDS,
		aws.Redshift,
		aws.Dynamodb,
	},
	"serverless": {
		aws.APIGateway,
		aws.Lambda,
		aws.S3,
		aws.SNS,
		aws.SQS,
		aws.Dynamodb,
	},
	"storage": {
		aws.S3,
		aws.Snowball,
		aws.StorageGateway,
		aws.ElasticFileSystem,
	},
}

func getAwsServiceGroups() []string {
	var groups []string
	for k := range AWS_SERVICE_GROUPING {
		groups = append(groups, k)
	}
	return groups
}
