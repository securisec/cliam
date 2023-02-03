package policy

// Microsoft_AzureStackHCI_skus policy
var Microsoft_AzureStackHCI_skus = map[string]Policy{
	"Skus_ListByOffer": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/publishers/{{.publisherName}}/offers/{{.offerName}}/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Skus_ListByOffer",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Skus_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/publishers/{{.publisherName}}/offers/{{.offerName}}/skus/{{.skuName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Skus_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
