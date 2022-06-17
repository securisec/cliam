package policy

var Microsoft_DBforPostgreSQL_Databases = map[string]Policy{
	"Databases_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/databases/{{.databaseName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Databases_Get",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
	"Databases_ListByServer": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{{.serverName}}/databases",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Databases_ListByServer",
		Resource:    "Microsoft.DBforPostgreSQL",
	},
}
