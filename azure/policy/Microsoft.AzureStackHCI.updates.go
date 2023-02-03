package policy

// Microsoft_AzureStackHCI_updates policy
var Microsoft_AzureStackHCI_updates = map[string]Policy{
	"Updates_Post": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updates/{{.updateName}}/apply",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Updates_Post",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Updates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Updates_List",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"Updates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updates/{{.updateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "Updates_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
