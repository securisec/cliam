package policy

// Microsoft_ApiManagement_apimtagresources policy
var Microsoft_ApiManagement_apimtagresources = map[string]Policy{
	"TagResource_ListByService": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/tagResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "TagResource_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
}
