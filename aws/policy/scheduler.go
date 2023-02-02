package policy

// SchedulerPolicies policy
var SchedulerPolicies = map[string]Service{
	"ListScheduleGroups": {
		Method:        "GET",
		ServiceSuffix: "schedule-groups",
		Permission:    "ListScheduleGroups",
	},
	"ListSchedules": {
		Method:        "GET",
		ServiceSuffix: "schedules",
		Permission:    "ListSchedules",
	},

	// extra
	"GetSchedule": {
		ServiceSuffix:          "/schedules/{{.name}}",
		Permission:             "GetSchedule",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"GetScheduleGroup": {
		ServiceSuffix:          "/schedule-groups/{{.name}}",
		Permission:             "GetScheduleGroup",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "name",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
