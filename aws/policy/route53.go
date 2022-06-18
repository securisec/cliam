package policy

var Route53Policies = map[string]Service{
	"ListGeoLocations": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/geolocations",
		Permission:    "ListGeoLocations",
	},
	"GetHealthCheckCount": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/healthcheckcount",
		Permission:    "GetHealthCheckCount",
	},
	"GetHostedZoneCount": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/hostedzonecount",
		Permission:    "GetHostedZoneCount",
	},
	"GetTrafficPolicyInstanceCount": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/trafficpolicyinstancecount",
		Permission:    "GetTrafficPolicyInstanceCount",
	},
	"ListHealthChecks": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/healthcheck",
		Permission:    "ListHealthChecks",
	},
	"ListHostedZones": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/hostedzone",
		Permission:    "ListHostedZones",
	},
	"ListHostedZonesByName": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/hostedzonesbyname",
		Permission:    "ListHostedZonesByName",
	},
	"ListQueryLoggingConfigs": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/queryloggingconfig",
		Permission:    "ListQueryLoggingConfigs",
	},
	"ListReusableDelegationSets": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/delegationset",
		Permission:    "ListReusableDelegationSets",
	},
	"ListTrafficPolicies": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/trafficpolicies",
		Permission:    "ListTrafficPolicies",
	},
	"ListTrafficPolicyInstances": {
		IgnoreRegion:  true,
		ServiceSuffix: "2013-04-01/trafficpolicyinstances",
		Permission:    "ListTrafficPolicyInstances",
	},

	// extra policies
	"GetAccountLimit": {
		ServiceSuffix:          "/2013-04-01/accountlimit/{{.type}}",
		Permission:             "GetAccountLimit",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "type",
	},
	"GetChange": {
		ServiceSuffix:          "/2013-04-01/change/{{.id}}",
		Permission:             "GetChange",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetDNSSEC": {
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.hosted_zone_id}}/dnssec",
		Permission:             "GetDNSSEC",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
	"GetHealthCheck": {
		ServiceSuffix:          "/2013-04-01/healthcheck/{{.health_check_id}}",
		Permission:             "GetHealthCheck",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "health_check_id",
	},
	"GetHealthCheckLastFailureReason": {
		ServiceSuffix:          "/2013-04-01/healthcheck/{{.health_check_id}}/lastfailurereason",
		Permission:             "GetHealthCheckLastFailureReason",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "health_check_id",
	},
	"GetHealthCheckStatus": {
		ServiceSuffix:          "/2013-04-01/healthcheck/{{.health_check_id}}/status",
		Permission:             "GetHealthCheckStatus",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "health_check_id",
	},
	"GetHostedZone": {
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.id}}",
		Permission:             "GetHostedZone",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetQueryLoggingConfig": {
		ServiceSuffix:          "/2013-04-01/queryloggingconfig/{{.id}}",
		Permission:             "GetQueryLoggingConfig",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetReusableDelegationSet": {
		ServiceSuffix:          "/2013-04-01/delegationset/{{.id}}",
		Permission:             "GetReusableDelegationSet",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"GetTrafficPolicyInstance": {
		ServiceSuffix:          "/2013-04-01/trafficpolicyinstance/{{.id}}",
		Permission:             "GetTrafficPolicyInstance",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListResourceRecordSets": {
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.hosted_zone_id}}/rrset",
		Permission:             "ListResourceRecordSets",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
	"ListTrafficPolicyInstancesByHostedZone": {
		ServiceSuffix:          "/2013-04-01/trafficpolicyinstances/hostedzone",
		Permission:             "ListTrafficPolicyInstancesByHostedZone",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
	"ListTrafficPolicyVersions": {
		ServiceSuffix:          "/2013-04-01/trafficpolicies/{{.id}}/versions",
		Permission:             "ListTrafficPolicyVersions",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "id",
	},
	"ListVPCAssociationAuthorizations": {
		ServiceSuffix:          "/2013-04-01/hostedzone/{{.hosted_zone_id}}/authorizevpcassociation",
		Permission:             "ListVPCAssociationAuthorizations",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "hosted_zone_id",
	},
}
