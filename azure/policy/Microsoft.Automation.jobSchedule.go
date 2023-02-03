package policy

// Microsoft_Automation_jobSchedule policy
var Microsoft_Automation_jobSchedule = map[string]Policy{
	"JobSchedule_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/jobSchedules/{{.jobScheduleId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "JobSchedule_Get",
		Resource:    "Microsoft.Automation",
	},
	"JobSchedule_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/jobSchedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "JobSchedule_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
