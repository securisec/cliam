package policy

// Microsoft_LabServices_OperationResults policy
var Microsoft_LabServices_OperationResults = map[string]Policy{
	"OperationResults_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.LabServices/operationResults/{{.operationResultId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-01",
		},
		OperationID: "OperationResults_Get",
		Resource:    "Microsoft.LabServices",
	},
}
