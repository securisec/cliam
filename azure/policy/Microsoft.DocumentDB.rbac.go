package policy

var Microsoft_DocumentDB_rbac = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleDefinitions/{{.roleDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlRoleDefinition",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlRoleDefinitions",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleAssignments/{{.roleAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_GetSqlRoleAssignment",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "SqlResources_ListSqlRoleAssignments",
		Resource:    "Microsoft.DocumentDB",
	},
}
