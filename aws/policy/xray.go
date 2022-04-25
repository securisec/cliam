package policy

import "github.com/securisec/cliam/shared"

var XrayPolicies = []Service{
	{
		ServiceSuffix: "Groups",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetGroups",
	},
	{
		ServiceSuffix: "SamplingStatisticSummaries",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetSamplingStatisticSummaries",
	},
	{
		ServiceSuffix: "GetSamplingRules",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetSamplingRules",
	},
	{
		ServiceSuffix: "EncryptionConfig",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "EncryptionConfig",
	},
}
