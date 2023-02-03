package policy

// Microsoft_HDInsight_privateEndpointConnections policy
var Microsoft_HDInsight_privateEndpointConnections = map[string]Policy{
	"PrivateEndpointConnections_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnections_ListByCluster",
		Resource:    "Microsoft.HDInsight",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.HDInsight",
	},
}
