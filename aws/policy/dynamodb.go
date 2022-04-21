package policy

import "github.com/securisec/cliam/shared"

var DynamoDBPolicies = []Service{
	// disabling this because it seems to be always valid
	// {
	// 	Method:   "POST",
	// 	JsonData: `{}`,
	// 	Headers: map[string]string{
	// 		shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
	// 		aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeEndpoints",
	// 	},
	// 	Permission: "DescribeEndpoints",
	// },
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.DescribeLimits",
		},
		Permission: "DescribeLimits",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListBackups",
		},
		Permission: "ListBackups",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListContributorInsights",
		},
		Permission: "ListContributorInsights",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListExports",
		},
		Permission: "ListExports",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListGlobalTables",
		},
		Permission: "ListGlobalTables",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_JSON,
			aws_X_AMZ_TARGET:           "DynamoDB_20120810.ListTables",
		},
		Permission: "ListTables",
	},
}
