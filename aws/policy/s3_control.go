package policy

// S3ControlPolicies policy
var S3ControlPolicies = map[string]Service{

	// extra
	"GetPublicAccessBlock": {
		ServiceSuffix:          "/v20180820/configuration/publicAccessBlock",
		Permission:             "GetPublicAccessBlock",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
	"ListAccessPoints": {
		ServiceSuffix:          "/v20180820/accesspoint",
		Permission:             "ListAccessPoints",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
	"ListAccessPointsForObjectLambda": {
		ServiceSuffix:          "/v20180820/accesspointforobjectlambda",
		Permission:             "ListAccessPointsForObjectLambda",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
	"ListJobs": {
		ServiceSuffix:          "/v20180820/jobs",
		Permission:             "ListJobs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
	"ListMultiRegionAccessPoints": {
		ServiceSuffix:          "/v20180820/mrap/instances",
		Permission:             "ListMultiRegionAccessPoints",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
	"ListRegionalBuckets": {
		ServiceSuffix:          "/v20180820/bucket",
		Permission:             "ListRegionalBuckets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
	"ListStorageLensConfigurations": {
		ServiceSuffix:          "/v20180820/storagelens",
		Permission:             "ListStorageLensConfigurations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "account_id",
	},
}
