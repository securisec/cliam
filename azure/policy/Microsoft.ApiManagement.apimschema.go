package policy

var Microsoft_ApiManagement_apimschema = map[string]Policy{
	"GlobalSchema_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/schemas",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GlobalSchema_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"GlobalSchema_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/schemas/{{.schemaId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "GlobalSchema_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
