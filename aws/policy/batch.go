package policy

import "github.com/securisec/cliam/shared"

var BatchPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "v1/describecomputeenvironments",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeComputeEnvironments",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/describejobdefinitions",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeJobDefinitions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/describejobqueues",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DescribeJobQueues",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/listjobs",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/listschedulingpolicies",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "ListSchedulingPolicies",
	},

	// extra
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeJobs",
			"Version": "2016-08-10",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeJobs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "jobs",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "jobs",
	},
	{
		Method: "POST",
		FormData: map[string]string{
			"Action":  "DescribeSchedulingPolicies",
			"Version": "2016-08-10",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "DescribeSchedulingPolicies",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arns",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "arns",
	},
	{
		ServiceSuffix:          "/v1/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
