package policy

// Microsoft_ApiManagement_apimemailtemplates policy
var Microsoft_ApiManagement_apimemailtemplates = map[string]Policy{
	"EmailTemplate_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/templates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "EmailTemplate_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"EmailTemplate_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/templates/{{.templateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "EmailTemplate_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
