package policy

var Microsoft_ApiManagement_apimapisByTags = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/apisByTags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Api_ListByTags",
		Resource:    "Microsoft.ApiManagement",
	},
}
