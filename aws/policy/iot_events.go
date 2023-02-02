package policy

import "github.com/securisec/cliam/shared"

// IOTEventsPolicies policy
var IOTEventsPolicies = map[string]Service{
	"DescribeLoggingOptions": {
		Method:        "GET",
		ServiceSuffix: "logging",
		Permission:    "DescribeLoggingOptions",
	},
	"ListAlarmModels": {
		Method:        "GET",
		ServiceSuffix: "alarm-models",
		Permission:    "ListAlarmModels",
	},
	"ListDetectorModels": {
		Method:        "GET",
		ServiceSuffix: "detector-models",
		Permission:    "ListDetectorModels",
	},
	"ListInputs": {
		Method:        "GET",
		ServiceSuffix: "inputs",
		Permission:    "ListInputs",
	},

	// extra
	"DescribeAlarmModel": {
		ServiceSuffix:          "/alarm-models/{{.alarm_model_name}}",
		Permission:             "DescribeAlarmModel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "alarm_model_name",
	},
	"DescribeDetectorModel": {
		ServiceSuffix:          "/detector-models/{{.detector_model_name}}",
		Permission:             "DescribeDetectorModel",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_model_name",
	},
	"DescribeDetectorModelAnalysis": {
		ServiceSuffix:          "/analysis/detector-models/{{.analysis_id}}",
		Permission:             "DescribeDetectorModelAnalysis",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "analysis_id",
	},
	"DescribeInput": {
		ServiceSuffix:          "/inputs/{{.input_name}}",
		Permission:             "DescribeInput",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "input_name",
	},
	"GetDetectorModelAnalysisResults": {
		ServiceSuffix:          "/analysis/detector-models/{{.analysis_id}}/results",
		Permission:             "GetDetectorModelAnalysisResults",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "analysis_id",
	},
	"ListAlarmModelVersions": {
		ServiceSuffix:          "/alarm-models/{{.alarm_model_name}}/versions",
		Permission:             "ListAlarmModelVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "alarm_model_name",
	},
	"ListDetectorModelVersions": {
		ServiceSuffix:          "/detector-models/{{.detector_model_name}}/versions",
		Permission:             "ListDetectorModelVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "detector_model_name",
	},
	"ListInputRoutings": {
		Method: "POST",
		FormData: map[string]string{
			"Action":  "ListInputRoutings",
			"Version": "2018-07-27",
		},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: shared.CONTENT_TYPE_URL_ENCODED,
		},
		Permission:             "ListInputRoutings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "inputIdentifier",
		ExtraComponentLocation: "form",
		ExtraCommandLineFlag:   "input_identifier",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
