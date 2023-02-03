package policy

// Microsoft_ApiManagement_apimbackends policy
var Microsoft_ApiManagement_apimbackends = map[string]Policy{
	"Backend_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/backends",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Backend_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Backend_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/backends/{{.backendId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Backend_Get",
		Resource:    "Microsoft.ApiManagement",
	},
	"Backend_Reconnect": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/backends/{{.backendId}}/reconnect",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Backend_Reconnect",
		Resource:    "Microsoft.ApiManagement",
	},
}
