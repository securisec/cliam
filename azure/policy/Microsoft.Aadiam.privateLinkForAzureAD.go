package policy

var Microsoft_Aadiam_privateLinkForAzureAD = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd/{{.policyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "privateLinkForAzureAd_Get",
		Resource:    "Microsoft.Aadiam",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/microsoft.aadiam/privateLinkForAzureAd",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "privateLinkForAzureAd_ListBySubscription",
		Resource:    "Microsoft.Aadiam",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "privateLinkForAzureAd_List",
		Resource:    "Microsoft.Aadiam",
	},
}
