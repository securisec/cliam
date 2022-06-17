package policy

var Microsoft_ServiceBus_DisasterRecoveryConfig = map[string]Policy{
	"DisasterRecoveryConfigs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_List",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_Get",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_BreakPairing": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/breakPairing",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_BreakPairing",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_FailOver": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/{{.alias}}/failover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_FailOver",
		Resource:    "Microsoft.ServiceBus",
	},
}
