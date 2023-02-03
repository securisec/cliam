package policy

// Microsoft_Migrate_resourcemovercollection policy
var Microsoft_Migrate_resourcemovercollection = map[string]Policy{
	"MoveCollections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_Get",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_Prepare": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/prepare",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_Prepare",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_InitiateMove": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/initiateMove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_InitiateMove",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_Commit": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/commit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_Commit",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_Discard": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/discard",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_Discard",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_ResolveDependencies": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/resolveDependencies",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_ResolveDependencies",
		Resource:    "Microsoft.Migrate",
	},
	"MoveResources_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/moveResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveResources_List",
		Resource:    "Microsoft.Migrate",
	},
	"UnresolvedDependencies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/unresolvedDependencies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "UnresolvedDependencies_Get",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_BulkRemove": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/bulkRemove",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_BulkRemove",
		Resource:    "Microsoft.Migrate",
	},
	"MoveResources_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/moveResources/{{.moveResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveResources_Get",
		Resource:    "Microsoft.Migrate",
	},
	"OperationsDiscovery_Get": {
		Path:   "/providers/Microsoft.Migrate/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "OperationsDiscovery_Get",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_ListMoveCollectionsBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Migrate/moveCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_ListMoveCollectionsBySubscription",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_ListMoveCollectionsByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_ListMoveCollectionsByResourceGroup",
		Resource:    "Microsoft.Migrate",
	},
	"MoveCollections_ListRequiredFor": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/moveCollections/{{.moveCollectionName}}/requiredFor",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-08-01",
		},
		OperationID: "MoveCollections_ListRequiredFor",
		Resource:    "Microsoft.Migrate",
	},
}
