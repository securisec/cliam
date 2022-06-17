package policy

var Microsoft_ServiceBus_networksets = map[string]Policy{
	"Namespaces_GetNetworkRuleSet": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/networkRuleSets/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_GetNetworkRuleSet",
		Resource:    "Microsoft.ServiceBus",
	},
	"Namespaces_ListNetworkRuleSets": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/networkRuleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListNetworkRuleSets",
		Resource:    "Microsoft.ServiceBus",
	},
}
