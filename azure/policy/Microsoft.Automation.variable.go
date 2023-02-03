package policy

// Microsoft_Automation_variable policy
var Microsoft_Automation_variable = map[string]Policy{
	"Variable_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/variables/{{.variableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Variable_Get",
		Resource:    "Microsoft.Automation",
	},
	"Variable_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/variables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Variable_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
