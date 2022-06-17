package policy

var Microsoft_ApiManagement_apimprivatelink = map[string]Policy{
	"PrivateEndpointConnection_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PrivateEndpointConnection_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"PrivateEndpointConnection_GetByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PrivateEndpointConnection_GetByName",
		Resource:    "Microsoft.ApiManagement",
	},
	"PrivateEndpointConnection_ListPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PrivateEndpointConnection_ListPrivateLinkResources",
		Resource:    "Microsoft.ApiManagement",
	},
	"PrivateEndpointConnection_GetPrivateLinkResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/privateLinkResources/{{.privateLinkSubResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PrivateEndpointConnection_GetPrivateLinkResource",
		Resource:    "Microsoft.ApiManagement",
	},
}
