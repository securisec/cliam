package policy

import "github.com/securisec/cliam/shared"

var ElasticachePolicies = map[string]Service{
	"DescribeCacheClusters": {
		ServiceSuffix: "?Action=DescribeCacheClusters&Version=2015-02-02",
		Permission:    "DescribeCacheClusters",
	},
	"DescribeCacheEngineVersions": {
		ServiceSuffix: "?Action=DescribeCacheEngineVersions&Version=2015-02-02",
		Permission:    "DescribeCacheEngineVersions",
	},
	"DescribeCacheParameterGroups": {
		ServiceSuffix: "?Action=DescribeCacheParameterGroups&Version=2015-02-02",
		Permission:    "DescribeCacheParameterGroups",
	},
	"DescribeCacheSecurityGroups": {
		ServiceSuffix: "?Action=DescribeCacheSecurityGroups",
		Permission:    "DescribeCacheSecurityGroups",
	},
	"DescribeCacheSubnetGroups": {
		ServiceSuffix: "?Action=DescribeCacheSubnetGroups&Version=2015-02-02",
		Permission:    "DescribeCacheSubnetGroups",
	},
	"DescribeReplicationGroups": {
		ServiceSuffix: "?Action=DescribeReplicationGroups&Version=2015-02-02",
		Permission:    "DescribeReplicationGroups",
	},
	"DescribeReservedCacheNodes": {
		ServiceSuffix: "?Action=DescribeReservedCacheNodes&Version=2015-02-02",
		Permission:    "DescribeReservedCacheNodes",
	},
	"DescribeReservedCacheNodesOfferings": {
		ServiceSuffix: "?Action=DescribeReservedCacheNodesOfferings&Version=2015-02-02",
		Permission:    "DescribeReservedCacheNodesOfferings",
	},
	"DescribeSnapshots": {
		ServiceSuffix: "?Action=DescribeSnapshots&Version=2015-02-02",
		Permission:    "DescribeSnapshots",
	},
	"ListAllowedNodeTypeModifications": {
		ServiceSuffix: "?Action=ListAllowedNodeTypeModifications&Version=2015-02-02",
		Permission:    "ListAllowedNodeTypeModifications",
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
