package policy

var Microsoft_ServiceBus_CheckNameAvailability = map[string]Policy{
	"Namespaces_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceBus/CheckNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Namespaces_CheckNameAvailability",
		Resource:    "Microsoft.ServiceBus",
	},
	"DisasterRecoveryConfigs_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceBus/namespaces/{{.namespaceName}}/disasterRecoveryConfigs/CheckNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DisasterRecoveryConfigs_CheckNameAvailability",
		Resource:    "Microsoft.ServiceBus",
	},
}
