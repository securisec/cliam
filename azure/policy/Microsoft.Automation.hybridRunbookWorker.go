package policy

// Microsoft_Automation_hybridRunbookWorker policy
var Microsoft_Automation_hybridRunbookWorker = map[string]Policy{
	"HybridRunbookWorkers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/hybridRunbookWorkerGroups/{{.hybridRunbookWorkerGroupName}}/hybridRunbookWorkers/{{.hybridRunbookWorkerId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "HybridRunbookWorkers_Get",
		Resource:    "Microsoft.Automation",
	},
	"HybridRunbookWorkers_Move": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/hybridRunbookWorkerGroups/{{.hybridRunbookWorkerGroupName}}/hybridRunbookWorkers/{{.hybridRunbookWorkerId}}/move",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "HybridRunbookWorkers_Move",
		Resource:    "Microsoft.Automation",
	},
	"HybridRunbookWorkers_ListByHybridRunbookWorkerGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Automation/automationAccounts/{{.automationAccountName}}/hybridRunbookWorkerGroups/{{.hybridRunbookWorkerGroupName}}/hybridRunbookWorkers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-08",
		},
		OperationID: "HybridRunbookWorkers_ListByHybridRunbookWorkerGroup",
		Resource:    "Microsoft.Automation",
	},
}
