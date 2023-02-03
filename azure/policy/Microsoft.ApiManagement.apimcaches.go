package policy

// Microsoft_ApiManagement_apimcaches policy
var Microsoft_ApiManagement_apimcaches = map[string]Policy{
	"Cache_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/caches",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Cache_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Cache_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/caches/{{.cacheId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Cache_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
