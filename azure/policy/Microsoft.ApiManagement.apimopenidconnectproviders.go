package policy

var Microsoft_ApiManagement_apimopenidconnectproviders = map[string]Policy{
	"OpenIdConnectProvider_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/openidConnectProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "OpenIdConnectProvider_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"OpenIdConnectProvider_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/openidConnectProviders/{{.opid}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "OpenIdConnectProvider_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"OpenIdConnectProvider_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/openidConnectProviders/{{.opid}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "OpenIdConnectProvider_ListSecrets",
		Resource:    "Microsoft.ApiManagement",
	},
}
