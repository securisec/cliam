package policy

import "github.com/securisec/cliam/shared"

// AthenaPolicies policy
var AthenaPolicies = map[string]Service{
	"ListApplicationDPUSizes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListApplicationDPUSizes",
		},
		Permission: "ListApplicationDPUSizes",
	},
	"ListDataCatalogs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListDataCatalogs",
		},
		Permission: "ListDataCatalogs",
	},
	"ListEngineVersions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListEngineVersions",
		},
		Permission: "ListEngineVersions",
	},
	"ListNamedQueries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListNamedQueries",
		},
		Permission: "ListNamedQueries",
	},
	"ListQueryExecutions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListQueryExecutions",
		},
		Permission: "ListQueryExecutions",
	},
	"ListWorkGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListWorkGroups",
		},
		Permission: "ListWorkGroups",
	},

	// extra
	"GetCalculationExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetCalculationExecution",
		},
		Permission:             "GetCalculationExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CalculationExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "calculation_execution_id",
	},
	"GetCalculationExecutionCode": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetCalculationExecutionCode",
		},
		Permission:             "GetCalculationExecutionCode",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CalculationExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "calculation_execution_id",
	},
	"GetCalculationExecutionStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetCalculationExecutionStatus",
		},
		Permission:             "GetCalculationExecutionStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CalculationExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "calculation_execution_id",
	},
	"GetDataCatalog": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetDataCatalog",
		},
		Permission:             "GetDataCatalog",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetNamedQuery": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetNamedQuery",
		},
		Permission:             "GetNamedQuery",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NamedQueryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "named_query_id",
	},
	"GetNotebookMetadata": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetNotebookMetadata",
		},
		Permission:             "GetNotebookMetadata",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NotebookId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "notebook_id",
	},
	"GetQueryExecution": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetQueryExecution",
		},
		Permission:             "GetQueryExecution",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_execution_id",
	},
	"GetQueryResults": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetQueryResults",
		},
		Permission:             "GetQueryResults",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_execution_id",
	},
	"GetQueryRuntimeStatistics": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetQueryRuntimeStatistics",
		},
		Permission:             "GetQueryRuntimeStatistics",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryExecutionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_execution_id",
	},
	"GetSession": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetSession",
		},
		Permission:             "GetSession",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SessionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_id",
	},
	"GetSessionStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetSessionStatus",
		},
		Permission:             "GetSessionStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SessionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_id",
	},
	"GetWorkGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.GetWorkGroup",
		},
		Permission:             "GetWorkGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "work_group",
	},
	"ListCalculationExecutions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListCalculationExecutions",
		},
		Permission:             "ListCalculationExecutions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SessionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_id",
	},
	"ListDatabases": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListDatabases",
		},
		Permission:             "ListDatabases",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CatalogName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "catalog_name",
	},
	"ListExecutors": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListExecutors",
		},
		Permission:             "ListExecutors",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SessionId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_id",
	},
	"ListNotebookMetadata": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListNotebookMetadata",
		},
		Permission:             "ListNotebookMetadata",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "work_group",
	},
	"ListNotebookSessions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListNotebookSessions",
		},
		Permission:             "ListNotebookSessions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NotebookId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "notebook_id",
	},
	"ListPreparedStatements": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListPreparedStatements",
		},
		Permission:             "ListPreparedStatements",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "work_group",
	},
	"ListSessions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListSessions",
		},
		Permission:             "ListSessions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "WorkGroup",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "work_group",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonAthena.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
