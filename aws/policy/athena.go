package policy

import "github.com/securisec/cliam/shared"

var AthenaPolicies = map[string]Service{
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
