package policy

    var Microsoft_StreamAnalytics_privateEndpoints = map[string]Policy{
        "PrivateEndpoints_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/clusters/{{.clusterName}}/privateEndpoints/{{.privateEndpointName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "PrivateEndpoints_Get",
    Resource:       "Microsoft.StreamAnalytics",
},
"PrivateEndpoints_ListByCluster": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/clusters/{{.clusterName}}/privateEndpoints",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2020-03-01",
    },
	OperationID:    "PrivateEndpoints_ListByCluster",
    Resource:       "Microsoft.StreamAnalytics",
},
    }
    