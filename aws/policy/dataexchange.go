package policy

var DataExchangePolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "v1/data-sets",
		Permission:    "ListDataSets",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v1/event-actions",
		Permission:    "ListEventActions",
	},
	{
		Method:        "GET",
		ServiceSuffix: "v1/jobs",
		Permission:    "ListJobs",
	},

	// extra
	{
		ServiceSuffix:          "/v1/data-sets/{{.data_set_id}}",
		Permission:             "GetDataSet",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "data_set_id",
	},
	{
		ServiceSuffix:          "/v1/event-actions/{{.event_action_id}}",
		Permission:             "GetEventAction",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "event_action_id",
	},
	{
		ServiceSuffix:          "/v1/jobs/{{.job_id}}",
		Permission:             "GetJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	{
		ServiceSuffix:          "/v1/data-sets/{{.data_set_id}}/revisions",
		Permission:             "ListDataSetRevisions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "data_set_id",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
