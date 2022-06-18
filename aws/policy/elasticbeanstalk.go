package policy

import "github.com/securisec/cliam/shared"

var ElasticBeanStalkPolicies = map[string]Service{
	"DescribeApplications": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeApplications&Version=2010-12-01",
		Permission:    "DescribeApplications",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},
	"DescribeApplicationVersions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeApplicationVersions&Version=2010-12-01",
		Permission:    "DescribeApplicationVersions",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},
	"DescribeEvents": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeEvents&Version=2010-12-01",
		Permission:    "DescribeEvents",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},
	"DescribeAccountAttributes": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=DescribeAccountAttributes&Version=2010-12-01",
		Permission:    "DescribeAccountAttributes",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},
	"ListAvailableSolutionStacks": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListAvailableSolutionStacks&Version=2010-12-01",
		Permission:    "ListAvailableSolutionStacks",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},
	"ListPlatformBranches": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPlatformBranches&Version=2010-12-01",
		Permission:    "ListPlatformBranches",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},
	"ListPlatformVersions": {
		Method:        "POST",
		JsonData:      map[string]string{},
		ServiceSuffix: "?Action=ListPlatformVersions&Version=2010-12-01",
		Permission:    "ListPlatformVersions",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
	},

	// extra
	"DescribeInstanceHealth": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeInstanceHealth",
			"Version": "2012-06-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeInstanceHealth",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_name",
	},
	"DescribeLoadBalancerAttributes": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeLoadBalancerAttributes",
			"Version": "2012-06-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeLoadBalancerAttributes",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_name",
	},
	"DescribeTags": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeTags",
			"Version": "2012-06-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "LoadBalancerNames",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "load_balancer_names",
	},

	// extra
	"DescribeConfigurationSettings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeConfigurationSettings",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeConfigurationSettings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ApplicationName",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "application_name",
	},
	"ListTagsForResource": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListTagsForResource",
			"Version": "2010-12-01",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
