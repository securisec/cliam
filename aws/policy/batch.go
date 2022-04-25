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
}
