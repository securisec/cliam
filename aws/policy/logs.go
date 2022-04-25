package policy

import "github.com/securisec/cliam/shared"

var LogsPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeDestinations",
		},
		Permission: "DescribeDestinations",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeExportTasks",
		},
		Permission: "DescribeExportTasks",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeLogGroups",
		},
		Permission: "DescribeLogGroups",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeMetricFilters",
		},
		Permission: "DescribeMetricFilters",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeQueries",
		},
		Permission: "DescribeQueries",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeQueryDefinitions",
		},
		Permission: "DescribeQueryDefinitions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeResourcePolicies",
		},
		Permission: "DescribeResourcePolicies",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.DescribeLogStreams",
		},
		Permission:             "DescribeLogStreams",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_group_name",
	},
	{
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
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Logs_20140328.GetLogGroupFields",
		},
		Permission:             "GetLogGroupFields",
		IsExtra:                true,
		ExtraComponentBodyKey:  "logGroupName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "log_group_name",
	},
	{
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
	{
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
	{
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
