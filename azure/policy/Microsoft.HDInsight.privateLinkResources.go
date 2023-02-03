package policy

// Microsoft_HDInsight_privateLinkResources policy
var Microsoft_HDInsight_privateLinkResources = map[string]Policy{
	"PrivateLinkResources_ListByCluster": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkResources_ListByCluster",
		Resource:    "Microsoft.HDInsight",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.HDInsight",
	},
}
