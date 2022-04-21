package policy

import "github.com/securisec/cliam/shared"

var BatchPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "v1/describecomputeenvironments",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeComputeEnvironments",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/describejobdefinitions",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeJobDefinitions",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/describejobqueues",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeJobQueues",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/listjobs",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListJobs",
	},
	{
		Method:        "POST",
		ServiceSuffix: "v1/listschedulingpolicies",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "ListSchedulingPolicies",
	},
}
