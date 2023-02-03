package policy

// Microsoft_Automation_dscNodeConfiguration policy
var Microsoft_Automation_dscNodeConfiguration = map[string]Policy{
	"DscNodeConfiguration_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/nodeConfigurations/{{.nodeConfigurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "DscNodeConfiguration_Get",
		Resource:    "Microsoft.Automation",
	},
	"DscNodeConfiguration_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/nodeConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "DscNodeConfiguration_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
