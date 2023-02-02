package policy

import "github.com/securisec/cliam/shared"

// ECRPublicPolicies policy
var ECRPublicPolicies = map[string]Service{
	"DescribeRegistries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.DescribeRegistries",
		},
		Permission: "DescribeRegistries",
	},
	"DescribeRepositories": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.DescribeRepositories",
		},
		Permission: "DescribeRepositories",
	},
	"GetAuthorizationToken": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.GetAuthorizationToken",
		},
		Permission: "GetAuthorizationToken",
	},
	"GetRegistryCatalogData": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.GetRegistryCatalogData",
		},
		Permission: "GetRegistryCatalogData",
	},
	"PutRegistryCatalogData": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.PutRegistryCatalogData",
		},
		Permission: "PutRegistryCatalogData",
	},

	// extra
	"DescribeImageTags": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.DescribeImageTags",
		},
		Permission:             "DescribeImageTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"DescribeImages": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.DescribeImages",
		},
		Permission:             "DescribeImages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"GetRepositoryCatalogData": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.GetRepositoryCatalogData",
		},
		Permission:             "GetRepositoryCatalogData",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"GetRepositoryPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.GetRepositoryPolicy",
		},
		Permission:             "GetRepositoryPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "SpencerFrontendService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
