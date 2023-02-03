package policy

// Microsoft_Maintenance_Maintenance policy
var Microsoft_Maintenance_Maintenance = map[string]Policy{
	"PublicMaintenanceConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "PublicMaintenanceConfigurations_List",
		Resource:    "Microsoft.Maintenance",
	},
	"PublicMaintenanceConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "PublicMaintenanceConfigurations_Get",
		Resource:    "Microsoft.Maintenance",
	},
	"ApplyUpdates_GetParent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.providerName}}/{{.resourceParentType}}/{{.resourceParentName}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Maintenance/applyUpdates/{{.applyUpdateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "ApplyUpdates_GetParent",
		Resource:    "Microsoft.Maintenance",
	},
	"ApplyUpdates_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.providerName}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Maintenance/applyUpdates/{{.applyUpdateName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "ApplyUpdates_Get",
		Resource:    "Microsoft.Maintenance",
	},
	"ConfigurationAssignments_ListParent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.providerName}}/{{.resourceParentType}}/{{.resourceParentName}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Maintenance/configurationAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "ConfigurationAssignments_ListParent",
		Resource:    "Microsoft.Maintenance",
	},
	"ConfigurationAssignments_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.providerName}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Maintenance/configurationAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "ConfigurationAssignments_List",
		Resource:    "Microsoft.Maintenance",
	},
	"MaintenanceConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Maintenance/maintenanceConfigurations/{{.resourceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "MaintenanceConfigurations_Get",
		Resource:    "Microsoft.Maintenance",
	},
	"MaintenanceConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Maintenance/maintenanceConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "MaintenanceConfigurations_List",
		Resource:    "Microsoft.Maintenance",
	},
	"MaintenanceConfigurationsForResourceGroup_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maintenance/maintenanceConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "MaintenanceConfigurationsForResourceGroup_List",
		Resource:    "Microsoft.Maintenance",
	},
	"ApplyUpdates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Maintenance/applyUpdates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "ApplyUpdates_List",
		Resource:    "Microsoft.Maintenance",
	},
	"ApplyUpdateForResourceGroup_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Maintenance/applyUpdates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "ApplyUpdateForResourceGroup_List",
		Resource:    "Microsoft.Maintenance",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Maintenance/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Maintenance",
	},
	"Updates_ListParent": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.providerName}}/{{.resourceParentType}}/{{.resourceParentName}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Maintenance/updates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Updates_ListParent",
		Resource:    "Microsoft.Maintenance",
	},
	"Updates_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/{{.providerName}}/{{.resourceType}}/{{.resourceName}}/providers/Microsoft.Maintenance/updates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-05-01",
		},
		OperationID: "Updates_List",
		Resource:    "Microsoft.Maintenance",
	},
}
