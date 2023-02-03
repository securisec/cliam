package policy

// Microsoft_Automation_sourceControlSyncJob policy
var Microsoft_Automation_sourceControlSyncJob = map[string]Policy{
	"SourceControlSyncJob_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/sourceControls/{{.sourceControlName}}/sourceControlSyncJobs/{{.sourceControlSyncJobId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SourceControlSyncJob_Get",
		Resource:    "Microsoft.Automation",
	},
	"SourceControlSyncJob_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/sourceControls/{{.sourceControlName}}/sourceControlSyncJobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SourceControlSyncJob_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
