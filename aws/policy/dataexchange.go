package policy

var DataExchangePolicies = map[string]Service{
	"ListDataSets": {
		Method:        "GET",
		ServiceSuffix: "v1/data-sets",
		Permission:    "ListDataSets",
	},
	"ListEventActions": {
		Method:        "GET",
		ServiceSuffix: "v1/event-actions",
		Permission:    "ListEventActions",
	},
	"ListJobs": {
		Method:        "GET",
		ServiceSuffix: "v1/jobs",
		Permission:    "ListJobs",
	},

	// extra
	"GetDataSet": {
		ServiceSuffix:          "/v1/data-sets/{{.data_set_id}}",
		Permission:             "GetDataSet",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "data_set_id",
	},
	"GetEventAction": {
		ServiceSuffix:          "/v1/event-actions/{{.event_action_id}}",
		Permission:             "GetEventAction",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "event_action_id",
	},
	"GetJob": {
		ServiceSuffix:          "/v1/jobs/{{.job_id}}",
		Permission:             "GetJob",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "job_id",
	},
	"ListDataSetRevisions": {
		ServiceSuffix:          "/v1/data-sets/{{.data_set_id}}/revisions",
		Permission:             "ListDataSetRevisions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "data_set_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
