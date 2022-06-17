package policy

var Microsoft_Aadiam_privateLinkResources = map[string]Policy{
	"PrivateLinkResources_ListByPrivateLinkPolicy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd/{{.policyName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "PrivateLinkResources_ListByPrivateLinkPolicy",
		Resource:    "Microsoft.Aadiam",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd/{{.policyName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.Aadiam",
	},
}
