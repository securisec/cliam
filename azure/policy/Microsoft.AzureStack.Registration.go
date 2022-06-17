package policy

var Microsoft_AzureStack_Registration = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_List",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureStack/registrations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_ListBySubscription",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_Get",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/getactivationkey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_GetActivationKey",
		Resource:    "Microsoft.AzureStack",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/enableRemoteManagement",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2017-06-01",
		},
		OperationID: "Registrations_EnableRemoteManagement",
		Resource:    "Microsoft.AzureStack",
	},
}
