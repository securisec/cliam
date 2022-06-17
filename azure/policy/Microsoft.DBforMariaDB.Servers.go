package policy

var Microsoft_DBforMariaDB_Servers = map[string]Policy{
	"Servers_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMariaDB/servers/{{.serverName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-01-01",
		},
		OperationID: "Servers_Start",
		Resource:    "Microsoft.DBforMariaDB",
	},
	"Servers_Stop": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DBforMariaDB/servers/{{.serverName}}/stop",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-01-01",
		},
		OperationID: "Servers_Stop",
		Resource:    "Microsoft.DBforMariaDB",
	},
}
