package policy

// IOTJobsDataPolicies policy
var IOTJobsDataPolicies = map[string]Service{

	// extra
	"GetPendingJobExecutions": {
		ServiceSuffix:          "/things/{{.thing_name}}/jobs",
		Permission:             "GetPendingJobExecutions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "thing_name",
	},
}
