package policy

// Microsoft_Synapse_keys policy
var Microsoft_Synapse_keys = map[string]Policy{
	"Keys_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/keys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Keys_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
	"Keys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/keys/{{.keyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Keys_Get",
		Resource:    "Microsoft.Synapse",
	},
}
