package policy

// IOTFleetHubPolicies policy
var IOTFleetHubPolicies = map[string]Service{
	"ListApplications": {
		Method:        "GET",
		ServiceSuffix: "applications",
		Permission:    "ListApplications",
	},

	// extra
	"DescribeApplication": {
		ServiceSuffix:          "/applications/{{.application_id}}",
		Permission:             "DescribeApplication",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "application_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
