package policy

// Microsoft_HDInsight_applications policy
var Microsoft_HDInsight_applications = map[string]Policy{
	"Applications_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/applications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Applications_ListByCluster",
		Resource:    "Microsoft.HDInsight",
	},
	"Applications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/applications/{{.applicationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Applications_Get",
		Resource:    "Microsoft.HDInsight",
	},
	"Applications_GetAzureAsyncOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/applications/{{.applicationName}}/azureasyncoperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Applications_GetAzureAsyncOperationStatus",
		Resource:    "Microsoft.HDInsight",
	},
}
