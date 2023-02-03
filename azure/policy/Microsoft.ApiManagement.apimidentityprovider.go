package policy

// Microsoft_ApiManagement_apimidentityprovider policy
var Microsoft_ApiManagement_apimidentityprovider = map[string]Policy{
	"IdentityProvider_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/identityProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "IdentityProvider_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"IdentityProvider_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/identityProviders/{{.identityProviderName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "IdentityProvider_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"IdentityProvider_ListSecrets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/identityProviders/{{.identityProviderName}}/listSecrets",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "IdentityProvider_ListSecrets",
		Resource:    "Microsoft.ApiManagement",
	},
}
