package policy

var Route53Policies = []Service{
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/geolocations",
		Permission:    "ListGeoLocations",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/healthcheckcount",
		Permission:    "GetHealthCheckCount",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/hostedzonecount",
		Permission:    "GetHostedZoneCount",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/trafficpolicyinstancecount",
		Permission:    "GetTrafficPolicyInstanceCount",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/healthcheck",
		Permission:    "ListHealthChecks",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/hostedzone",
		Permission:    "ListHostedZones",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/hostedzonesbyname",
		Permission:    "ListHostedZonesByName",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/queryloggingconfig",
		Permission:    "ListQueryLoggingConfigs",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/delegationset",
		Permission:    "ListReusableDelegationSets",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/trafficpolicies",
		Permission:    "ListTrafficPolicies",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/trafficpolicyinstances",
		Permission:    "ListTrafficPolicyInstances",
	},

	// extra policies
	{
		ServiceSuffix:          "/2013-04-01/accountlimit/{{.type}}",
		Permission:             "GetAccountLimit",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "type",
	},
	{
		ServiceSuffix:          "/2013-04-01/change/{{.id}}",
		Permission:             "GetChange",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.hosted_zone_id}}/dnssec",
		Permission:             "GetDNSSEC",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
	{
		ServiceSuffix:          "/2013-04-01/healthcheck/{{.health_check_id}}",
		Permission:             "GetHealthCheck",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "health_check_id",
	},
	{
		ServiceSuffix:          "/2013-04-01/healthcheck/{{.health_check_id}}/lastfailurereason",
		Permission:             "GetHealthCheckLastFailureReason",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "health_check_id",
	},
	{
		ServiceSuffix:          "/2013-04-01/healthcheck/{{.health_check_id}}/status",
		Permission:             "GetHealthCheckStatus",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "health_check_id",
	},
	{
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.id}}",
		Permission:             "GetHostedZone",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2013-04-01/queryloggingconfig/{{.id}}",
		Permission:             "GetQueryLoggingConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2013-04-01/delegationset/{{.id}}",
		Permission:             "GetReusableDelegationSet",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2013-04-01/trafficpolicyinstance/{{.id}}",
		Permission:             "GetTrafficPolicyInstance",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.hosted_zone_id}}/rrset",
		Permission:             "ListResourceRecordSets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
	{
		ServiceSuffix:          "/2013-04-01/trafficpolicyinstances/hostedzone",
		Permission:             "ListTrafficPolicyInstancesByHostedZone",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
	{
		ServiceSuffix:          "/2013-04-01/trafficpolicies/{{.id}}/versions",
		Permission:             "ListTrafficPolicyVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	{
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.hosted_zone_id}}/authorizevpcassociation",
		Permission:             "ListVPCAssociationAuthorizations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
}
