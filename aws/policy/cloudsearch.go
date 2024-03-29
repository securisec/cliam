package policy

import "github.com/securisec/cliam/shared"

// CloudSearchPolicies policy
var CloudSearchPolicies = map[string]Service{
	"DescribeDomains": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDomains",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeDomains",
	},
	"ListDomainNames": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListDomainNames",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListDomainNames",
	},

	// extra
	"DescribeAnalysisSchemes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAnalysisSchemes",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeAnalysisSchemes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeAvailabilityOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeAvailabilityOptions",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeAvailabilityOptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeDomainEndpointOptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeDomainEndpointOptions",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeDomainEndpointOptions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeExpressions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeExpressions",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeExpressions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeIndexFields": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeIndexFields",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeIndexFields",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeScalingParameters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeScalingParameters",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeScalingParameters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeServiceAccessPolicies": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeServiceAccessPolicies",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeServiceAccessPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
	"DescribeSuggesters": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSuggesters",
			"Version": "2013-01-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSuggesters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain_name",
	},
}
