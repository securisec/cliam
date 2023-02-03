package policy

// Microsoft_Automation_schedule policy
var Microsoft_Automation_schedule = map[string]Policy{
	"Schedule_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/schedules/{{.scheduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Schedule_Get",
		Resource:    "Microsoft.Automation",
	},
	"Schedule_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/schedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Schedule_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
