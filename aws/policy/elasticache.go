package policy

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
}
