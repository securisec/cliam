package policy

    var Microsoft_ServiceBus_Queue = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Queues_ListByNamespace",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Queues_Get",
    Resource:       "Microsoft.ServiceBus",
},
    }
    