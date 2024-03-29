package policy

// Microsoft_ApiManagement_apimauthorizationservers policy
var Microsoft_ApiManagement_apimauthorizationservers = map[string]Policy{
	"AuthorizationServer_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/authorizationServers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "AuthorizationServer_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"AuthorizationServer_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/authorizationServers/{{.authsid}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "AuthorizationServer_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"AuthorizationServer_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/authorizationServers/{{.authsid}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "AuthorizationServer_ListSecrets",
		Resource:    "Microsoft.ApiManagement",
	},
}
