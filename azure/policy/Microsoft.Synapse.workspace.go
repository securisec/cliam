package policy

// Microsoft_Synapse_workspace policy
var Microsoft_Synapse_workspace = map[string]Policy{
	"Workspaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Workspaces_ListByResourceGroup",
		Resource:    "Microsoft.Synapse",
	},
	"Workspaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Workspaces_Get",
		Resource:    "Microsoft.Synapse",
	},
	"Workspaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Synapse/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "Workspaces_List",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceAadAdmins_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/administrators/activeDirectory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceAadAdmins_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceSqlAadAdmins_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/sqlAdministrators/activeDirectory",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceSqlAadAdmins_Get",
		Resource:    "Microsoft.Synapse",
	},
	"WorkspaceManagedIdentitySqlControlSettings_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/managedIdentitySqlControlSettings/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "WorkspaceManagedIdentitySqlControlSettings_Get",
		Resource:    "Microsoft.Synapse",
	},
	"RestorableDroppedSqlPools_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/restorableDroppedSqlPools/{{.restorableDroppedSqlPoolId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "RestorableDroppedSqlPools_Get",
		Resource:    "Microsoft.Synapse",
	},
	"RestorableDroppedSqlPools_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/restorableDroppedSqlPools",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "RestorableDroppedSqlPools_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
}
