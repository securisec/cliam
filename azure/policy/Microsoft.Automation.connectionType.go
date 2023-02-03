package policy

// Microsoft_Automation_connectionType policy
var Microsoft_Automation_connectionType = map[string]Policy{
	"ConnectionType_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/connectionTypes/{{.connectionTypeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "ConnectionType_Get",
		Resource:    "Microsoft.Automation",
	},
	"ConnectionType_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/connectionTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "ConnectionType_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
