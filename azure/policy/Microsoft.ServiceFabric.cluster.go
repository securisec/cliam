package policy

// Microsoft_ServiceFabric_cluster policy
var Microsoft_ServiceFabric_cluster = map[string]Policy{
	"Clusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Clusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_ListByResourceGroup",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Clusters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/clusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_List",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ClusterVersions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/clusterVersions/{{.clusterVersion}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ClusterVersions_Get",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ClusterVersions_GetByEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/environments/{{.environment}}/clusterVersions/{{.clusterVersion}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ClusterVersions_GetByEnvironment",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ClusterVersions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/clusterVersions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ClusterVersions_List",
		Resource:    "Microsoft.ServiceFabric",
	},
	"ClusterVersions_ListByEnvironment": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/environments/{{.environment}}/clusterVersions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "ClusterVersions_ListByEnvironment",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Clusters_ListUpgradableVersions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/listUpgradableVersions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Clusters_ListUpgradableVersions",
		Resource:    "Microsoft.ServiceFabric",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.ServiceFabric/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.ServiceFabric",
	},
}
