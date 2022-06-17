package policy

var Microsoft_AzureStack_CustomerSubscription = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/customerSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "CustomerSubscriptions_List",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/customerSubscriptions/{{.customerSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "CustomerSubscriptions_Get",
		Resource:    "Microsoft.AzureStack",
	},
}
