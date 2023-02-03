package policy

// Microsoft_ApiManagement_apimtags policy
var Microsoft_ApiManagement_apimtags = map[string]Policy{
	"Tag_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
	"Tag_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tags/{{.tagId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Tag_Get",
		Resource:    "Microsoft.ApiManagement",
	},
}
