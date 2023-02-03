package policy

// Microsoft_Migrate_hubmigrate policy
var Microsoft_Migrate_hubmigrate = map[string]Policy{
	"MigrateProjectsController_GetMigrateProject": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/migrateProjects/{{.migrateProjectName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "MigrateProjectsController_GetMigrateProject",
		Resource:    "Microsoft.Migrate",
	},
	"PrivateEndpointConnectionsController_GetPrivateEndpointConnections": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/migrateProjects/{{.migrateProjectName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "PrivateEndpointConnectionsController_GetPrivateEndpointConnections",
		Resource:    "Microsoft.Migrate",
	},
	"PrivateEndpointConnectionController_GetPrivateEndpointConnection": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/migrateProjects/{{.migrateProjectName}}/privateEndpointConnections/{{.peConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "PrivateEndpointConnectionController_GetPrivateEndpointConnection",
		Resource:    "Microsoft.Migrate",
	},
	"PrivateLinkResourceController_GetPrivateLinkResource": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/migrateProjects/{{.migrateProjectName}}/privateLinkResources/{{.privateLinkResourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "PrivateLinkResourceController_GetPrivateLinkResource",
		Resource:    "Microsoft.Migrate",
	},
	"PrivateLinkResourceController_GetPrivateLinkResources": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/migrateProjects/{{.migrateProjectName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "PrivateLinkResourceController_GetPrivateLinkResources",
		Resource:    "Microsoft.Migrate",
	},
	"Projects_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Migrate/migrateProjects",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "Projects_ListBySubscription",
		Resource:    "Microsoft.Migrate",
	},
	"Projects_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Migrate/migrateProjects",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "Projects_List",
		Resource:    "Microsoft.Migrate",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Migrate/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Migrate",
	},
}
