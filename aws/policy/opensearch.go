package policy

import "github.com/securisec/cliam/shared"

// OpensearchPolicies policies
var OpensearchPolicies = map[string]Service{
	"DescribeInboundConnections": {
		Method:        "POST",
		ServiceSuffix: "2021-01-01/opensearch/cc/inboundConnection/search",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeInboundConnections",
	},
	"DescribeOutboundConnections": {
		Method:        "POST",
		ServiceSuffix: "2021-01-01/opensearch/cc/outboundConnection/search",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeOutboundConnections",
	},
	"DescribePackages": {
		Method:        "POST",
		ServiceSuffix: "2021-01-01/packages/describe",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribePackages",
	},
	"DescribeReservedInstanceOfferings": {
		Method:        "GET",
		ServiceSuffix: "2021-01-01/opensearch/reservedInstanceOfferings",
		Permission:    "DescribeReservedInstanceOfferings",
	},
	"DescribeReservedInstances": {
		Method:        "GET",
		ServiceSuffix: "2021-01-01/opensearch/reservedInstances",
		Permission:    "DescribeReservedInstances",
	},
	"GetCompatibleVersions": {
		Method:        "GET",
		ServiceSuffix: "2021-01-01/opensearch/compatibleVersions",
		Permission:    "GetCompatibleVersions",
	},
	"ListDomainNames": {
		Method:        "GET",
		ServiceSuffix: "2021-01-01/domain",
		Permission:    "ListDomainNames",
		ResponseParser: &ResponseParser{
			ResponseFormat: "json",
			KeysToExtract: []CommandLineFlagMap{
				{Flag: "domain_name", ResponseKey: "DomainName"},
			},
			ObjectPath: []string{"DomainNames"},
		},
	},
	"ListVersions": {
		Method:        "GET",
		ServiceSuffix: "2021-01-01/opensearch/versions",
		Permission:    "ListVersions",
	},
	"ListVpcEndpoints": {
		Method:        "GET",
		ServiceSuffix: "2021-01-01/opensearch/vpcEndpoints",
		Permission:    "ListVpcEndpoints",
	},

	// extra
	"DescribeDomain": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}",
		Permission:             "DescribeDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeDomainAutoTunes": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}/autoTunes",
		Permission:             "DescribeDomainAutoTunes",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeDomainChangeProgress": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}/progress",
		Permission:             "DescribeDomainChangeProgress",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeDomainConfig": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}/config",
		Permission:             "DescribeDomainConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeDomains": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDomains",
			"Version": "2021-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDomains",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainNames",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_names",
	},
	"DescribeDryRunProgress": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}/dryRun",
		Permission:             "DescribeDryRunProgress",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeVpcEndpoints": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeVpcEndpoints",
			"Version": "2021-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeVpcEndpoints",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VpcEndpointIds",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "vpc_endpoint_ids",
	},
	"GetPackageVersionHistory": {
		ServiceSuffix:          "/2021-01-01/packages/{{.package_i_d}}/history",
		Permission:             "GetPackageVersionHistory",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "package_id",
	},
	"GetUpgradeHistory": {
		ServiceSuffix:          "/2021-01-01/opensearch/upgradeDomain/{{.domain_name}}/history",
		Permission:             "GetUpgradeHistory",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"GetUpgradeStatus": {
		ServiceSuffix:          "/2021-01-01/opensearch/upgradeDomain/{{.domain_name}}/status",
		Permission:             "GetUpgradeStatus",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"ListDomainsForPackage": {
		ServiceSuffix:          "/2021-01-01/packages/{{.package_i_d}}/domains",
		Permission:             "ListDomainsForPackage",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "package_id",
	},
	"ListInstanceTypeDetails": {
		ServiceSuffix:          "/2021-01-01/opensearch/instanceTypeDetails/{{.engine_version}}",
		Permission:             "ListInstanceTypeDetails",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "engine_version",
	},
	"ListPackagesForDomain": {
		ServiceSuffix:          "/2021-01-01/domain/{{.domain_name}}/packages",
		Permission:             "ListPackagesForDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"ListTags": {
		ServiceSuffix:          "/2021-01-01/tags/",
		Permission:             "ListTags",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
	"ListVpcEndpointAccess": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}/listVpcEndpointAccess",
		Permission:             "ListVpcEndpointAccess",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
	"ListVpcEndpointsForDomain": {
		ServiceSuffix:          "/2021-01-01/opensearch/domain/{{.domain_name}}/vpcEndpoints",
		Permission:             "ListVpcEndpointsForDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain_name",
	},
}
