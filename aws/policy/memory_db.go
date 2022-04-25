package policy

import "github.com/securisec/cliam/shared"

var MemoryDBPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeAcLs",
		},
		Permission: "DescribeAcLs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeClusters",
		},
		Permission: "DescribeClusters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeEngineVersions",
		},
		Permission: "DescribeEngineVersions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeParameterGroups",
		},
		Permission: "DescribeParameterGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeServiceUpdates",
		},
		Permission: "DescribeServiceUpdates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeSubnetGroups",
		},
		Permission: "DescribeSubnetGroups",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeParameters",
		},
		Permission:             "DescribeParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ParameterGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "parameter_group_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.ListAllowedNodeTypeUpdates",
		},
		Permission:             "ListAllowedNodeTypeUpdates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ClusterName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "cluster_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.ListTags",
		},
		Permission:             "ListTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
