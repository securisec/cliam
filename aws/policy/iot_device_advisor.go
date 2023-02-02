package policy

// IOTDeviceAdvisorPolicies policy
var IOTDeviceAdvisorPolicies = map[string]Service{
	"GetEndpoint": {
		Method:        "GET",
		ServiceSuffix: "endpoint",
		Permission:    "GetEndpoint",
	},
	"ListSuiteDefinitions": {
		Method:        "GET",
		ServiceSuffix: "suiteDefinitions",
		Permission:    "ListSuiteDefinitions",
	},
	"ListSuiteRuns": {
		Method:        "GET",
		ServiceSuffix: "suiteRuns",
		Permission:    "ListSuiteRuns",
	},

	// extra
	"GetSuiteDefinition": {
		ServiceSuffix:          "/suiteDefinitions/{{.suite_definition_id}}",
		Permission:             "GetSuiteDefinition",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "suite_definition_id",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
