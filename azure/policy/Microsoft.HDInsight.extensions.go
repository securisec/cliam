package policy

// Microsoft_HDInsight_extensions policy
var Microsoft_HDInsight_extensions = map[string]Policy{
	"Extensions_GetMonitoringStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/extensions/clustermonitoring",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Extensions_GetMonitoringStatus",
		Resource:    "Microsoft.HDInsight",
	},
	"Extensions_GetAzureMonitorStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/extensions/azureMonitor",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Extensions_GetAzureMonitorStatus",
		Resource:    "Microsoft.HDInsight",
	},
	"Extensions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/extensions/{{.extensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Extensions_Get",
		Resource:    "Microsoft.HDInsight",
	},
	"Extensions_GetAzureAsyncOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/extensions/{{.extensionName}}/azureAsyncOperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Extensions_GetAzureAsyncOperationStatus",
		Resource:    "Microsoft.HDInsight",
	},
}
