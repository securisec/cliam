package policy

// Microsoft_AzureStackHCI_updateSummaries policy
var Microsoft_AzureStackHCI_updateSummaries = map[string]Policy{
	"UpdateSummaries_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updateSummaries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "UpdateSummaries_List",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"UpdateSummaries_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updateSummaries/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "UpdateSummaries_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
