package policy

// IOTEventsDataPolicies policy
var IOTEventsDataPolicies = map[string]Service{

	// extra
	"DescribeAlarm": {
		ServiceSuffix:          "/alarms/{{.alarm_model_name}}/keyValues/",
		Permission:             "DescribeAlarm",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "alarm_model_name",
	},
	"DescribeDetector": {
		ServiceSuffix:          "/detectors/{{.detector_model_name}}/keyValues/",
		Permission:             "DescribeDetector",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_model_name",
	},
	"ListAlarms": {
		ServiceSuffix:          "/alarms/{{.alarm_model_name}}",
		Permission:             "ListAlarms",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "alarm_model_name",
	},
	"ListDetectors": {
		ServiceSuffix:          "/detectors/{{.detector_model_name}}",
		Permission:             "ListDetectors",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_model_name",
	},
}
