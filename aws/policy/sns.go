package policy

import "github.com/securisec/cliam/shared"

// SNSPolicies policy
var SNSPolicies = map[string]Service{
	"GetSMSAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSMSAttributes",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetSMSAttributes",
	},
	"GetSMSSandboxAccountStatus": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSMSSandboxAccountStatus",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "GetSMSSandboxAccountStatus",
	},
	"ListOriginationNumbers": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListOriginationNumbers",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListOriginationNumbers",
	},
	"ListPhoneNumbersOptedOut": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListPhoneNumbersOptedOut",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListPhoneNumbersOptedOut",
	},
	"ListPlatformApplications": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListPlatformApplications",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListPlatformApplications",
	},
	"ListSMSSandboxPhoneNumbers": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListSMSSandboxPhoneNumbers",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListSMSSandboxPhoneNumbers",
	},
	"ListSubscriptions": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListSubscriptions",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListSubscriptions",
	},
	"ListTopics": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTopics",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission: "ListTopics",
	},

	// extra
	"GetDataProtectionPolicy": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetDataProtectionPolicy",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetDataProtectionPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"GetEndpointAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetEndpointAttributes",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetEndpointAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "endpoint_arn",
	},
	"GetPlatformApplicationAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetPlatformApplicationAttributes",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetPlatformApplicationAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PlatformApplicationArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "platform_application_arn",
	},
	"GetSubscriptionAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSubscriptionAttributes",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetSubscriptionAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SubscriptionArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "subscription_arn",
	},
	"GetTopicAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTopicAttributes",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTopicAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TopicArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "topic_arn",
	},
	"ListEndpointsByPlatformApplication": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListEndpointsByPlatformApplication",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListEndpointsByPlatformApplication",
		IsExtra:                true,
		ExtraComponentBodyKey:  "PlatformApplicationArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "platform_application_arn",
	},
	"ListSubscriptionsByTopic": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListSubscriptionsByTopic",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListSubscriptionsByTopic",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TopicArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "topic_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2010-03-31",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
