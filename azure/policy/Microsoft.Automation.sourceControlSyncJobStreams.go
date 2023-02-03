package policy

// Microsoft_Automation_sourceControlSyncJobStreams policy
var Microsoft_Automation_sourceControlSyncJobStreams = map[string]Policy{
	"SourceControlSyncJobStreams_ListBySyncJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/sourceControls/{{.sourceControlName}}/sourceControlSyncJobs/{{.sourceControlSyncJobId}}/streams",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SourceControlSyncJobStreams_ListBySyncJob",
		Resource:    "Microsoft.Automation",
	},
	"SourceControlSyncJobStreams_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/sourceControls/{{.sourceControlName}}/sourceControlSyncJobs/{{.sourceControlSyncJobId}}/streams/{{.streamId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "SourceControlSyncJobStreams_Get",
		Resource:    "Microsoft.Automation",
	},
}
