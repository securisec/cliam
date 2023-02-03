package policy

// Microsoft_Synapse_library policy
var Microsoft_Synapse_library = map[string]Policy{
	"Library_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/libraries/{{.libraryName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Library_Get",
		Resource:    "Microsoft.Synapse",
	},
	"Libraries_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/libraries",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Libraries_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
}
