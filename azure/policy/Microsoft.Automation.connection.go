package policy

// Microsoft_Automation_connection policy
var Microsoft_Automation_connection = map[string]Policy{
	"Connection_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/connections/{{.connectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Connection_Get",
		Resource:    "Microsoft.Automation",
	},
	"Connection_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/connections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Connection_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
