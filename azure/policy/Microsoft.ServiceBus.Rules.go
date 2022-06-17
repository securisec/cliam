package policy

var Microsoft_ServiceBus_Rules = map[string]Policy{
	"Rules_ListBySubscriptions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/subscriptions/{{.subscriptionName}}/rules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Rules_ListBySubscriptions",
		Resource:    "Microsoft.ServiceBus",
	},
	"Rules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/subscriptions/{{.subscriptionName}}/rules/{{.ruleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Rules_Get",
		Resource:    "Microsoft.ServiceBus",
	},
}
