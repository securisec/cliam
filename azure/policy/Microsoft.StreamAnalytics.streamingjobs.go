package policy

// Microsoft_StreamAnalytics_streamingjobs policy
var Microsoft_StreamAnalytics_streamingjobs = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.StreamAnalytics/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"StreamingJobs_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "StreamingJobs_Get",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"StreamingJobs_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "StreamingJobs_ListByResourceGroup",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"StreamingJobs_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StreamAnalytics/streamingjobs",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "StreamingJobs_List",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"StreamingJobs_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "StreamingJobs_Start",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"StreamingJobs_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "StreamingJobs_Stop",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"StreamingJobs_Scale": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/scale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "StreamingJobs_Scale",
		Resource:    "Microsoft.StreamAnalytics",
	},
}
