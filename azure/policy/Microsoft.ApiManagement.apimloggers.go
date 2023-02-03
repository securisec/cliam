package policy

// Microsoft_ApiManagement_apimloggers policy
var Microsoft_ApiManagement_apimloggers = map[string]Policy{
	"Logger_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/loggers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Logger_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Logger_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/loggers/{{.loggerId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Logger_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
