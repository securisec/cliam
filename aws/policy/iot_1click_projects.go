package policy

// IOT1ClickProjectsPolicies policy
var IOT1ClickProjectsPolicies = map[string]Service{
	"ListProjects": {
		Method:        "GET",
		ServiceSuffix: "projects",
		Permission:    "ListProjects",
	},

	// extra
	"DescribeProject": {
		ServiceSuffix:          "/projects/{{.project_name}}",
		Permission:             "DescribeProject",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "project_name",
	},
	"ListPlacements": {
		ServiceSuffix:          "/projects/{{.project_name}}/placements",
		Permission:             "ListPlacements",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "project_name",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
