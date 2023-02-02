package policy

import "github.com/securisec/cliam/shared"

// DMSPolicies policy
var DMSPolicies = map[string]Service{
	"DescribeAccountAttributes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeAccountAttributes",
		},
		Permission: "DescribeAccountAttributes",
	},
	"DescribeApplicableIndividualAssessments": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeApplicableIndividualAssessments",
		},
		Permission: "DescribeApplicableIndividualAssessments",
	},
	"DescribeCertificates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeCertificates",
		},
		Permission: "DescribeCertificates",
	},
	"DescribeConnections": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeConnections",
		},
		Permission: "DescribeConnections",
	},
	"DescribeEndpointTypes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpointTypes",
		},
		Permission: "DescribeEndpointTypes",
	},
	"DescribeEndpoints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEndpoints",
		},
		Permission: "DescribeEndpoints",
	},
	"DescribeEventCategories": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEventCategories",
		},
		Permission: "DescribeEventCategories",
	},
	"DescribeEventSubscriptions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEventSubscriptions",
		},
		Permission: "DescribeEventSubscriptions",
	},
	"DescribeEvents": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeEvents",
		},
		Permission: "DescribeEvents",
	},
	"DescribeFleetAdvisorCollectors": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeFleetAdvisorCollectors",
		},
		Permission: "DescribeFleetAdvisorCollectors",
	},
	"DescribeFleetAdvisorDatabases": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeFleetAdvisorDatabases",
		},
		Permission: "DescribeFleetAdvisorDatabases",
	},
	"DescribeFleetAdvisorLsaAnalysis": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeFleetAdvisorLsaAnalysis",
		},
		Permission: "DescribeFleetAdvisorLsaAnalysis",
	},
	"DescribeFleetAdvisorSchemaObjectSummary": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeFleetAdvisorSchemaObjectSummary",
		},
		Permission: "DescribeFleetAdvisorSchemaObjectSummary",
	},
	"DescribeFleetAdvisorSchemas": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeFleetAdvisorSchemas",
		},
		Permission: "DescribeFleetAdvisorSchemas",
	},
	"DescribeOrderableReplicationInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeOrderableReplicationInstances",
		},
		Permission: "DescribeOrderableReplicationInstances",
	},
	"DescribePendingMaintenanceActions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribePendingMaintenanceActions",
		},
		Permission: "DescribePendingMaintenanceActions",
	},
	"DescribeReplicationInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationInstances",
		},
		Permission: "DescribeReplicationInstances",
	},
	"DescribeReplicationSubnetGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationSubnetGroups",
		},
		Permission: "DescribeReplicationSubnetGroups",
	},
	"DescribeReplicationTaskAssessmentResults": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskAssessmentResults",
		},
		Permission: "DescribeReplicationTaskAssessmentResults",
	},
	"DescribeReplicationTaskAssessmentRuns": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskAssessmentRuns",
		},
		Permission: "DescribeReplicationTaskAssessmentRuns",
	},
	"DescribeReplicationTaskIndividualAssessments": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTaskIndividualAssessments",
		},
		Permission: "DescribeReplicationTaskIndividualAssessments",
	},
	"DescribeReplicationTasks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.DescribeReplicationTasks",
		},
		Permission: "DescribeReplicationTasks",
	},
	"ListTagsForResource": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.ListTagsForResource",
		},
		Permission: "ListTagsForResource",
	},
	"RunFleetAdvisorLsaAnalysis": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.RunFleetAdvisorLsaAnalysis",
		},
		Permission: "RunFleetAdvisorLsaAnalysis",
	},
	"UpdateSubscriptionsToEventBridge": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonDMSv20160101.UpdateSubscriptionsToEventBridge",
		},
		Permission: "UpdateSubscriptionsToEventBridge",
	},

	// extra
	"DescribeEndpointSettings": {
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
	"DescribeRefreshSchemasStatus": {
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
	"DescribeReplicationInstanceTaskLogs": {
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
	"DescribeSchemas": {
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
	"DescribeTableStatistics": {
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
