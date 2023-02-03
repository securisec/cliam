package policy

// Microsoft_Automation_python2package policy
var Microsoft_Automation_python2package = map[string]Policy{
	"Python2Package_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/python2Packages/{{.packageName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Python2Package_Get",
		Resource:    "Microsoft.Automation",
	},
	"Python2Package_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/python2Packages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Python2Package_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
