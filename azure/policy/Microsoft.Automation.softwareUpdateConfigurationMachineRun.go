package policy

// Microsoft_Automation_softwareUpdateConfigurationMachineRun policy
var Microsoft_Automation_softwareUpdateConfigurationMachineRun = map[string]Policy{
	"SoftwareUpdateConfigurationMachineRuns_GetById": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/softwareUpdateConfigurationMachineRuns/{{.softwareUpdateConfigurationMachineRunId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SoftwareUpdateConfigurationMachineRuns_GetById",
		Resource:    "Microsoft.Automation",
	},
	"SoftwareUpdateConfigurationMachineRuns_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/softwareUpdateConfigurationMachineRuns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SoftwareUpdateConfigurationMachineRuns_List",
		Resource:    "Microsoft.Automation",
	},
}
