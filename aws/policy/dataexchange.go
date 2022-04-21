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
}
