package policy

import "github.com/securisec/cliam/shared"

var ElasticachePolicies = []Service{
	{
		ServiceSuffix: "?Action=DescribeCacheClusters&Version=2015-02-02",
		Permission:    "DescribeCacheClusters",
	},
	{
		ServiceSuffix: "?Action=DescribeCacheEngineVersions&Version=2015-02-02",
		Permission:    "DescribeCacheEngineVersions",
	},
	{
		ServiceSuffix: "?Action=DescribeCacheParameterGroups&Version=2015-02-02",
		Permission:    "DescribeCacheParameterGroups",
	},
	{
		ServiceSuffix: "?Action=DescribeCacheSecurityGroups",
		Permission:    "DescribeCacheSecurityGroups",
	},
	{
		ServiceSuffix: "?Action=DescribeCacheSubnetGroups&Version=2015-02-02",
		Permission:    "DescribeCacheSubnetGroups",
	},
	{
		ServiceSuffix: "?Action=DescribeReplicationGroups&Version=2015-02-02",
		Permission:    "DescribeReplicationGroups",
	},
	{
		ServiceSuffix: "?Action=DescribeReservedCacheNodes&Version=2015-02-02",
		Permission:    "DescribeReservedCacheNodes",
	},
	{
		ServiceSuffix: "?Action=DescribeReservedCacheNodesOfferings&Version=2015-02-02",
		Permission:    "DescribeReservedCacheNodesOfferings",
	},
	{
		ServiceSuffix: "?Action=DescribeSnapshots&Version=2015-02-02",
		Permission:    "DescribeSnapshots",
	},
	{
		ServiceSuffix: "?Action=ListAllowedNodeTypeModifications&Version=2015-02-02",
		Permission:    "ListAllowedNodeTypeModifications",
	},

	// extra
	{
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
	{
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
	{
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
