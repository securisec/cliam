package policy

// Microsoft_AzureStackHCI_clusters policy
var Microsoft_AzureStackHCI_clusters = map[string]Policy{
	"Clusters_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureStackHCI/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Clusters_ListBySubscription",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Clusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Clusters_ListByResourceGroup",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Clusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Clusters_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Clusters_UploadCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/uploadCertificate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Clusters_UploadCertificate",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Clusters_CreateIdentity": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/createClusterIdentity",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Clusters_CreateIdentity",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Clusters_ExtendSoftwareAssuranceBenefit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/extendSoftwareAssuranceBenefit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Clusters_ExtendSoftwareAssuranceBenefit",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
