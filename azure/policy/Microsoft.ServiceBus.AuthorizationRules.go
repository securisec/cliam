package policy

    var Microsoft_ServiceBus_AuthorizationRules = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Namespaces_ListAuthorizationRules",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Namespaces_GetAuthorizationRule",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}/listKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Namespaces_ListKeys",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}/regenerateKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Namespaces_RegenerateKeys",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Queues_ListAuthorizationRules",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules/{{.authorizationRuleName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Queues_GetAuthorizationRule",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules/{{.authorizationRuleName}}/ListKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Queues_ListKeys",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Queues_RegenerateKeys",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Topics_ListAuthorizationRules",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules/{{.authorizationRuleName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Topics_GetAuthorizationRule",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules/{{.authorizationRuleName}}/ListKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Topics_ListKeys",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "Topics_RegenerateKeys",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/authorizationRules",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "DisasterRecoveryConfigs_ListAuthorizationRules",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/authorizationRules/{{.authorizationRuleName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "DisasterRecoveryConfigs_GetAuthorizationRule",
    Resource:       "Microsoft.ServiceBus",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/authorizationRules/{{.authorizationRuleName}}/listKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-11-01",
    },
	OperationID:    "DisasterRecoveryConfigs_ListKeys",
    Resource:       "Microsoft.ServiceBus",
},
    }
    