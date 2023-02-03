package policy

// Microsoft_Network_networkManagerScopeConnection policy
var Microsoft_Network_networkManagerScopeConnection = map[string]Policy{
	"ScopeConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/scopeConnections/{{.scopeConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ScopeConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"ScopeConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/scopeConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ScopeConnections_List",
		Resource:    "Microsoft.Network",
	},
}
