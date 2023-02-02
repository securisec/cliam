package policy

import "github.com/securisec/cliam/shared"

// MemoryDBPolicies policy
var MemoryDBPolicies = map[string]Service{
	"DescribeACLs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeACLs",
		},
		Permission: "DescribeACLs",
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
	"DescribeEvents": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeEvents",
		},
		Permission: "DescribeEvents",
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
	"DescribeReservedNodes": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeReservedNodes",
		},
		Permission: "DescribeReservedNodes",
	},
	"DescribeReservedNodesOfferings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeReservedNodesOfferings",
		},
		Permission: "DescribeReservedNodesOfferings",
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
	"DescribeUsers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonMemoryDB.DescribeUsers",
		},
		Permission: "DescribeUsers",
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
