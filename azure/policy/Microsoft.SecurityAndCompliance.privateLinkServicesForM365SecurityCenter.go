package policy

    var Microsoft_SecurityAndCompliance_privateLinkServicesForM365SecurityCenter = map[string]Policy{
        "privateLinkServicesForM365SecurityCenter_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{{.resourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForM365SecurityCenter_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForM365SecurityCenter_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForM365SecurityCenter_List",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"privateLinkServicesForM365SecurityCenter_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "privateLinkServicesForM365SecurityCenter_ListByResourceGroup",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsSec_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{{.resourceName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsSec_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateEndpointConnectionsSec_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{{.resourceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateEndpointConnectionsSec_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResourcesSec_ListByService": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{{.resourceName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResourcesSec_ListByService",
    Resource:       "Microsoft.SecurityAndCompliance",
},
"PrivateLinkResourcesSec_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.SecurityAndCompliance/privateLinkServicesForM365SecurityCenter/{{.resourceName}}/privateLinkResources/{{.groupName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-03-08",
    },
	OperationID:    "PrivateLinkResourcesSec_Get",
    Resource:       "Microsoft.SecurityAndCompliance",
},
    }
    