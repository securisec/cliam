package policy

// Microsoft_Automation_runbook policy
var Microsoft_Automation_runbook = map[string]Policy{
	"RunbookDraft_GetContent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/content",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "RunbookDraft_GetContent",
		Resource:    "Microsoft.Automation",
	},
	"RunbookDraft_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "RunbookDraft_Get",
		Resource:    "Microsoft.Automation",
	},
	"Runbook_Publish": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/publish",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Runbook_Publish",
		Resource:    "Microsoft.Automation",
	},
	"RunbookDraft_UndoEdit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/undoEdit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "RunbookDraft_UndoEdit",
		Resource:    "Microsoft.Automation",
	},
	"Runbook_GetContent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/content",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Runbook_GetContent",
		Resource:    "Microsoft.Automation",
	},
	"Runbook_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Runbook_Get",
		Resource:    "Microsoft.Automation",
	},
	"Runbook_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Runbook_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
	"TestJobStreams_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/testJob/streams/{{.jobStreamId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "TestJobStreams_Get",
		Resource:    "Microsoft.Automation",
	},
	"TestJobStreams_ListByTestJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/testJob/streams",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "TestJobStreams_ListByTestJob",
		Resource:    "Microsoft.Automation",
	},
	"TestJob_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/testJob",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "TestJob_Get",
		Resource:    "Microsoft.Automation",
	},
	"TestJob_Resume": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/testJob/resume",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "TestJob_Resume",
		Resource:    "Microsoft.Automation",
	},
	"TestJob_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/testJob/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "TestJob_Stop",
		Resource:    "Microsoft.Automation",
	},
	"TestJob_Suspend": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/runbooks/{{.runbookName}}/draft/testJob/suspend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "TestJob_Suspend",
		Resource:    "Microsoft.Automation",
	},
}
