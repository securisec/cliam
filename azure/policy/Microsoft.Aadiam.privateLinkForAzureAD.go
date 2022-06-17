package policy

var Microsoft_Aadiam_privateLinkForAzureAD = map[string]Policy{
	"privateLinkForAzureAd_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd/{{.policyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "privateLinkForAzureAd_Get",
		Resource:    "Microsoft.Aadiam",
	},
	"privateLinkForAzureAd_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/microsoft.aadiam/privateLinkForAzureAd",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "privateLinkForAzureAd_ListBySubscription",
		Resource:    "Microsoft.Aadiam",
	},
	"privateLinkForAzureAd_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "privateLinkForAzureAd_List",
		Resource:    "Microsoft.Aadiam",
	},
}
