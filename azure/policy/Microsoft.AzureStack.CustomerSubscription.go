package policy

// Microsoft_AzureStack_CustomerSubscription policy
var Microsoft_AzureStack_CustomerSubscription = map[string]Policy{
	"CustomerSubscriptions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/customerSubscriptions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CustomerSubscriptions_List",
		Resource:    "Microsoft.AzureStack",
	},
	"CustomerSubscriptions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroup}}/providers/Microsoft.AzureStack/registrations/{{.registrationName}}/customerSubscriptions/{{.customerSubscriptionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-06-01",
		},
		OperationID: "CustomerSubscriptions_Get",
		Resource:    "Microsoft.AzureStack",
	},
}
