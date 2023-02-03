package policy

// Microsoft_StreamAnalytics_clusters policy
var Microsoft_StreamAnalytics_clusters = map[string]Policy{
	"Clusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Clusters_Get",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Clusters_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.StreamAnalytics/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Clusters_ListBySubscription",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Clusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Clusters_ListByResourceGroup",
		Resource:    "Microsoft.StreamAnalytics",
	},
	"Clusters_ListStreamingJobs": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.StreamAnalytics/clusters/{{.clusterName}}/listStreamingJobs",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "Clusters_ListStreamingJobs",
		Resource:    "Microsoft.StreamAnalytics",
	},
}
