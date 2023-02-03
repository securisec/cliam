package policy

// Microsoft_ApiManagement_apimdiagnostics policy
var Microsoft_ApiManagement_apimdiagnostics = map[string]Policy{
	"Diagnostic_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/diagnostics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Diagnostic_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Diagnostic_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/diagnostics/{{.diagnosticId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Diagnostic_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
