package policy

var Microsoft_DocumentDB_notebook = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "NotebookWorkspaces_ListByDatabaseAccount",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "NotebookWorkspaces_Get",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}/listConnectionInfo",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "NotebookWorkspaces_ListConnectionInfo",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}/regenerateAuthToken",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "NotebookWorkspaces_RegenerateAuthToken",
		Resource:    "Microsoft.DocumentDB",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DocumentDB/databaseAccounts/{{.accountName}}/notebookWorkspaces/{{.notebookWorkspaceName}}/start",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-10-15",
		},
		OperationID: "NotebookWorkspaces_Start",
		Resource:    "Microsoft.DocumentDB",
	},
}
