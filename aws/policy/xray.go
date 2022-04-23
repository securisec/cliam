package policy

import "github.com/securisec/cliam/shared"

var XrayPolicies = []Service{
	{
		ServiceSuffix: "Groups",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetGroups",
	},
	{
		ServiceSuffix: "SamplingStatisticSummaries",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetSamplingStatisticSummaries",
	},
	{
		ServiceSuffix: "GetSamplingRules",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "GetSamplingRules",
	},
	{
		ServiceSuffix: "EncryptionConfig",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "EncryptionConfig",
	},
}
