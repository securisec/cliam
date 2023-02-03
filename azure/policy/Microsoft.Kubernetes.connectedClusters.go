package policy

// Microsoft_Kubernetes_connectedClusters policy
var Microsoft_Kubernetes_connectedClusters = map[string]Policy{
	"ConnectedCluster_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Kubernetes/connectedClusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ConnectedCluster_Get",
		Resource:    "Microsoft.Kubernetes",
	},
	"ConnectedCluster_ListClusterUserCredential": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Kubernetes/connectedClusters/{{.clusterName}}/listClusterUserCredential",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ConnectedCluster_ListClusterUserCredential",
		Resource:    "Microsoft.Kubernetes",
	},
	"ConnectedCluster_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Kubernetes/connectedClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ConnectedCluster_ListByResourceGroup",
		Resource:    "Microsoft.Kubernetes",
	},
	"ConnectedCluster_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Kubernetes/connectedClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "ConnectedCluster_ListBySubscription",
		Resource:    "Microsoft.Kubernetes",
	},
	"Operations_Get": {
		Path:   "/providers/Microsoft.Kubernetes/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Operations_Get",
		Resource:    "Microsoft.Kubernetes",
	},
}
