package policy

import "github.com/securisec/enumerate/shared"

var DynamoDBPolicies = []Service{
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListTables",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "ListTables",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListGlobalTables",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "ListGlobalTables",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListBackups",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "ListBackups",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeLimits",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "DescribeLimits",
	},
	{
		ServiceSuffix: "",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeEndpoints",
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
		},
		Permission: "DescribeEndpoints",
	},
}
