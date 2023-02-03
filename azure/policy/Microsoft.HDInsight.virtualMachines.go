package policy

// Microsoft_HDInsight_virtualMachines policy
var Microsoft_HDInsight_virtualMachines = map[string]Policy{
	"VirtualMachines_ListHosts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/listHosts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "VirtualMachines_ListHosts",
		Resource:    "Microsoft.HDInsight",
	},
	"VirtualMachines_RestartHosts": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/restartHosts",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "VirtualMachines_RestartHosts",
		Resource:    "Microsoft.HDInsight",
	},
	"VirtualMachines_GetAsyncOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/restartHosts/azureasyncoperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "VirtualMachines_GetAsyncOperationStatus",
		Resource:    "Microsoft.HDInsight",
	},
}
