package policy

import "github.com/securisec/cliam/shared"

// EmailPolicies policy
var EmailPolicies = map[string]Service{
	"DescribeActiveReceiptRuleSet": {
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
	"GetAccountSendingEnabled": {
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
	"GetSendQuota": {
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
	"GetSendStatistics": {
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
	"ListConfigurationSets": {
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
	"ListCustomVerificationEmailTemplates": {
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
	"ListIdentities": {
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
	"ListReceiptFilters": {
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
	"ListReceiptRuleSets": {
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
	"ListTemplates": {
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
	"ListVerifiedEmailAddresses": {
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
	"SetActiveReceiptRuleSet": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "SetActiveReceiptRuleSet",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "SetActiveReceiptRuleSet",
	},
	"UpdateAccountSendingEnabled": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "UpdateAccountSendingEnabled",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "UpdateAccountSendingEnabled",
	},

	// extra
	"DescribeConfigurationSet": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConfigurationSet",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeConfigurationSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ConfigurationSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "configuration_set_name",
	},
	"DescribeReceiptRuleSet": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeReceiptRuleSet",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeReceiptRuleSet",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RuleSetName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "rule_set_name",
	},
	"GetCustomVerificationEmailTemplate": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetCustomVerificationEmailTemplate",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetCustomVerificationEmailTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TemplateName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "template_name",
	},
	"GetIdentityDkimAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIdentityDkimAttributes",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIdentityDkimAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Identities",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "identities",
	},
	"GetIdentityMailFromDomainAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIdentityMailFromDomainAttributes",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIdentityMailFromDomainAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Identities",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "identities",
	},
	"GetIdentityNotificationAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIdentityNotificationAttributes",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIdentityNotificationAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Identities",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "identities",
	},
	"GetIdentityVerificationAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetIdentityVerificationAttributes",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetIdentityVerificationAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Identities",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "identities",
	},
	"GetTemplate": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTemplate",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTemplate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TemplateName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "template_name",
	},
	"ListIdentityPolicies": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListIdentityPolicies",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListIdentityPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Identity",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "identity",
	},
}
