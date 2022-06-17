package policy

var Microsoft_DocumentDB_privateEndpointConnection = map[string]Policy{
	"PrivateEndpointConnections_ListByDatabaseAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PrivateEndpointConnections_ListByDatabaseAccount",
		Resource:    "Microsoft.DocumentDB",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.DocumentDB",
	},
}
