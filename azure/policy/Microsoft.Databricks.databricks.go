package policy

    var Microsoft_Databricks_databricks = map[string]Policy{
        "Workspaces_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces/{{.workspaceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-04-01",
    },
	OperationID:    "Workspaces_Get",
    Resource:       "Microsoft.Databricks",
},
"Workspaces_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Databricks/workspaces",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-04-01",
    },
	OperationID:    "Workspaces_ListByResourceGroup",
    Resource:       "Microsoft.Databricks",
},
"Workspaces_ListBySubscription": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Databricks/workspaces",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-04-01",
    },
	OperationID:    "Workspaces_ListBySubscription",
    Resource:       "Microsoft.Databricks",
},
"Operations_List": {
    Path: "/providers/Microsoft.Databricks/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2018-04-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.Databricks",
},
    }
    