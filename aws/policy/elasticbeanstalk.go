package policy

import "github.com/securisec/cliam/shared"

var ElasticBeanStalkPolicies = []Service{
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeApplications&Version=2010-12-01",
		Permission:    "DescribeApplications",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeApplicationVersions&Version=2010-12-01",
		Permission:    "DescribeApplicationVersions",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeEvents&Version=2010-12-01",
		Permission:    "DescribeEvents",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountAttributes&Version=2010-12-01",
		Permission:    "DescribeAccountAttributes",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListAvailableSolutionStacks&Version=2010-12-01",
		Permission:    "ListAvailableSolutionStacks",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPlatformBranches&Version=2010-12-01",
		Permission:    "ListPlatformBranches",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
	{
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPlatformVersions&Version=2010-12-01",
		Permission:    "ListPlatformVersions",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
	},
}
