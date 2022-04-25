package policy

import "github.com/securisec/cliam/shared"

var ECRPublicPolicies = []Service{
	{
		Permission: "DescribeRepositories",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			"x-amz-target":             "SpencerFrontendService.DescribeRepositories",
		},
	},
	{
		Permission: "DescribeRegistries",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			"x-amz-target":             "SpencerFrontendService.DescribeRegistries",
		},
	},
	{
		Permission: "GetAuthorizationToken",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			"x-amz-target":             "SpencerFrontendService.GetAuthorizationToken",
		},
	},
	{
		Permission: "GetRegistryCatalogData",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			"x-amz-target":             "SpencerFrontendService.GetRegistryCatalogData",
		},
	},

	// extra policies
	{
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
	{
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
	{
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
	{
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
	{
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
