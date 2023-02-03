package policy

// Microsoft_Automation_operations policy
var Microsoft_Automation_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Automation/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Automation",
	},
	"convertGraphRunbookContent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/convertGraphRunbookContent",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "convertGraphRunbookContent",
		Resource:    "Microsoft.Automation",
	},
}
