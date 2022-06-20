package policy

    var Microsoft_SecurityAndCompliance_privateLinkServicesForEDMUpload = map[string]Policy{
        "privateLinkServicesForEDMUpload_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload/{{.resourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForEDMUpload_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForEDMUpload_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForEDMUpload_List",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForEDMUpload_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForEDMUpload_ListByResourceGroup",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsForEDM_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload/{{.resourceName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsForEDM_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsForEDM_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsForEDM_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResources_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload/{{.resourceName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResources_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResources_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForEDMUpload/{{.resourceName}}/privateLinkResources/{{.groupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResources_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
    }
    