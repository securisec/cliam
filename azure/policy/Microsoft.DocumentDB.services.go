package policy

// Microsoft_DocumentDB_services policy
var Microsoft_DocumentDB_services = map[string]Policy{
	"Service_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/services",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "Service_List",
		Resource:    "Microsoft.DocumentDB",
	},
	"Service_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/services/{{.serviceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "Service_Get",
		Resource:    "Microsoft.DocumentDB",
	},
}
