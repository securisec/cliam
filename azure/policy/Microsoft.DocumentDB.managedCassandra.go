package policy

var Microsoft_DocumentDB_managedCassandra = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/cassandraClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_ListBySubscription",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_ListByResourceGroup",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_Get",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/invokeCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_InvokeCommand",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/dataCenters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraDataCenters_List",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/dataCenters/{{.dataCenterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraDataCenters_Get",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_Deallocate",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_Start",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/status",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "CassandraClusters_Status",
		Resource:    "Microsoft.DocumentDB",
	},
}
