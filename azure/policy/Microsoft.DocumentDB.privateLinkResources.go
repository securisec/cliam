package policy

var Microsoft_DocumentDB_privateLinkResources = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PrivateLinkResources_ListByDatabaseAccount",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/privateLinkResources/{{.groupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "PrivateLinkResources_Get",
		Resource:    "Microsoft.DocumentDB",
	},
}
