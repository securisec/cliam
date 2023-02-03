package policy

// Microsoft_DocumentDB_rbac policy
var Microsoft_DocumentDB_rbac = map[string]Policy{
	"SqlResources_GetSqlRoleDefinition": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleDefinitions/{{.roleDefinitionId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "SqlResources_GetSqlRoleDefinition",
		Resource:    "Microsoft.DocumentDB",
	},
	"SqlResources_ListSqlRoleDefinitions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "SqlResources_ListSqlRoleDefinitions",
		Resource:    "Microsoft.DocumentDB",
	},
	"SqlResources_GetSqlRoleAssignment": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleAssignments/{{.roleAssignmentId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "SqlResources_GetSqlRoleAssignment",
		Resource:    "Microsoft.DocumentDB",
	},
	"SqlResources_ListSqlRoleAssignments": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/sqlRoleAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "SqlResources_ListSqlRoleAssignments",
		Resource:    "Microsoft.DocumentDB",
	},
}
