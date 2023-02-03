package policy

// Microsoft_AzureStackHCI_publishers policy
var Microsoft_AzureStackHCI_publishers = map[string]Policy{
	"Publishers_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/publishers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Publishers_ListByCluster",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Publishers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/publishers/{{.publisherName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Publishers_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
