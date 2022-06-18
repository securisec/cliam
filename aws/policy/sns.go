package policy

import "github.com/securisec/cliam/shared"

var SNSPolicies = map[string]Service{
	"GetSMSAttributes": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetSMSAttributes&Version=2010-03-31",
		Permission:    "GetSMSAttributes",
	},
	"GetSMSSandboxAccountStatus": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=GetSMSSandboxAccountStatus&Version=2010-03-31",
		Permission:    "GetSMSSandboxAccountStatus",
	},
	"ListOriginationNumbers": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListOriginationNumbers&Version=2010-03-31",
		Permission:    "ListOriginationNumbers",
	},
	"ListPhoneNumbersOptedOut": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPhoneNumbersOptedOut&Version=2010-03-31",
		Permission:    "ListPhoneNumbersOptedOut",
	},
	"ListPlatformApplications": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPlatformApplications&Version=2010-03-31",
		Permission:    "ListPlatformApplications",
	},
	"ListSMSSandboxPhoneNumbers": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListSMSSandboxPhoneNumbers&Version=2010-03-31",
		Permission:    "ListSMSSandboxPhoneNumbers",
	},
	"ListSubscriptions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListSubscriptions&Version=2010-03-31",
		Permission:    "ListSubscriptions",
	},
	"ListTopics": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListTopics&Version=2010-03-31",
		Permission:    "ListTopics",
	},

	// extra policies
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
