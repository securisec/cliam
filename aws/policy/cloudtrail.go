package policy

import "github.com/securisec/cliam/shared"

// CloudTrailPolicies policy
var CloudTrailPolicies = map[string]Service{
	"DescribeTrails": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.DescribeTrails",
		},
		Permission: "DescribeTrails",
	},
	"ListChannels": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListChannels",
		},
		Permission: "ListChannels",
	},
	"ListEventDataStores": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListEventDataStores",
		},
		Permission: "ListEventDataStores",
	},
	"ListImports": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListImports",
		},
		Permission: "ListImports",
	},
	"ListPublicKeys": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListPublicKeys",
		},
		Permission: "ListPublicKeys",
	},
	"ListTrails": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListTrails",
		},
		Permission: "ListTrails",
	},
	"LookupEvents": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.LookupEvents",
		},
		Permission: "LookupEvents",
	},
	"StartImport": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.StartImport",
		},
		Permission: "StartImport",
	},

	// extra
	"DescribeQuery": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.DescribeQuery",
		},
		Permission:             "DescribeQuery",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_id",
	},
	"GetChannel": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetChannel",
		},
		Permission:             "GetChannel",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Channel",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "channel",
	},
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
	"GetImport": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetImport",
		},
		Permission:             "GetImport",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ImportId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "import_id",
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
	"GetQueryResults": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetQueryResults",
		},
		Permission:             "GetQueryResults",
		IsExtra:                true,
		ExtraComponentBodyKey:  "QueryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_id",
	},
	"GetResourcePolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.GetResourcePolicy",
		},
		Permission:             "GetResourcePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
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
	"ListImportFailures": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "com.amazonaws.cloudtrail.v20131101.CloudTrail_20131101.ListImportFailures",
		},
		Permission:             "ListImportFailures",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ImportId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "import_id",
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
