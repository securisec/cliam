package policy

// Microsoft_AzureStackHCI_updateRuns policy
var Microsoft_AzureStackHCI_updateRuns = map[string]Policy{
	"UpdateRuns_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updates/{{.updateName}}/updateRuns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "UpdateRuns_List",
		Resource:    "Microsoft.AzureStackHCI",
	},
	"UpdateRuns_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureStackHCI/clusters/{{.clusterName}}/updates/{{.updateName}}/updateRuns/{{.updateRunName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-12-01",
		},
		OperationID: "UpdateRuns_Get",
		Resource:    "Microsoft.AzureStackHCI",
	},
}
