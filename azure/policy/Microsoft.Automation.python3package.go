package policy

// Microsoft_Automation_python3package policy
var Microsoft_Automation_python3package = map[string]Policy{
	"Python3Package_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/python3Packages/{{.packageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Python3Package_Get",
		Resource:    "Microsoft.Automation",
	},
	"Python3Package_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/python3Packages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Python3Package_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
