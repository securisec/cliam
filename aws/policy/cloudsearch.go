package policy

import "github.com/securisec/iam-enumerate/shared"

var CloudsearchPolicies = []Service{
	{
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
	{
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
}
