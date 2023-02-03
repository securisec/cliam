package policy

// Microsoft_PowerBI_powerbiembedded policy
var Microsoft_PowerBI_powerbiembedded = map[string]Policy{
	"WorkspaceCollections_getByName": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/workspaceCollections/{{.workspaceCollectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_getByName",
		Resource:    "Microsoft.PowerBI",
	},
	"WorkspaceCollections_checkNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBI/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_checkNameAvailability",
		Resource:    "Microsoft.PowerBI",
	},
	"WorkspaceCollections_listByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/workspaceCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_listByResourceGroup",
		Resource:    "Microsoft.PowerBI",
	},
	"WorkspaceCollections_listBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.PowerBI/workspaceCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_listBySubscription",
		Resource:    "Microsoft.PowerBI",
	},
	"WorkspaceCollections_getAccessKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/workspaceCollections/{{.workspaceCollectionName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_getAccessKeys",
		Resource:    "Microsoft.PowerBI",
	},
	"WorkspaceCollections_regenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/workspaceCollections/{{.workspaceCollectionName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_regenerateKey",
		Resource:    "Microsoft.PowerBI",
	},
	"getAvailableOperations": {
		Path:   "/providers/Microsoft.PowerBI/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "getAvailableOperations",
		Resource:    "Microsoft.PowerBI",
	},
	"Workspaces_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.PowerBI/workspaceCollections/{{.workspaceCollectionName}}/workspaces",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "Workspaces_List",
		Resource:    "Microsoft.PowerBI",
	},
	"WorkspaceCollections_migrate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/moveResources",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-01-29",
		},
		OperationID: "WorkspaceCollections_migrate",
		Resource:    "Microsoft.PowerBI",
	},
}
