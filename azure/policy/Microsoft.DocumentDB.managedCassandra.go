package policy

// Microsoft_DocumentDB_managedCassandra policy
var Microsoft_DocumentDB_managedCassandra = map[string]Policy{
	"CassandraClusters_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DocumentDB/cassandraClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_ListBySubscription",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraClusters_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_ListByResourceGroup",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraClusters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_Get",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraClusters_InvokeCommand": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/invokeCommand",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_InvokeCommand",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraDataCenters_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/dataCenters",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraDataCenters_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraDataCenters_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/dataCenters/{{.dataCenterName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraDataCenters_Get",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraClusters_Deallocate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/deallocate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_Deallocate",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraClusters_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_Start",
		Resource:    "Microsoft.DocumentDB",
	},
	"CassandraClusters_Status": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/cassandraClusters/{{.clusterName}}/status",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "CassandraClusters_Status",
		Resource:    "Microsoft.DocumentDB",
	},
}
