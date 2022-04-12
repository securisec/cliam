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
}
