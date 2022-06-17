package policy

var Microsoft_ApiManagement_apimproductsByTags = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/productsByTags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "Product_ListByTags",
		Resource:    "Microsoft.ApiManagement",
	},
}