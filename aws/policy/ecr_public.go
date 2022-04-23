package policy

import "github.com/securisec/cliam/shared"

var ECRPublicPolicies = []Service{
	{
		Permission: "DescribeRepositories",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			"x-amz-target":             "SpencerFrontendService.DescribeRepositories",
		},
	},
	{
		Permission: "DescribeRegistries",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			"x-amz-target":             "SpencerFrontendService.DescribeRegistries",
		},
	},
	{
		Permission: "GetAuthorizationToken",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			"x-amz-target":             "SpencerFrontendService.GetAuthorizationToken",
		},
	},
	{
		Permission: "GetRegistryCatalogData",
		Method:     "POST",
		JsonData:   map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			"x-amz-target":             "SpencerFrontendService.GetRegistryCatalogData",
		},
	},
}
