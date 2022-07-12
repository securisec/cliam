package policy

    var Microsoft_StreamAnalytics_transformations = map[string]Policy{
        "Transformations_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/streamingjobs/{{.jobName}}/transformations/{{.transformationName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "Transformations_Get",
    Resource:       "Microsoft.StreamAnalytics",
},
    }
    