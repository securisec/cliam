package policy

import "github.com/securisec/cliam/shared"

var EmailPolicies = []Service{
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListConfigurationSets",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListConfigurationSets",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeActiveReceiptRuleSet",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "DescribeActiveReceiptRuleSet",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAccountSendingEnabled",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetAccountSendingEnabled",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSendQuota",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetSendQuota",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSendStatistics",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetSendStatistics",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListCustomVerificationEmailTemplates",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListCustomVerificationEmailTemplates",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListIdentities",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListIdentities",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListReceiptFilters",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListReceiptFilters",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListReceiptRuleSets",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListReceiptRuleSets",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTemplates",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListTemplates",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListVerifiedEmailAddresses",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListVerifiedEmailAddresses",
	},
}
