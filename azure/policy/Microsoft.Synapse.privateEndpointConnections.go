package policy

// Microsoft_Synapse_privateEndpointConnections policy
var Microsoft_Synapse_privateEndpointConnections = map[string]Policy{
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Synapse",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Synapse",
	},
}
