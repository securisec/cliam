package policy

import "github.com/securisec/cliam/shared"

// CodeArtifactPolicies policies
var CodeArtifactPolicies = map[string]Service{
	"ListDomains": {
		Method:        "POST",
		ServiceSuffix: "v1/domains",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListDomains",
	},
	"ListRepositories": {
		Method:        "POST",
		ServiceSuffix: "v1/repositories",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListRepositories",
	},

	// extra
	"DescribeDomain": {
		ServiceSuffix:          "/v1/domain",
		Permission:             "DescribeDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain",
	},
	"GetAuthorizationToken": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "GetAuthorizationToken",
			"Version": "2018-09-22",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "GetAuthorizationToken",
		IsExtra:                true,
		ExtraComponentBodyKey:  "domain",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain",
	},
	"GetDomainPermissionsPolicy": {
		ServiceSuffix:          "/v1/domain/permissions/policy",
		Permission:             "GetDomainPermissionsPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "domain",
	},
	"ListRepositoriesInDomain": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListRepositoriesInDomain",
			"Version": "2018-09-22",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListRepositoriesInDomain",
		IsExtra:                true,
		ExtraComponentBodyKey:  "domain",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "domain",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2018-09-22",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
