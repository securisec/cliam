package policy

// Microsoft_HDInsight_cluster policy
var Microsoft_HDInsight_cluster = map[string]Policy{
	"Clusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_Get",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_ListByResourceGroup",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_Resize": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/roles/{{.roleName}}/resize",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_Resize",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_UpdateAutoScaleConfiguration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/roles/{{.roleName}}/autoscale",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_UpdateAutoScaleConfiguration",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HDInsight/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_List",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_RotateDiskEncryptionKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/rotatediskencryptionkey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_RotateDiskEncryptionKey",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_GetGatewaySettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/getGatewaySettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_GetGatewaySettings",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_UpdateGatewaySettings": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/updateGatewaySettings",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_UpdateGatewaySettings",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_GetAzureAsyncOperationStatus": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/azureasyncoperations/{{.operationId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_GetAzureAsyncOperationStatus",
		Resource:    "Microsoft.HDInsight",
	},
	"Clusters_UpdateIdentityCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HDInsight/clusters/{{.clusterName}}/updateClusterIdentityCertificate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_UpdateIdentityCertificate",
		Resource:    "Microsoft.HDInsight",
	},
}
