package policy

import "github.com/securisec/cliam/shared"

var AccessAnalyzersPolicies = map[string]Service{

	// extra
	"GetAnalyzer": {
		ServiceSuffix:          "/analyzer/{{.analyzer_name}}",
		Permission:             "GetAnalyzer",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "analyzer_name",
	},
	"GetGeneratedPolicy": {
		ServiceSuffix:          "/policy/generation/{{.job_id}}",
		Permission:             "GetGeneratedPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	"ListAccessPreviews": {
		ServiceSuffix:          "/access-preview",
		Permission:             "ListAccessPreviews",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "analyzer_arn",
	},
	"ListAnalyzedResources": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListAnalyzedResources",
			"Version": "2019-11-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListAnalyzedResources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "analyzerArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "analyzer_arn",
	},
	"ListArchiveRules": {
		ServiceSuffix:          "/analyzer/{{.analyzer_name}}/archive-rule",
		Permission:             "ListArchiveRules",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "analyzer_name",
	},
	"ListFindings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListFindings",
			"Version": "2019-11-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListFindings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "analyzerArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "analyzer_arn",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
