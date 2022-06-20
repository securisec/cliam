package policy

    var Microsoft_HybridCompute_privateLinkScopes = map[string]Policy{
        "PrivateLinkScopes_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridCompute/privateLinkScopes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkScopes_List",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateLinkScopes_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/privateLinkScopes",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkScopes_ListByResourceGroup",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateLinkScopes_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/privateLinkScopes/{{.scopeName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkScopes_Get",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateLinkResources_ListByPrivateLinkScope": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/privateLinkScopes/{{.scopeName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkResources_ListByPrivateLinkScope",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateLinkResources_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/privateLinkScopes/{{.scopeName}}/privateLinkResources/{{.groupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkResources_Get",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateEndpointConnections_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/privateLinkScopes/{{.scopeName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateEndpointConnections_Get",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateEndpointConnections_ListByPrivateLinkScope": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/privateLinkScopes/{{.scopeName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateEndpointConnections_ListByPrivateLinkScope",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateLinkScopes_GetValidationDetails": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridCompute/locations/{{.location}}/privateLinkScopes/{{.privateLinkScopeId}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkScopes_GetValidationDetails",
    Resource:       "Microsoft.HybridCompute",
},
"PrivateLinkScopes_GetValidationDetailsForMachine": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/privateLinkScopes/current",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2022-03-10",
    },
	OperationID:    "PrivateLinkScopes_GetValidationDetailsForMachine",
    Resource:       "Microsoft.HybridCompute",
},
    }
    