package policy

// Microsoft_Automation_sourceControl policy
var Microsoft_Automation_sourceControl = map[string]Policy{
	"SourceControl_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/sourceControls/{{.sourceControlName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SourceControl_Get",
		Resource:    "Microsoft.Automation",
	},
	"SourceControl_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/sourceControls",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SourceControl_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
