package policy

import "github.com/securisec/cliam/shared"

// ElasticachePolicies policy
var ElasticachePolicies = map[string]Service{
	"DescribeCacheClusters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCacheClusters",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCacheClusters",
	},
	"DescribeCacheEngineVersions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCacheEngineVersions",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCacheEngineVersions",
	},
	"DescribeCacheParameterGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCacheParameterGroups",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCacheParameterGroups",
	},
	"DescribeCacheSecurityGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCacheSecurityGroups",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCacheSecurityGroups",
	},
	"DescribeCacheSubnetGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCacheSubnetGroups",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeCacheSubnetGroups",
	},
	"DescribeEvents": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEvents",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeEvents",
	},
	"DescribeGlobalReplicationGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeGlobalReplicationGroups",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeGlobalReplicationGroups",
	},
	"DescribeReplicationGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReplicationGroups",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReplicationGroups",
	},
	"DescribeReservedCacheNodes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedCacheNodes",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedCacheNodes",
	},
	"DescribeReservedCacheNodesOfferings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReservedCacheNodesOfferings",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeReservedCacheNodesOfferings",
	},
	"DescribeServiceUpdates": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeServiceUpdates",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeServiceUpdates",
	},
	"DescribeSnapshots": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSnapshots",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeSnapshots",
	},
	"DescribeUpdateActions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeUpdateActions",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeUpdateActions",
	},
	"DescribeUserGroups": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeUserGroups",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeUserGroups",
	},
	"DescribeUsers": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeUsers",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeUsers",
	},
	"ListAllowedNodeTypeModifications": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListAllowedNodeTypeModifications",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListAllowedNodeTypeModifications",
	},

	// extra
	"DescribeCacheParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeCacheParameters",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeCacheParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CacheParameterGroupName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "cache_parameter_group_name",
	},
	"DescribeEngineDefaultParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeEngineDefaultParameters",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeEngineDefaultParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CacheParameterGroupFamily",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "cache_parameter_group_family",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2015-02-02",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_name",
	},
}
