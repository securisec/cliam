package policy

var Microsoft_ApiManagement_apimpolicydescriptions = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ApiManagement/service/{{.serviceName}}/policyDescriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "PolicyDescription_ListByService",
		Resource:    "Microsoft.ApiManagement",
	},
}
