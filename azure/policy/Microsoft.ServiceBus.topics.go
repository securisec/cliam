package policy

// Microsoft_ServiceBus_topics policy
var Microsoft_ServiceBus_topics = map[string]Policy{
	"Topics_ListByNamespace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Topics_ListByNamespace",
		Resource:    "Microsoft.ServiceBus",
	},
	"Topics_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Topics_Get",
		Resource:    "Microsoft.ServiceBus",
	},
}
