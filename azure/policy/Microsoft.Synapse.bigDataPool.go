package policy

// Microsoft_Synapse_bigDataPool policy
var Microsoft_Synapse_bigDataPool = map[string]Policy{
	"BigDataPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/bigDataPools/{{.bigDataPoolName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "BigDataPools_Get",
		Resource:    "Microsoft.Synapse",
	},
	"BigDataPools_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/bigDataPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "BigDataPools_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
}
