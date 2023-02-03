package policy

// Microsoft_Automation_softwareUpdateConfigurationRun policy
var Microsoft_Automation_softwareUpdateConfigurationRun = map[string]Policy{
	"SoftwareUpdateConfigurationRuns_GetById": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/softwareUpdateConfigurationRuns/{{.softwareUpdateConfigurationRunId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SoftwareUpdateConfigurationRuns_GetById",
		Resource:    "Microsoft.Automation",
	},
	"SoftwareUpdateConfigurationRuns_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/softwareUpdateConfigurationRuns",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SoftwareUpdateConfigurationRuns_List",
		Resource:    "Microsoft.Automation",
	},
}
