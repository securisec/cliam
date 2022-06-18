package policy

import "github.com/securisec/cliam/shared"

var MemoryDBPolicies = map[string]Service{
	"DescribeAcLs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeAcLs",
		},
		Permission: "DescribeAcLs",
	},
	"DescribeClusters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeClusters",
		},
		Permission: "DescribeClusters",
	},
	"DescribeEngineVersions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeEngineVersions",
		},
		Permission: "DescribeEngineVersions",
	},
	"DescribeParameterGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeParameterGroups",
		},
		Permission: "DescribeParameterGroups",
	},
	"DescribeServiceUpdates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeServiceUpdates",
		},
		Permission: "DescribeServiceUpdates",
	},
	"DescribeSnapshots": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeSnapshots",
		},
		Permission: "DescribeSnapshots",
	},
	"DescribeSubnetGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeSubnetGroups",
		},
		Permission: "DescribeSubnetGroups",
	},

	// extra
	"DescribeParameters": {
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
	"ListAllowedNodeTypeUpdates": {
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
	"ListTags": {
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
