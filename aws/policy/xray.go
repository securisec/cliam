package policy

import "github.com/securisec/cliam/shared"

// XrayPolicies policy
var XrayPolicies = map[string]Service{
	"GetEncryptionConfig": {
		Method:        "POST",
		ServiceSuffix: "EncryptionConfig",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetEncryptionConfig",
	},
	"GetGroup": {
		Method:        "POST",
		ServiceSuffix: "GetGroup",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetGroup",
	},
	"GetGroups": {
		Method:        "POST",
		ServiceSuffix: "Groups",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetGroups",
	},
	"GetSamplingRules": {
		Method:        "POST",
		ServiceSuffix: "GetSamplingRules",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetSamplingRules",
	},
	"GetSamplingStatisticSummaries": {
		Method:        "POST",
		ServiceSuffix: "SamplingStatisticSummaries",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "GetSamplingStatisticSummaries",
	},
	"ListResourcePolicies": {
		Method:        "POST",
		ServiceSuffix: "ListResourcePolicies",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListResourcePolicies",
	},
	"UpdateGroup": {
		Method:        "POST",
		ServiceSuffix: "UpdateGroup",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "UpdateGroup",
	},

	// extra
	"GetInsight": {
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
	"GetInsightEvents": {
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
	"GetSamplingTargets": {
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
	"GetTraceGraph": {
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
	"ListTagsForResource": {
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
