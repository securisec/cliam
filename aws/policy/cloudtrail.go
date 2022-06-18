package policy

import "github.com/securisec/cliam/shared"

var CloudtrailPolicies = map[string]Service{
	"DescribeTrails": {
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "DescribeTrails",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.DescribeTrails",
		},
	},
	"ListPublicKeys": {
		Method:     "POST",
		JsonData:   map[string]string{},
		Permission: "ListPublicKeys",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListPublicKeys",
		},
	},

	// extra
	"GetEventDataStore": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetEventDataStore",
		},
		Permission:             "GetEventDataStore",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EventDataStore",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "event_data_store",
	},
	"GetEventSelectors": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetEventSelectors",
		},
		Permission:             "GetEventSelectors",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TrailName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "trail_name",
	},
	"GetInsightSelectors": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetInsightSelectors",
		},
		Permission:             "GetInsightSelectors",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TrailName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "trail_name",
	},
	"GetTrail": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetTrail",
		},
		Permission:             "GetTrail",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetTrailStatus": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetTrailStatus",
		},
		Permission:             "GetTrailStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListQueries": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListQueries",
		},
		Permission:             "ListQueries",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EventDataStore",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "event_data_store",
	},
	"ListTags": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListTags",
		},
		Permission:             "ListTags",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceIdList",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_id_list",
	},
}
