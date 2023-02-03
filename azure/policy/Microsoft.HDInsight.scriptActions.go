package policy

// Microsoft_HDInsight_scriptActions policy
var Microsoft_HDInsight_scriptActions = map[string]Policy{
	"Clusters_ExecuteScriptActions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/executeScriptActions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_ExecuteScriptActions",
		Resource:    "Microsoft.HDInsight",
	},
	"ScriptActions_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/scriptActions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ScriptActions_ListByCluster",
		Resource:    "Microsoft.HDInsight",
	},
	"ScriptActions_GetExecutionDetail": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/scriptExecutionHistory/{{.scriptExecutionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ScriptActions_GetExecutionDetail",
		Resource:    "Microsoft.HDInsight",
	},
	"ScriptExecutionHistory_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/scriptExecutionHistory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ScriptExecutionHistory_ListByCluster",
		Resource:    "Microsoft.HDInsight",
	},
	"ScriptExecutionHistory_Promote": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/scriptExecutionHistory/{{.scriptExecutionId}}/promote",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ScriptExecutionHistory_Promote",
		Resource:    "Microsoft.HDInsight",
	},
	"ScriptActions_GetExecutionAsyncOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/executeScriptActions/azureasyncoperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ScriptActions_GetExecutionAsyncOperationStatus",
		Resource:    "Microsoft.HDInsight",
	},
}
