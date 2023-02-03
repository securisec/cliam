package policy

// Microsoft_Automation_dscConfiguration policy
var Microsoft_Automation_dscConfiguration = map[string]Policy{
	"DscConfiguration_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/configurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "DscConfiguration_Get",
		Resource:    "Microsoft.Automation",
	},
	"DscConfiguration_GetContent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/configurations/{{.configurationName}}/content",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "DscConfiguration_GetContent",
		Resource:    "Microsoft.Automation",
	},
	"DscConfiguration_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/configurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "DscConfiguration_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
