package policy

import "github.com/securisec/cliam/shared"

var SMSPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetApp",
		},
		Permission: "GetApp",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetAppLaunchConfiguration",
		},
		Permission: "GetAppLaunchConfiguration",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetAppReplicationConfiguration",
		},
		Permission: "GetAppReplicationConfiguration",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetConnectors",
		},
		Permission: "GetConnectors",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetReplicationJobs",
		},
		Permission: "GetReplicationJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetServers",
		},
		Permission: "GetServers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.ListApps",
		},
		Permission: "ListApps",
	},
}
