package policy

// Microsoft_ApiManagement_apimapiversionsets policy
var Microsoft_ApiManagement_apimapiversionsets = map[string]Policy{
	"ApiVersionSet_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apiVersionSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiVersionSet_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"ApiVersionSet_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apiVersionSets/{{.versionSetId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "ApiVersionSet_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
