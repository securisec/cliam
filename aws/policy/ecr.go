package policy

import "github.com/securisec/cliam/shared"

var ECRPolicies = map[string]Service{
	"DescribeRepositories": {
		ServiceSuffix: "",
		Permission:    "DescribeRepositories",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerRegistry_V20150921.DescribeRepositories",
		},
	},
	"GetAuthorizationToken": {
		ServiceSuffix: "",
		Permission:    "GetAuthorizationToken",
		Method:        "POST",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			"Content-Type": aws_JSON_1_1,
			"x-amz-target": "AmazonEC2ContainerRegistry_V20150921.GetAuthorizationToken",
		},
	},

	// extra policies
	"DescribeImages": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerRegistry_V20150921.DescribeImages",
		},
		Permission:             "DescribeImages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"GetLifecyclePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerRegistry_V20150921.GetLifecyclePolicy",
		},
		Permission:             "GetLifecyclePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"GetLifecyclePolicyPreview": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerRegistry_V20150921.GetLifecyclePolicyPreview",
		},
		Permission:             "GetLifecyclePolicyPreview",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"GetRepositoryPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerRegistry_V20150921.GetRepositoryPolicy",
		},
		Permission:             "GetRepositoryPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"ListImages": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerRegistry_V20150921.ListImages",
		},
		Permission:             "ListImages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "repositoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "repository_name",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AmazonEC2ContainerRegistry_V20150921.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
