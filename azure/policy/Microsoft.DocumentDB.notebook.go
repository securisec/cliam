package policy

// Microsoft_DocumentDB_notebook policy
var Microsoft_DocumentDB_notebook = map[string]Policy{
	"NotebookWorkspaces_ListByDatabaseAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "NotebookWorkspaces_ListByDatabaseAccount",
		Resource:    "Microsoft.DocumentDB",
	},
	"NotebookWorkspaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "NotebookWorkspaces_Get",
		Resource:    "Microsoft.DocumentDB",
	},
	"NotebookWorkspaces_ListConnectionInfo": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}/listConnectionInfo",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "NotebookWorkspaces_ListConnectionInfo",
		Resource:    "Microsoft.DocumentDB",
	},
	"NotebookWorkspaces_RegenerateAuthToken": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}/regenerateAuthToken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "NotebookWorkspaces_RegenerateAuthToken",
		Resource:    "Microsoft.DocumentDB",
	},
	"NotebookWorkspaces_Start": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-08-15",
		},
		OperationID: "NotebookWorkspaces_Start",
		Resource:    "Microsoft.DocumentDB",
	},
}
