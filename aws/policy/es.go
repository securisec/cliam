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
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeInboundCrossClusterSearchConnections",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeOutboundCrossClusterSearchConnections",
	},
	{
		ServiceSuffix: "2021-01-01/packages/describe",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribePackages",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/inboundConnection/search",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeInboundConnections",
	},
	{
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeOutboundConnections",
	},
	{
		ServiceSuffix: "2015-01-01/packages/describe",
		Method:        "POST",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeOutboundConnections",
	},
}
