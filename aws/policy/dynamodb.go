package policy

import "github.com/securisec/cliam/shared"

var DynamoDBPolicies = []Service{
	// disabling this because it seems to be always valid
	// {
	// 	Method:   "POST",
	// 	JsonData: map[string]string{},
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
	// 		aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeEndpoints",
	// 	},
	// 	Permission: "DescribeEndpoints",
	// },
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeLimits",
		},
		Permission: "DescribeLimits",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListBackups",
		},
		Permission: "ListBackups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListContributorInsights",
		},
		Permission: "ListContributorInsights",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListExports",
		},
		Permission: "ListExports",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListGlobalTables",
		},
		Permission: "ListGlobalTables",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListTables",
		},
		Permission: "ListTables",
	},

	// extras
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeBackup",
		},
		Permission:             "DescribeBackup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "BackupArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "backup_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeContinuousBackups",
		},
		Permission:             "DescribeContinuousBackups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeContributorInsights",
		},
		Permission:             "DescribeContributorInsights",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeExport",
		},
		Permission:             "DescribeExport",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ExportArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "export_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeGlobalTable",
		},
		Permission:             "DescribeGlobalTable",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GlobalTableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "global_table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeGlobalTableSettings",
		},
		Permission:             "DescribeGlobalTableSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "GlobalTableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "global_table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeKinesisStreamingDestination",
		},
		Permission:             "DescribeKinesisStreamingDestination",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeTable",
		},
		Permission:             "DescribeTable",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeTableReplicaAutoScaling",
		},
		Permission:             "DescribeTableReplicaAutoScaling",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeTimeToLive",
		},
		Permission:             "DescribeTimeToLive",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TableName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "table_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListTagsOfResource",
		},
		Permission:             "ListTagsOfResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
