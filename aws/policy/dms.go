package policy

import "github.com/securisec/cliam/shared"

var DMSPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeAccountAttributes",
		},
		Permission: "DescribeAccountAttributes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeApplicableIndividualAssessments",
		},
		Permission: "DescribeApplicableIndividualAssessments",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeCertificates",
		},
		Permission: "DescribeCertificates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeConnections",
		},
		Permission: "DescribeConnections",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpointTypes",
		},
		Permission: "DescribeEndpointTypes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpoints",
		},
		Permission: "DescribeEndpoints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEventCategories",
		},
		Permission: "DescribeEventCategories",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEventSubscriptions",
		},
		Permission: "DescribeEventSubscriptions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeOrderableReplicationInstances",
		},
		Permission: "DescribeOrderableReplicationInstances",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribePendingMaintenanceActions",
		},
		Permission: "DescribePendingMaintenanceActions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationInstances",
		},
		Permission: "DescribeReplicationInstances",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationSubnetGroups",
		},
		Permission: "DescribeReplicationSubnetGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskAssessmentResults",
		},
		Permission: "DescribeReplicationTaskAssessmentResults",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskAssessmentRuns",
		},
		Permission: "DescribeReplicationTaskAssessmentRuns",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskIndividualAssessments",
		},
		Permission: "DescribeReplicationTaskIndividualAssessments",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTasks",
		},
		Permission: "DescribeReplicationTasks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.ListTagsForResource",
		},
		Permission: "ListTagsForResource",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpointSettings",
		},
		Permission:             "DescribeEndpointSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EngineName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "engine_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeRefreshSchemasStatus",
		},
		Permission:             "DescribeRefreshSchemasStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationInstanceTaskLogs",
		},
		Permission:             "DescribeReplicationInstanceTaskLogs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ReplicationInstanceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "replication_instance_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeSchemas",
		},
		Permission:             "DescribeSchemas",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeTableStatistics",
		},
		Permission:             "DescribeTableStatistics",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ReplicationTaskArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "replication_task_arn",
	},
}
