package policy

// Microsoft_Synapse_operations policy
var Microsoft_Synapse_operations = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Synapse/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Synapse",
	},
	"Operations_GetLocationHeaderResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/operationResults/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_GetLocationHeaderResult",
		Resource:    "Microsoft.Synapse",
	},
	"Operations_GetAzureAsyncHeaderResult": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/operationStatuses/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_GetAzureAsyncHeaderResult",
		Resource:    "Microsoft.Synapse",
	},
}
