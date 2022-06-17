package policy

var Microsoft_ServiceBus_Queue = map[string]Policy{
	"Queues_ListByNamespace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Queues_ListByNamespace",
		Resource:    "Microsoft.ServiceBus",
	},
	"Queues_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Queues_Get",
		Resource:    "Microsoft.ServiceBus",
	},
}
