package policy

import "github.com/securisec/cliam/shared"

var DSPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeDirectories",
		},
		Permission: "DescribeDirectories",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeEventTopics",
		},
		Permission: "DescribeEventTopics",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.DescribeTrusts",
		},
		Permission: "DescribeTrusts",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.GetDirectoryLimits",
		},
		Permission: "GetDirectoryLimits",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "DirectoryService_20150416.ListLogSubscriptions",
		},
		Permission: "ListLogSubscriptions",
	},
}
