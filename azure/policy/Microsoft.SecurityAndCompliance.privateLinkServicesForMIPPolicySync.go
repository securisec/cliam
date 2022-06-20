package policy

    var Microsoft_SecurityAndCompliance_privateLinkServicesForMIPPolicySync = map[string]Policy{
        "privateLinkServicesForMIPPolicySync_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{{.resourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForMIPPolicySync_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForMIPPolicySync_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForMIPPolicySync_List",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForMIPPolicySync_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForMIPPolicySync_ListByResourceGroup",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsForMIPPolicySync_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{{.resourceName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsForMIPPolicySync_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsForMIPPolicySync_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsForMIPPolicySync_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResourcesForMIPPolicySync_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{{.resourceName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResourcesForMIPPolicySync_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResourcesForMIPPolicySync_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForMIPPolicySync/{{.resourceName}}/privateLinkResources/{{.groupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResourcesForMIPPolicySync_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
    }
    