package policy

    var Microsoft_SecurityAndCompliance_privateLinkServicesForO365ManagementActivityAPI = map[string]Policy{
        "privateLinkServicesForO365ManagementActivityAPI_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{{.resourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForO365ManagementActivityAPI_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForO365ManagementActivityAPI_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForO365ManagementActivityAPI_List",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForO365ManagementActivityAPI_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForO365ManagementActivityAPI_ListByResourceGroup",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsAdtAPI_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{{.resourceName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsAdtAPI_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsAdtAPI_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsAdtAPI_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResourcesAdtAPI_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{{.resourceName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResourcesAdtAPI_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResourcesAdtAPI_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForO365ManagementActivityAPI/{{.resourceName}}/privateLinkResources/{{.groupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResourcesAdtAPI_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
    }
    