package policy

// Microsoft_OperationalInsights_Tables policy
var Microsoft_OperationalInsights_Tables = map[string]Policy{
	"Tables_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/workspaces/{{.workspaceName}}/tables",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Tables_ListByWorkspace",
		Resource:    "Microsoft.OperationalInsights",
	},
	"Tables_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/workspaces/{{.workspaceName}}/tables/{{.tableName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Tables_Get",
		Resource:    "Microsoft.OperationalInsights",
	},
	"Tables_Migrate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/workspaces/{{.workspaceName}}/tables/{{.tableName}}/migrate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Tables_Migrate",
		Resource:    "Microsoft.OperationalInsights",
	},
	"Tables_CancelSearch": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/workspaces/{{.workspaceName}}/tables/{{.tableName}}/cancelSearch",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Tables_CancelSearch",
		Resource:    "Microsoft.OperationalInsights",
	},
}
