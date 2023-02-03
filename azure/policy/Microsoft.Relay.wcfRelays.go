package policy

// Microsoft_Relay_wcfRelays policy
var Microsoft_Relay_wcfRelays = map[string]Policy{
	"WCFRelays_ListByNamespace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/wcfRelays",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "WCFRelays_ListByNamespace",
		Resource:    "Microsoft.Relay",
	},
	"WCFRelays_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/wcfRelays/{{.relayName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "WCFRelays_Get",
		Resource:    "Microsoft.Relay",
	},
}
