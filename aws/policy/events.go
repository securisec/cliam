package policy

import "github.com/securisec/cliam/shared"

// EventsPolicies policy
var EventsPolicies = map[string]Service{
	"DescribeEventBus": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeEventBus",
		},
		Permission: "DescribeEventBus",
	},
	"ListApiDestinations": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListApiDestinations",
		},
		Permission: "ListApiDestinations",
	},
	"ListArchives": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListArchives",
		},
		Permission: "ListArchives",
	},
	"ListConnections": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListConnections",
		},
		Permission: "ListConnections",
	},
	"ListEventBuses": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListEventBuses",
		},
		Permission: "ListEventBuses",
	},
	"ListEventSources": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListEventSources",
		},
		Permission: "ListEventSources",
	},
	"ListReplays": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListReplays",
		},
		Permission: "ListReplays",
	},
	"ListRules": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListRules",
		},
		Permission: "ListRules",
	},
	"PutPermission": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.PutPermission",
		},
		Permission: "PutPermission",
	},
	"RemovePermission": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.RemovePermission",
		},
		Permission: "RemovePermission",
	},

	// extra
	"DescribeApiDestination": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeApiDestination",
		},
		Permission:             "DescribeApiDestination",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"DescribeArchive": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeArchive",
		},
		Permission:             "DescribeArchive",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ArchiveName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "archive_name",
	},
	"DescribeConnection": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeConnection",
		},
		Permission:             "DescribeConnection",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"DescribeEventSource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeEventSource",
		},
		Permission:             "DescribeEventSource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"DescribePartnerEventSource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribePartnerEventSource",
		},
		Permission:             "DescribePartnerEventSource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"DescribeReplay": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeReplay",
		},
		Permission:             "DescribeReplay",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ReplayName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "replay_name",
	},
	"DescribeRule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.DescribeRule",
		},
		Permission:             "DescribeRule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"ListPartnerEventSourceAccounts": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListPartnerEventSourceAccounts",
		},
		Permission:             "ListPartnerEventSourceAccounts",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EventSourceName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "event_source_name",
	},
	"ListPartnerEventSources": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListPartnerEventSources",
		},
		Permission:             "ListPartnerEventSources",
		IsExtra:                true,
		ExtraComponentBodyKey:  "NamePrefix",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name_prefix",
	},
	"ListRuleNamesByTarget": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListRuleNamesByTarget",
		},
		Permission:             "ListRuleNamesByTarget",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TargetArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "target_arn",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	"ListTargetsByRule": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSEvents.ListTargetsByRule",
		},
		Permission:             "ListTargetsByRule",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Rule",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "rule",
	},
}
