package policy

import "github.com/securisec/cliam/shared"

var SMSPolicies = map[string]Service{
	"GetApp": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetApp",
		},
		Permission: "GetApp",
	},
	"GetAppLaunchConfiguration": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetAppLaunchConfiguration",
		},
		Permission: "GetAppLaunchConfiguration",
	},
	"GetAppReplicationConfiguration": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetAppReplicationConfiguration",
		},
		Permission: "GetAppReplicationConfiguration",
	},
	"GetConnectors": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetConnectors",
		},
		Permission: "GetConnectors",
	},
	"GetReplicationJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetReplicationJobs",
		},
		Permission: "GetReplicationJobs",
	},
	"GetServers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetServers",
		},
		Permission: "GetServers",
	},
	"ListApps": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.ListApps",
		},
		Permission: "ListApps",
	},

	// extra
	"GetAppValidationConfiguration": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetAppValidationConfiguration",
		},
		Permission:             "GetAppValidationConfiguration",
		IsExtra:                true,
		ExtraComponentBodyKey:  "appId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "app_id",
	},
	"GetAppValidationOutput": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetAppValidationOutput",
		},
		Permission:             "GetAppValidationOutput",
		IsExtra:                true,
		ExtraComponentBodyKey:  "appId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "app_id",
	},
	"GetReplicationRuns": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSServerMigrationService_V2016_10_24.GetReplicationRuns",
		},
		Permission:             "GetReplicationRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "replicationJobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "replication_job_id",
	},
}
