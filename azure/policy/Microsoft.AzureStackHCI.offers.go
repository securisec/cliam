package policy

// Microsoft_AzureStackHCI_offers policy
var Microsoft_AzureStackHCI_offers = map[string]Policy{
	"Offers_ListByPublisher": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/publishers/{{.publisherName}}/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Offers_ListByPublisher",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Offers_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/offers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Offers_ListByCluster",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Offers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/publishers/{{.publisherName}}/offers/{{.offerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Offers_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
