package policy

// Microsoft_DocumentDB_privateLinkResources policy
var Microsoft_DocumentDB_privateLinkResources = map[string]Policy{
	"PrivateLinkResources_ListByDatabaseAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "PrivateLinkResources_ListByDatabaseAccount",
		Resource:    "Microsoft.DocumentDB",
	},
	"PrivateLinkResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.DocumentDB",
	},
}
