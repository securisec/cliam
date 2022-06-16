package policy

var Microsoft_ApiManagement_apimprivatelink = []Policy{
    {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateEndpointConnections",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "PrivateEndpointConnection_ListByService",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "PrivateEndpointConnection_GetByName",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateLinkResources",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "PrivateEndpointConnection_ListPrivateLinkResources",
    Resource:       "Microsoft.ApiManagement",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateLinkResources/{{.privateLinkSubResourceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-08-01",
    },
	OperationID:    "PrivateEndpointConnection_GetPrivateLinkResource",
    Resource:       "Microsoft.ApiManagement",
},
}
