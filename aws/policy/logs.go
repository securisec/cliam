package policy

import "github.com/securisec/cliam/shared"

// LogsPolicies policy
var LogsPolicies = map[string]Service{
	"DescribeDestinations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeDestinations",
		},
		Permission: "DescribeDestinations",
	},
	"DescribeExportTasks": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeExportTasks",
		},
		Permission: "DescribeExportTasks",
	},
	"DescribeLogGroups": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeLogGroups",
		},
		Permission: "DescribeLogGroups",
	},
	"DescribeLogStreams": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeLogStreams",
		},
		Permission: "DescribeLogStreams",
	},
	"DescribeMetricFilters": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeMetricFilters",
		},
		Permission: "DescribeMetricFilters",
	},
	"DescribeQueries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeQueries",
		},
		Permission: "DescribeQueries",
	},
	"DescribeQueryDefinitions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeQueryDefinitions",
		},
		Permission: "DescribeQueryDefinitions",
	},
	"DescribeResourcePolicies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeResourcePolicies",
		},
		Permission: "DescribeResourcePolicies",
	},
	"FilterLogEvents": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.FilterLogEvents",
		},
		Permission: "FilterLogEvents",
	},
	"GetLogGroupFields": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.GetLogGroupFields",
		},
		Permission: "GetLogGroupFields",
	},
	"PutResourcePolicy": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.PutResourcePolicy",
		},
		Permission: "PutResourcePolicy",
	},

	// extra
	"DescribeSubscriptionFilters": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeSubscriptionFilters",
		},
		Permission:             "DescribeSubscriptionFilters",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_group_name",
	},
	"GetDataProtectionPolicy": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.GetDataProtectionPolicy",
		},
		Permission:             "GetDataProtectionPolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logGroupIdentifier",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_group_identifier",
	},
	"GetLogEvents": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.GetLogEvents",
		},
		Permission:             "GetLogEvents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logStreamName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_stream_name",
	},
	"GetLogRecord": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.GetLogRecord",
		},
		Permission:             "GetLogRecord",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logRecordPointer",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_record_pointer",
	},
	"GetQueryResults": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.GetQueryResults",
		},
		Permission:             "GetQueryResults",
		IsExtra:                true,
		ExtraComponentBodyKey:  "queryId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "query_id",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTagsLogGroup": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.ListTagsLogGroup",
		},
		Permission:             "ListTagsLogGroup",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_group_name",
	},
}
