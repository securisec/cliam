package policy

import "github.com/securisec/cliam/shared"

var GluePolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCatalogImportStatus",
		},
		Permission: "GetCatalogImportStatus",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetClassifiers",
		},
		Permission: "GetClassifiers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetConnections",
		},
		Permission: "GetConnections",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCrawlerMetrics",
		},
		Permission: "GetCrawlerMetrics",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetCrawlers",
		},
		Permission: "GetCrawlers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDataCatalogEncryptionSettings",
		},
		Permission: "GetDataCatalogEncryptionSettings",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDatabases",
		},
		Permission: "GetDatabases",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDataflowGraph",
		},
		Permission: "GetDataflowGraph",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetDevEndpoints",
		},
		Permission: "GetDevEndpoints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetJobs",
		},
		Permission: "GetJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetMlTransforms",
		},
		Permission: "GetMlTransforms",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetResourcePolicies",
		},
		Permission: "GetResourcePolicies",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetResourcePolicy",
		},
		Permission: "GetResourcePolicy",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSchemaVersion",
		},
		Permission: "GetSchemaVersion",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetSecurityConfigurations",
		},
		Permission: "GetSecurityConfigurations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.GetTriggers",
		},
		Permission: "GetTriggers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListBlueprints",
		},
		Permission: "ListBlueprints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListCrawlers",
		},
		Permission: "ListCrawlers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListDevEndpoints",
		},
		Permission: "ListDevEndpoints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListJobs",
		},
		Permission: "ListJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListMlTransforms",
		},
		Permission: "ListMlTransforms",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListRegistries",
		},
		Permission: "ListRegistries",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListSchemas",
		},
		Permission: "ListSchemas",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListSessions",
		},
		Permission: "ListSessions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListTriggers",
		},
		Permission: "ListTriggers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSGlue.ListWorkflows",
		},
		Permission: "ListWorkflows",
	},

	// extra
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
	{
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
