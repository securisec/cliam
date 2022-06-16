package policy

var Microsoft_DBforPostgreSQL_Databases = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/databases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Databases_Get",
		Resource:    "Microsoft.DBforPostgreSQL",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/databases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Databases_ListByServer",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
}
