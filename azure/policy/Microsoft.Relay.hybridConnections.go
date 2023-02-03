package policy

// Microsoft_Relay_hybridConnections policy
var Microsoft_Relay_hybridConnections = map[string]Policy{
	"HybridConnections_ListByNamespace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/hybridConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "HybridConnections_ListByNamespace",
		Resource:    "Microsoft.Relay",
	},
	"HybridConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/hybridConnections/{{.hybridConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "HybridConnections_Get",
		Resource:    "Microsoft.Relay",
	},
}
