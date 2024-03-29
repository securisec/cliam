package policy

import "github.com/securisec/cliam/shared"

// DeviceFarmPolicies policy
var DeviceFarmPolicies = map[string]Service{
	"GetAccountSettings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetAccountSettings",
		},
		Permission: "GetAccountSettings",
	},
	"GetOfferingStatus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetOfferingStatus",
		},
		Permission: "GetOfferingStatus",
	},
	"GetTestGridSession": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetTestGridSession",
		},
		Permission: "GetTestGridSession",
	},
	"ListDeviceInstances": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListDeviceInstances",
		},
		Permission: "ListDeviceInstances",
	},
	"ListDevices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListDevices",
		},
		Permission: "ListDevices",
	},
	"ListInstanceProfiles": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListInstanceProfiles",
		},
		Permission: "ListInstanceProfiles",
	},
	"ListOfferingPromotions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListOfferingPromotions",
		},
		Permission: "ListOfferingPromotions",
	},
	"ListOfferingTransactions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListOfferingTransactions",
		},
		Permission: "ListOfferingTransactions",
	},
	"ListOfferings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListOfferings",
		},
		Permission: "ListOfferings",
	},
	"ListProjects": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListProjects",
		},
		Permission: "ListProjects",
	},
	"ListTestGridProjects": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListTestGridProjects",
		},
		Permission: "ListTestGridProjects",
	},
	"ListVPCEConfigurations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListVPCEConfigurations",
		},
		Permission: "ListVPCEConfigurations",
	},

	// extra
	"GetDevice": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetDevice",
		},
		Permission:             "GetDevice",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetDeviceInstance": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetDeviceInstance",
		},
		Permission:             "GetDeviceInstance",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetDevicePool": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetDevicePool",
		},
		Permission:             "GetDevicePool",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetDevicePoolCompatibility": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetDevicePoolCompatibility",
		},
		Permission:             "GetDevicePoolCompatibility",
		IsExtra:                true,
		ExtraComponentBodyKey:  "devicePoolArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "device_pool_arn",
	},
	"GetInstanceProfile": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetInstanceProfile",
		},
		Permission:             "GetInstanceProfile",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetJob",
		},
		Permission:             "GetJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetNetworkProfile": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetNetworkProfile",
		},
		Permission:             "GetNetworkProfile",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetProject": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetProject",
		},
		Permission:             "GetProject",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetRemoteAccessSession": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetRemoteAccessSession",
		},
		Permission:             "GetRemoteAccessSession",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetRun": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetRun",
		},
		Permission:             "GetRun",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetSuite": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetSuite",
		},
		Permission:             "GetSuite",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetTest": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetTest",
		},
		Permission:             "GetTest",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetTestGridProject": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetTestGridProject",
		},
		Permission:             "GetTestGridProject",
		IsExtra:                true,
		ExtraComponentBodyKey:  "projectArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "project_arn",
	},
	"GetUpload": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetUpload",
		},
		Permission:             "GetUpload",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"GetVPCEConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.GetVPCEConfiguration",
		},
		Permission:             "GetVPCEConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListDevicePools": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListDevicePools",
		},
		Permission:             "ListDevicePools",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListJobs": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListJobs",
		},
		Permission:             "ListJobs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListNetworkProfiles": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListNetworkProfiles",
		},
		Permission:             "ListNetworkProfiles",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListRemoteAccessSessions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListRemoteAccessSessions",
		},
		Permission:             "ListRemoteAccessSessions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListRuns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListRuns",
		},
		Permission:             "ListRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListSamples": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListSamples",
		},
		Permission:             "ListSamples",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListSuites": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListSuites",
		},
		Permission:             "ListSuites",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTestGridSessionActions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListTestGridSessionActions",
		},
		Permission:             "ListTestGridSessionActions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "sessionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_arn",
	},
	"ListTestGridSessionArtifacts": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListTestGridSessionArtifacts",
		},
		Permission:             "ListTestGridSessionArtifacts",
		IsExtra:                true,
		ExtraComponentBodyKey:  "sessionArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "session_arn",
	},
	"ListTestGridSessions": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListTestGridSessions",
		},
		Permission:             "ListTestGridSessions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "projectArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "project_arn",
	},
	"ListTests": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListTests",
		},
		Permission:             "ListTests",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListUniqueProblems": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListUniqueProblems",
		},
		Permission:             "ListUniqueProblems",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
	"ListUploads": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "DeviceFarm_20150623.ListUploads",
		},
		Permission:             "ListUploads",
		IsExtra:                true,
		ExtraComponentBodyKey:  "arn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "arn",
	},
}
