package policy

    var Microsoft_ServiceFabric_cluster = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Clusters_Get",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Clusters_ListByResourceGroup",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/clusters",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Clusters_List",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/clusterVersions/{{.clusterVersion}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ClusterVersions_Get",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/environments/{{.environment}}/clusterVersions/{{.clusterVersion}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ClusterVersions_GetByEnvironment",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/clusterVersions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ClusterVersions_List",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.ServiceFabric/locations/{{.location}}/environments/{{.environment}}/clusterVersions",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ClusterVersions_ListByEnvironment",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.ServiceFabric/clusters/{{.clusterName}}/listUpgradableVersions",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Clusters_ListUpgradableVersions",
    Resource:       "Microsoft.ServiceFabric",
},{
    Path: "/providers/Microsoft.ServiceFabric/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.ServiceFabric",
},
    }
    