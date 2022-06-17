package policy

    var Microsoft_ServiceBus_subscriptions = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/subscriptions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Subscriptions_ListByTopic",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/subscriptions/{{.subscriptionName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Subscriptions_Get",
    Resource:       "Microsoft.ServiceBus",
},
    }
    