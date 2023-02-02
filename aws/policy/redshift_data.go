package policy

import "github.com/securisec/cliam/shared"

// RedshiftDataPolicies policy
var RedshiftDataPolicies = map[string]Service{
	"ListStatements": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.ListStatements",
		},
		Permission: "ListStatements",
	},

	// extra
	"DescribeStatement": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.DescribeStatement",
		},
		Permission:             "DescribeStatement",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"DescribeTable": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.DescribeTable",
		},
		Permission:             "DescribeTable",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Database",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "database",
	},
	"GetStatementResult": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.GetStatementResult",
		},
		Permission:             "GetStatementResult",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Id",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "id",
	},
	"ListDatabases": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.ListDatabases",
		},
		Permission:             "ListDatabases",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Database",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "database",
	},
	"ListSchemas": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.ListSchemas",
		},
		Permission:             "ListSchemas",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Database",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "database",
	},
	"ListTables": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "RedshiftData.ListTables",
		},
		Permission:             "ListTables",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Database",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "database",
	},
}
