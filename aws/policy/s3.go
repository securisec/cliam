package policy

var S3Policies = []Service{
	{
		ServiceSuffix: "",
		Permission:    "ListBuckets",
	},

	// extra
	{
		ServiceSuffix:          "/{{.}}?accelerate",
		Permission:             "GetBucketAccelerateConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?acl",
		Permission:             "GetBucketAcl",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?cors",
		Permission:             "GetBucketCors",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?encryption",
		Permission:             "GetBucketEncryption",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?lifecycle",
		Permission:             "GetBucketLifecycle",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?lifecycle",
		Permission:             "GetBucketLifecycleConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?location",
		Permission:             "GetBucketLocation",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?logging",
		Permission:             "GetBucketLogging",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?notification",
		Permission:             "GetBucketNotification",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?notification",
		Permission:             "GetBucketNotificationConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?ownershipControls",
		Permission:             "GetBucketOwnershipControls",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?policy",
		Permission:             "GetBucketPolicy",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?policyStatus",
		Permission:             "GetBucketPolicyStatus",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?replication",
		Permission:             "GetBucketReplication",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?requestPayment",
		Permission:             "GetBucketRequestPayment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?tagging",
		Permission:             "GetBucketTagging",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?versioning",
		Permission:             "GetBucketVersioning",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?website",
		Permission:             "GetBucketWebsite",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?object-lock",
		Permission:             "GetObjectLockConfiguration",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?publicAccessBlock",
		Permission:             "GetPublicAccessBlock",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}",
		Permission:             "HeadBucket",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?analytics",
		Permission:             "ListBucketAnalyticsConfigurations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?intelligent-tiering",
		Permission:             "ListBucketIntelligentTieringConfigurations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?inventory",
		Permission:             "ListBucketInventoryConfigurations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?metrics",
		Permission:             "ListBucketMetricsConfigurations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?uploads",
		Permission:             "ListMultipartUploads",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?versions",
		Permission:             "ListObjectVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}",
		Permission:             "ListObjects",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
	{
		ServiceSuffix:          "/{{.}}?list-type=2",
		Permission:             "ListObjectsV2",
		ExtraComponentLocation: "path",
		IsExtra:                true,
	},
}
