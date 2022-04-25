package policy

import "github.com/securisec/cliam/shared"

var ESPolicies = []Service{
	{
		ServiceSuffix: "2021-01-01/domain",
		Permission:    "ListDomainNames",
	},
	{
		ServiceSuffix: "2015-01-01/es/reservedInstances",
		Permission:    "DescribeReservedElasticsearchInstances",
	},
	{
		ServiceSuffix: "2015-01-01/es/reservedInstanceOfferings",
		Permission:    "DescribeReservedElasticsearchInstancesOfferings",
	},
	{
		ServiceSuffix: "2015-01-01/es/compatibleVersions",
		Permission:    "GetCompatibleElasticsearchVersions",
	},
	{
		ServiceSuffix: "2015-01-01/es/versions",
		Permission:    "ListElasticsearchVersions",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/inboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeInboundCrossClusterSearchConnections",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOutboundCrossClusterSearchConnections",
	},
	{
		ServiceSuffix: "2021-01-01/packages/describe",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribePackages",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/inboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeInboundConnections",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOutboundConnections",
	},
	{
		ServiceSuffix: "2015-01-01/packages/describe",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOutboundConnections",
	},
}
