package policy

// Microsoft_StreamAnalytics_functions policy
var Microsoft_StreamAnalytics_functions = map[string]Policy{
	"Functions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/functions/{{.functionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Functions_Get",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Functions_ListByStreamingJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/functions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Functions_ListByStreamingJob",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Functions_Test": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/functions/{{.functionName}}/test",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Functions_Test",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Functions_RetrieveDefaultDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/functions/{{.functionName}}/retrieveDefaultDefinition",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Functions_RetrieveDefaultDefinition",
		Resource:    "Microsoft.StreamAnalytics",
	},
}
