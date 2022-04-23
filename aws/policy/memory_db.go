package policy

import "github.com/securisec/cliam/shared"

var MemoryDBPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeAcLs",
		},
		Permission: "DescribeAcLs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeClusters",
		},
		Permission: "DescribeClusters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeEngineVersions",
		},
		Permission: "DescribeEngineVersions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeParameterGroups",
		},
		Permission: "DescribeParameterGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeServiceUpdates",
		},
		Permission: "DescribeServiceUpdates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeSubnetGroups",
		},
		Permission: "DescribeSubnetGroups",
	},
}
