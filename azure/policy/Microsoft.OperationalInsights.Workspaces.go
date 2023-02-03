package policy

// Microsoft_OperationalInsights_Workspaces policy
var Microsoft_OperationalInsights_Workspaces = map[string]Policy{
	"Workspaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.OperationalInsights/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Workspaces_List",
		Resource:    "Microsoft.OperationalInsights",
	},
	"Workspaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Workspaces_ListByResourceGroup",
		Resource:    "Microsoft.OperationalInsights",
	},
	"Workspaces_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/workspaces/{{.workspaceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "Workspaces_Get",
		Resource:    "Microsoft.OperationalInsights",
	},
	"DeletedWorkspaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.OperationalInsights/deletedWorkspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "DeletedWorkspaces_List",
		Resource:    "Microsoft.OperationalInsights",
	},
	"DeletedWorkspaces_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.OperationalInsights/deletedWorkspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "DeletedWorkspaces_ListByResourceGroup",
		Resource:    "Microsoft.OperationalInsights",
	},
}
