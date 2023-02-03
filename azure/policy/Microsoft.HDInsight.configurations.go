package policy

// Microsoft_HDInsight_configurations policy
var Microsoft_HDInsight_configurations = map[string]Policy{
	"Configurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/configurations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Configurations_List",
		Resource:    "Microsoft.HDInsight",
	},
	"Configurations_Update": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/configurations/{{.configurationName}}",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Configurations_Update",
		Resource:    "Microsoft.HDInsight",
	},
	"Configurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/configurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Configurations_Get",
		Resource:    "Microsoft.HDInsight",
	},
}
