package policy

import "github.com/securisec/cliam/shared"

var MediaConvertPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "2017-08-29/endpoints",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
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

	// extra
	{
		ServiceSuffix:          "/2017-08-29/jobs/{{.id}}",
		Permission:             "GetJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2017-08-29/jobTemplates/{{.name}}",
		Permission:             "GetJobTemplate",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/2017-08-29/presets/{{.name}}",
		Permission:             "GetPreset",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/2017-08-29/queues/{{.name}}",
		Permission:             "GetQueue",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	{
		ServiceSuffix:          "/2017-08-29/tags/{{.arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "arn",
	},
}
