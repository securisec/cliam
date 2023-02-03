package policy

// Microsoft_DelegatedNetwork_orchestrators policy
var Microsoft_DelegatedNetwork_orchestrators = map[string]Policy{
	"OrchestratorInstanceService_GetDetails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DelegatedNetwork/orchestrators/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "OrchestratorInstanceService_GetDetails",
		Resource:    "Microsoft.DelegatedNetwork",
	},
	"OrchestratorInstanceService_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DelegatedNetwork/orchestrators",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "OrchestratorInstanceService_ListBySubscription",
		Resource:    "Microsoft.DelegatedNetwork",
	},
	"OrchestratorInstanceService_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DelegatedNetwork/orchestrators",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-15",
		},
		OperationID: "OrchestratorInstanceService_ListByResourceGroup",
		Resource:    "Microsoft.DelegatedNetwork",
	},
}
