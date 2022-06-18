package policy

import "github.com/securisec/cliam/shared"

var GluePolicies = map[string]Service{
	"GetCatalogImportStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCatalogImportStatus",
		},
		Permission: "GetCatalogImportStatus",
	},
	"GetClassifiers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetClassifiers",
		},
		Permission: "GetClassifiers",
	},
	"GetConnections": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetConnections",
		},
		Permission: "GetConnections",
	},
	"GetCrawlerMetrics": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCrawlerMetrics",
		},
		Permission: "GetCrawlerMetrics",
	},
	"GetCrawlers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCrawlers",
		},
		Permission: "GetCrawlers",
	},
	"GetDataCatalogEncryptionSettings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDataCatalogEncryptionSettings",
		},
		Permission: "GetDataCatalogEncryptionSettings",
	},
	"GetDatabases": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDatabases",
		},
		Permission: "GetDatabases",
	},
	"GetDataflowGraph": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDataflowGraph",
		},
		Permission: "GetDataflowGraph",
	},
	"GetDevEndpoints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDevEndpoints",
		},
		Permission: "GetDevEndpoints",
	},
	"GetJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetJobs",
		},
		Permission: "GetJobs",
	},
	"GetMlTransforms": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetMlTransforms",
		},
		Permission: "GetMlTransforms",
	},
	"GetResourcePolicies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetResourcePolicies",
		},
		Permission: "GetResourcePolicies",
	},
	"GetResourcePolicy": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetResourcePolicy",
		},
		Permission: "GetResourcePolicy",
	},
	"GetSchemaVersion": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSchemaVersion",
		},
		Permission: "GetSchemaVersion",
	},
	"GetSecurityConfigurations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSecurityConfigurations",
		},
		Permission: "GetSecurityConfigurations",
	},
	"GetTriggers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetTriggers",
		},
		Permission: "GetTriggers",
	},
	"ListBlueprints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListBlueprints",
		},
		Permission: "ListBlueprints",
	},
	"ListCrawlers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListCrawlers",
		},
		Permission: "ListCrawlers",
	},
	"ListDevEndpoints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListDevEndpoints",
		},
		Permission: "ListDevEndpoints",
	},
	"ListJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListJobs",
		},
		Permission: "ListJobs",
	},
	"ListMlTransforms": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListMlTransforms",
		},
		Permission: "ListMlTransforms",
	},
	"ListRegistries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListRegistries",
		},
		Permission: "ListRegistries",
	},
	"ListSchemas": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListSchemas",
		},
		Permission: "ListSchemas",
	},
	"ListSessions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListSessions",
		},
		Permission: "ListSessions",
	},
	"ListTriggers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListTriggers",
		},
		Permission: "ListTriggers",
	},
	"ListWorkflows": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListWorkflows",
		},
		Permission: "ListWorkflows",
	},

	// extra
	"GetBlueprint": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetBlueprint",
		},
		Permission:             "GetBlueprint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetBlueprintRuns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetBlueprintRuns",
		},
		Permission:             "GetBlueprintRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "BlueprintName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "blueprint_name",
	},
	"GetClassifier": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetClassifier",
		},
		Permission:             "GetClassifier",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetConnection": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetConnection",
		},
		Permission:             "GetConnection",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetCrawler": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCrawler",
		},
		Permission:             "GetCrawler",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetDatabase": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDatabase",
		},
		Permission:             "GetDatabase",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetDevEndpoint": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDevEndpoint",
		},
		Permission:             "GetDevEndpoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_name",
	},
	"GetJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetJob",
		},
		Permission:             "GetJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_name",
	},
	"GetJobBookmark": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetJobBookmark",
		},
		Permission:             "GetJobBookmark",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_name",
	},
	"GetJobRuns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetJobRuns",
		},
		Permission:             "GetJobRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_name",
	},
	"GetMLTaskRuns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetMLTaskRuns",
		},
		Permission:             "GetMLTaskRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransformId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "transform_id",
	},
	"GetMLTransform": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetMLTransform",
		},
		Permission:             "GetMLTransform",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TransformId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "transform_id",
	},
	"GetMapping": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetMapping",
		},
		Permission:             "GetMapping",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Source",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "source",
	},
	"GetRegistry": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetRegistry",
		},
		Permission:             "GetRegistry",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RegistryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "registry_id",
	},
	"GetSchema": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSchema",
		},
		Permission:             "GetSchema",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SchemaId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "schema_id",
	},
	"GetSecurityConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSecurityConfiguration",
		},
		Permission:             "GetSecurityConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetSession": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSession",
		},
		Permission:             "GetSession",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"GetTables": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetTables",
		},
		Permission:             "GetTables",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DatabaseName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "database_name",
	},
	"GetTags": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetTags",
		},
		Permission:             "GetTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetTrigger": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetTrigger",
		},
		Permission:             "GetTrigger",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetUserDefinedFunctions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetUserDefinedFunctions",
		},
		Permission:             "GetUserDefinedFunctions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Pattern",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "pattern",
	},
	"GetWorkflow": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetWorkflow",
		},
		Permission:             "GetWorkflow",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetWorkflowRuns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetWorkflowRuns",
		},
		Permission:             "GetWorkflowRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListSchemaVersions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListSchemaVersions",
		},
		Permission:             "ListSchemaVersions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SchemaId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "schema_id",
	},
	"ListStatements": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListStatements",
		},
		Permission:             "ListStatements",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SessionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_id",
	},
}
