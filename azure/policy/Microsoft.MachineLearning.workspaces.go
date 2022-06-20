package policy

    var Microsoft_MachineLearning_workspaces = map[string]Policy{
        "Operations_List": {
    Path: "/providers/Microsoft.MachineLearning/operations",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-10-01",
    },
	OperationID:    "Operations_List",
    Resource:       "Microsoft.MachineLearning",
},
"Workspaces_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearning/workspaces/{{.workspaceName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-10-01",
    },
	OperationID:    "Workspaces_Get",
    Resource:       "Microsoft.MachineLearning",
},
"Workspaces_ResyncStorageKeys": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearning/workspaces/{{.workspaceName}}/resyncStorageKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-10-01",
    },
	OperationID:    "Workspaces_ResyncStorageKeys",
    Resource:       "Microsoft.MachineLearning",
},
"Workspaces_ListWorkspaceKeys": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearning/workspaces/{{.workspaceName}}/listWorkspaceKeys",
	Method: "POST",
	QueryValues:   map[string]string{
        "api-version": "2019-10-01",
    },
	OperationID:    "Workspaces_ListWorkspaceKeys",
    Resource:       "Microsoft.MachineLearning",
},
"Workspaces_ListByResourceGroup": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearning/workspaces",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-10-01",
    },
	OperationID:    "Workspaces_ListByResourceGroup",
    Resource:       "Microsoft.MachineLearning",
},
"Workspaces_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.MachineLearning/workspaces",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2019-10-01",
    },
	OperationID:    "Workspaces_List",
    Resource:       "Microsoft.MachineLearning",
},
    }
    