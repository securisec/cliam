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

	// extra
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetInsight",
			"Version": "2016-04-12",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetInsight",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InsightId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "insight_id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetInsightEvents",
			"Version": "2016-04-12",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetInsightEvents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "InsightId",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "insight_id",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetSamplingTargets",
			"Version": "2016-04-12",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetSamplingTargets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "SamplingStatisticsDocuments",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "sampling_statistics_documents",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetTraceGraph",
			"Version": "2016-04-12",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetTraceGraph",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TraceIds",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "trace_ids",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2016-04-12",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_a_r_n",
	},
}
