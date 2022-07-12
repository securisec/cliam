package policy

    var Microsoft_StreamAnalytics_inputs = map[string]Policy{
        "Inputs_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/inputs/{{.inputName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "Inputs_Get",
    Resource:       "Microsoft.StreamAnalytics",
},
"Inputs_ListByStreamingJob": {
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/inputs",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "Inputs_ListByStreamingJob",
    Resource:       "Microsoft.StreamAnalytics",
},
"Inputs_Test": {
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/inputs/{{.inputName}}/test",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "Inputs_Test",
    Resource:       "Microsoft.StreamAnalytics",
},
    }
    