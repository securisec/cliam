package policy

// Microsoft_MachineLearningServices_workspaceFeatures policy
var Microsoft_MachineLearningServices_workspaceFeatures = map[string]Policy{
	"WorkspaceFeatures_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MachineLearningServices/workspaces/{{.workspaceName}}/features",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "WorkspaceFeatures_List",
		Resource:    "Microsoft.MachineLearningServices",
	},
}
