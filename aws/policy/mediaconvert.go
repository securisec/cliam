package policy

import "github.com/securisec/cliam/shared"

var MediaConvertPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "2017-08-29/endpoints",
		JsonData:      `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
		},
		Permission: "DescribeEndpoints",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2017-08-29/policy",
		Permission:    "GetPolicy",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2017-08-29/jobTemplates",
		Permission:    "ListJobTemplates",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2017-08-29/jobs",
		Permission:    "ListJobs",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2017-08-29/presets",
		Permission:    "ListPresets",
	},
	{
		Method:        "GET",
		ServiceSuffix: "2017-08-29/queues",
		Permission:    "ListQueues",
	},
}
