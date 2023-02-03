package policy

// Microsoft_ServiceBus_AuthorizationRules policy
var Microsoft_ServiceBus_AuthorizationRules = map[string]Policy{
	"Namespaces_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListAuthorizationRules",
		Resource:    "Microsoft.ServiceBus",
	},
	"Namespaces_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_GetAuthorizationRule",
		Resource:    "Microsoft.ServiceBus",
	},
	"Namespaces_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_ListKeys",
		Resource:    "Microsoft.ServiceBus",
	},
	"Namespaces_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/AuthorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_RegenerateKeys",
		Resource:    "Microsoft.ServiceBus",
	},
	"Queues_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Queues_ListAuthorizationRules",
		Resource:    "Microsoft.ServiceBus",
	},
	"Queues_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Queues_GetAuthorizationRule",
		Resource:    "Microsoft.ServiceBus",
	},
	"Queues_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules/{{.authorizationRuleName}}/ListKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Queues_ListKeys",
		Resource:    "Microsoft.ServiceBus",
	},
	"Queues_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/queues/{{.queueName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Queues_RegenerateKeys",
		Resource:    "Microsoft.ServiceBus",
	},
	"Topics_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Topics_ListAuthorizationRules",
		Resource:    "Microsoft.ServiceBus",
	},
	"Topics_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Topics_GetAuthorizationRule",
		Resource:    "Microsoft.ServiceBus",
	},
	"Topics_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules/{{.authorizationRuleName}}/ListKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Topics_ListKeys",
		Resource:    "Microsoft.ServiceBus",
	},
	"Topics_RegenerateKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/topics/{{.topicName}}/authorizationRules/{{.authorizationRuleName}}/regenerateKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Topics_RegenerateKeys",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_ListAuthorizationRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/authorizationRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_ListAuthorizationRules",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_GetAuthorizationRule": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/authorizationRules/{{.authorizationRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_GetAuthorizationRule",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/authorizationRules/{{.authorizationRuleName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_ListKeys",
		Resource:    "Microsoft.ServiceBus",
	},
}
