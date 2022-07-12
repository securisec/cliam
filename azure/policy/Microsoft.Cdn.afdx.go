package policy

    var Microsoft_Cdn_afdx = map[string]Policy{
        "CheckEndpointNameAvailability": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/checkEndpointNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CheckEndpointNameAvailability",
    Resource:       "Microsoft.Cdn",
},
"AFDProfiles_ListResourceUsage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/usages",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDProfiles_ListResourceUsage",
    Resource:       "Microsoft.Cdn",
},
"AFDProfiles_CheckHostNameAvailability": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/checkHostNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDProfiles_CheckHostNameAvailability",
    Resource:       "Microsoft.Cdn",
},
"AFDCustomDomains_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/customDomains",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDCustomDomains_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"AFDCustomDomains_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/customDomains/{{.customDomainName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDCustomDomains_Get",
    Resource:       "Microsoft.Cdn",
},
"AFDCustomDomains_RefreshValidationToken": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/customDomains/{{.customDomainName}}/refreshValidationToken",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDCustomDomains_RefreshValidationToken",
    Resource:       "Microsoft.Cdn",
},
"AFDEndpoints_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDEndpoints_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"AFDEndpoints_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints/{{.endpointName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDEndpoints_Get",
    Resource:       "Microsoft.Cdn",
},
"AFDEndpoints_PurgeContent": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints/{{.endpointName}}/purge",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDEndpoints_PurgeContent",
    Resource:       "Microsoft.Cdn",
},
"AFDEndpoints_ListResourceUsage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints/{{.endpointName}}/usages",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDEndpoints_ListResourceUsage",
    Resource:       "Microsoft.Cdn",
},
"AFDEndpoints_ValidateCustomDomain": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints/{{.endpointName}}/validateCustomDomain",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDEndpoints_ValidateCustomDomain",
    Resource:       "Microsoft.Cdn",
},
"AFDOriginGroups_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/originGroups",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDOriginGroups_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"AFDOriginGroups_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/originGroups/{{.originGroupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDOriginGroups_Get",
    Resource:       "Microsoft.Cdn",
},
"AFDOriginGroups_ListResourceUsage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/originGroups/{{.originGroupName}}/usages",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDOriginGroups_ListResourceUsage",
    Resource:       "Microsoft.Cdn",
},
"AFDOrigins_ListByOriginGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/originGroups/{{.originGroupName}}/origins",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDOrigins_ListByOriginGroup",
    Resource:       "Microsoft.Cdn",
},
"AFDOrigins_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/originGroups/{{.originGroupName}}/origins/{{.originName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "AFDOrigins_Get",
    Resource:       "Microsoft.Cdn",
},
"Routes_ListByEndpoint": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints/{{.endpointName}}/routes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Routes_ListByEndpoint",
    Resource:       "Microsoft.Cdn",
},
"Routes_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/afdEndpoints/{{.endpointName}}/routes/{{.routeName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Routes_Get",
    Resource:       "Microsoft.Cdn",
},
"RuleSets_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/ruleSets",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "RuleSets_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"RuleSets_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/ruleSets/{{.ruleSetName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "RuleSets_Get",
    Resource:       "Microsoft.Cdn",
},
"RuleSets_ListResourceUsage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/ruleSets/{{.ruleSetName}}/usages",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "RuleSets_ListResourceUsage",
    Resource:       "Microsoft.Cdn",
},
"Rules_ListByRuleSet": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/ruleSets/{{.ruleSetName}}/rules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Rules_ListByRuleSet",
    Resource:       "Microsoft.Cdn",
},
"Rules_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/ruleSets/{{.ruleSetName}}/rules/{{.ruleName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Rules_Get",
    Resource:       "Microsoft.Cdn",
},
"SecurityPolicies_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/securityPolicies",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "SecurityPolicies_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"SecurityPolicies_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/securityPolicies/{{.securityPolicyName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "SecurityPolicies_Get",
    Resource:       "Microsoft.Cdn",
},
"Secrets_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/secrets",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Secrets_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"Secrets_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/secrets/{{.secretName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Secrets_Get",
    Resource:       "Microsoft.Cdn",
},
"Validate_Secret": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cdn/validateSecret",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Validate_Secret",
    Resource:       "Microsoft.Cdn",
},
"LogAnalytics_GetLogAnalyticsMetrics": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getLogAnalyticsMetrics",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "LogAnalytics_GetLogAnalyticsMetrics",
    Resource:       "Microsoft.Cdn",
},
"LogAnalytics_GetLogAnalyticsRankings": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getLogAnalyticsRankings",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "LogAnalytics_GetLogAnalyticsRankings",
    Resource:       "Microsoft.Cdn",
},
"LogAnalytics_GetLogAnalyticsLocations": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getLogAnalyticsLocations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "LogAnalytics_GetLogAnalyticsLocations",
    Resource:       "Microsoft.Cdn",
},
"LogAnalytics_GetLogAnalyticsResources": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getLogAnalyticsResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "LogAnalytics_GetLogAnalyticsResources",
    Resource:       "Microsoft.Cdn",
},
"LogAnalytics_GetWafLogAnalyticsMetrics": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getWafLogAnalyticsMetrics",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "LogAnalytics_GetWafLogAnalyticsMetrics",
    Resource:       "Microsoft.Cdn",
},
"LogAnalytics_GetWafLogAnalyticsRankings": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getWafLogAnalyticsRankings",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "LogAnalytics_GetWafLogAnalyticsRankings",
    Resource:       "Microsoft.Cdn",
},
    }
    