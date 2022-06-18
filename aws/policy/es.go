package policy

import "github.com/securisec/cliam/shared"

var ESPolicies = map[string]Service{
	"ListDomainNames": {
		ServiceSuffix: "2021-01-01/domain",
		Permission:    "ListDomainNames",
	},
	"DescribeReservedElasticsearchInstances": {
		ServiceSuffix: "2015-01-01/es/reservedInstances",
		Permission:    "DescribeReservedElasticsearchInstances",
	},
	"DescribeReservedElasticsearchInstancesOfferings": {
		ServiceSuffix: "2015-01-01/es/reservedInstanceOfferings",
		Permission:    "DescribeReservedElasticsearchInstancesOfferings",
	},
	"GetCompatibleElasticsearchVersions": {
		ServiceSuffix: "2015-01-01/es/compatibleVersions",
		Permission:    "GetCompatibleElasticsearchVersions",
	},
	"ListElasticsearchVersions": {
		ServiceSuffix: "2015-01-01/es/versions",
		Permission:    "ListElasticsearchVersions",
	},
	"DescribeInboundCrossClusterSearchConnections": {
		ServiceSuffix: "2021-01-01/opensearch/cc/inboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeInboundCrossClusterSearchConnections",
	},
	"DescribeOutboundCrossClusterSearchConnections": {
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOutboundCrossClusterSearchConnections",
	},
	"DescribePackages": {
		ServiceSuffix: "2021-01-01/packages/describe",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribePackages",
	},
	"DescribeInboundConnections": {
		ServiceSuffix: "2021-01-01/opensearch/cc/inboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeInboundConnections",
	},
	"DescribeOutboundConnections": {
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOutboundConnections",
	},

	// extra
	"DescribeDomainAutoTunes": {
		ServiceSuffix:          "/2015-01-01/es/domain/{{.domain_name}}/autoTunes",
		Permission:             "DescribeDomainAutoTunes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeDomainChangeProgress": {
		ServiceSuffix:          "/2015-01-01/es/domain/{{.domain_name}}/progress",
		Permission:             "DescribeDomainChangeProgress",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeElasticsearchDomain": {
		ServiceSuffix:          "/2015-01-01/es/domain/{{.domain_name}}",
		Permission:             "DescribeElasticsearchDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeElasticsearchDomainConfig": {
		ServiceSuffix:          "/2015-01-01/es/domain/{{.domain_name}}/config",
		Permission:             "DescribeElasticsearchDomainConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeElasticsearchDomains": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeElasticsearchDomains",
			"Version": "2015-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeElasticsearchDomains",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainNames",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_names",
	},
	"GetPackageVersionHistory": {
		ServiceSuffix:          "/2015-01-01/packages/{{.package_i_d}}/history",
		Permission:             "GetPackageVersionHistory",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "package_i_d",
	},
	"GetUpgradeHistory": {
		ServiceSuffix:          "/2015-01-01/es/upgradeDomain/{{.domain_name}}/history",
		Permission:             "GetUpgradeHistory",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetUpgradeStatus": {
		ServiceSuffix:          "/2015-01-01/es/upgradeDomain/{{.domain_name}}/status",
		Permission:             "GetUpgradeStatus",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"ListDomainsForPackage": {
		ServiceSuffix:          "/2015-01-01/packages/{{.package_i_d}}/domains",
		Permission:             "ListDomainsForPackage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "package_i_d",
	},
	"ListElasticsearchInstanceTypes": {
		ServiceSuffix:          "/2015-01-01/es/instanceTypes/{{.elasticsearch_version}}",
		Permission:             "ListElasticsearchInstanceTypes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "elasticsearch_version",
	},
	"ListPackagesForDomain": {
		ServiceSuffix:          "/2015-01-01/domain/{{.domain_name}}/packages",
		Permission:             "ListPackagesForDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"ListTags": {
		ServiceSuffix:          "/2015-01-01/tags/",
		Permission:             "ListTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
}
