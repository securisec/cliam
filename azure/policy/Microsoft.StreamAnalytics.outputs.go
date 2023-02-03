package policy

// Microsoft_StreamAnalytics_outputs policy
var Microsoft_StreamAnalytics_outputs = map[string]Policy{
	"Outputs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/outputs/{{.outputName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Outputs_Get",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Outputs_ListByStreamingJob": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/outputs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Outputs_ListByStreamingJob",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Outputs_Test": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/outputs/{{.outputName}}/test",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Outputs_Test",
		Resource:    "Microsoft.StreamAnalytics",
	},
}
