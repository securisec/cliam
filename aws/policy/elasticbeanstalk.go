package policy

import "github.com/securisec/enumerate/shared"

var ElasticBeanStalkPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeApplications&Version=2010-12-01",
		Permission:    "DescribeApplications",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeApplicationVersions&Version=2010-12-01",
		Permission:    "DescribeApplicationVersions",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeEvents&Version=2010-12-01",
		Permission:    "DescribeEvents",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=DescribeAccountAttributes&Version=2010-12-01",
		Permission:    "DescribeAccountAttributes",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListAvailableSolutionStacks&Version=2010-12-01",
		Permission:    "ListAvailableSolutionStacks",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListPlatformBranches&Version=2010-12-01",
		Permission:    "ListPlatformBranches",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      `{}`,
		ServiceSuffix: "?Action=ListPlatformVersions&Version=2010-12-01",
		Permission:    "ListPlatformVersions",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
}
