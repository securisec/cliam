package policy

    var Microsoft_Cdn_cdn = map[string]Policy{
        "Profiles_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cdn/profiles",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Profiles_List",
    Resource:       "Microsoft.Cdn",
},
"Profiles_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Profiles_ListByResourceGroup",
    Resource:       "Microsoft.Cdn",
},
"Profiles_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Profiles_Get",
    Resource:       "Microsoft.Cdn",
},
"Profiles_GenerateSsoUri": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/generateSsoUri",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Profiles_GenerateSsoUri",
    Resource:       "Microsoft.Cdn",
},
"Profiles_ListSupportedOptimizationTypes": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/getSupportedOptimizationTypes",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Profiles_ListSupportedOptimizationTypes",
    Resource:       "Microsoft.Cdn",
},
"Profiles_ListResourceUsage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/checkResourceUsage",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Profiles_ListResourceUsage",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_ListByProfile": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_ListByProfile",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_Get",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_Start": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/start",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_Start",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_Stop": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/stop",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_Stop",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_PurgeContent": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/purge",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_PurgeContent",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_LoadContent": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/load",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_LoadContent",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_ValidateCustomDomain": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/validateCustomDomain",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_ValidateCustomDomain",
    Resource:       "Microsoft.Cdn",
},
"Endpoints_ListResourceUsage": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/checkResourceUsage",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Endpoints_ListResourceUsage",
    Resource:       "Microsoft.Cdn",
},
"Origins_ListByEndpoint": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/origins",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Origins_ListByEndpoint",
    Resource:       "Microsoft.Cdn",
},
"Origins_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/origins/{{.originName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Origins_Get",
    Resource:       "Microsoft.Cdn",
},
"OriginGroups_ListByEndpoint": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/originGroups",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "OriginGroups_ListByEndpoint",
    Resource:       "Microsoft.Cdn",
},
"OriginGroups_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/originGroups/{{.originGroupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "OriginGroups_Get",
    Resource:       "Microsoft.Cdn",
},
"CustomDomains_ListByEndpoint": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/customDomains",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CustomDomains_ListByEndpoint",
    Resource:       "Microsoft.Cdn",
},
"CustomDomains_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/customDomains/{{.customDomainName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CustomDomains_Get",
    Resource:       "Microsoft.Cdn",
},
"CustomDomains_DisableCustomHttps": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/customDomains/{{.customDomainName}}/disableCustomHttps",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CustomDomains_DisableCustomHttps",
    Resource:       "Microsoft.Cdn",
},
"CustomDomains_EnableCustomHttps": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/profiles/{{.profileName}}/endpoints/{{.endpointName}}/customDomains/{{.customDomainName}}/enableCustomHttps",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CustomDomains_EnableCustomHttps",
    Resource:       "Microsoft.Cdn",
},
"CheckNameAvailability": {
    Path: "/providers/Microsoft.Cdn/checkNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CheckNameAvailability",
    Resource:       "Microsoft.Cdn",
},
"CheckNameAvailabilityWithSubscription": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cdn/checkNameAvailability",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "CheckNameAvailabilityWithSubscription",
    Resource:       "Microsoft.Cdn",
},
"ValidateProbe": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cdn/validateProbe",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ValidateProbe",
    Resource:       "Microsoft.Cdn",
},
"ResourceUsage_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cdn/checkResourceUsage",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ResourceUsage_List",
    Resource:       "Microsoft.Cdn",
},
"Operations_List": {
    Path: "/providers/Microsoft.Cdn/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.Cdn",
},
"EdgeNodes_List": {
    Path: "/providers/Microsoft.Cdn/edgenodes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "EdgeNodes_List",
    Resource:       "Microsoft.Cdn",
},
    }
    