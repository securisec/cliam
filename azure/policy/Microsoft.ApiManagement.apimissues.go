package policy

var Microsoft_ApiManagement_apimissues = map[string]Policy{
	"Issue_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/issues",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Issue_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Issue_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/issues/{{.issueId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Issue_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
