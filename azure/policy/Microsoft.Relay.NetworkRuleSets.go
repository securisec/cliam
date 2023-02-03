package policy

// Microsoft_Relay_NetworkRuleSets policy
var Microsoft_Relay_NetworkRuleSets = map[string]Policy{
	"Namespaces_GetNetworkRuleSet": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Relay/namespaces/{{.namespaceName}}/networkRuleSets/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_GetNetworkRuleSet",
		Resource:    "Microsoft.Relay",
	},
}
