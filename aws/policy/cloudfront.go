package policy

var CloudfrontPolicies = []Service{
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2018-11-05/distribution",
		Permission:    "ListDistributions",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/cache-policy",
		Permission:    "ListCachePolicies",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/function",
		Permission:    "ListFunctions",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/key-group",
		Permission:    "ListKeyGroups",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/origin-request-policy",
		Permission:    "ListOriginRequestPolicies",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/realtime-log-config",
		Permission:    "ListRealtimeLogConfigs",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/response-headers-policy",
		Permission:    "ListResponseHeadersPolicies",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/origin-access-identity/cloudfront",
		Permission:    "ListCloudFrontOriginAccessIdentities",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/field-level-encryption",
		Permission:    "ListFieldLevelEncryptionConfigs",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/field-level-encryption-profile",
		Permission:    "ListFieldLevelEncryptionProfiles",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2020-05-31/streaming-distribution",
		Permission:    "ListStreamingDistributions",
	},
}
