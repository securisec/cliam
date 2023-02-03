package policy

// Microsoft_Automation_hybridRunbookWorkerGroup policy
var Microsoft_Automation_hybridRunbookWorkerGroup = map[string]Policy{
	"HybridRunbookWorkerGroup_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/hybridRunbookWorkerGroups/{{.hybridRunbookWorkerGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "HybridRunbookWorkerGroup_Get",
		Resource:    "Microsoft.Automation",
	},
	"HybridRunbookWorkerGroup_ListByAutomationAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/hybridRunbookWorkerGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "HybridRunbookWorkerGroup_ListByAutomationAccount",
		Resource:    "Microsoft.Automation",
	},
}
