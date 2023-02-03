package policy

// Microsoft_AzureArcData_azurearcdata policy
var Microsoft_AzureArcData_azurearcdata = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.AzureArcData/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.AzureArcData",
	},
	"SqlManagedInstances_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureArcData/sqlManagedInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "SqlManagedInstances_List",
		Resource:    "Microsoft.AzureArcData",
	},
	"SqlManagedInstances_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureArcData/sqlManagedInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "SqlManagedInstances_ListByResourceGroup",
		Resource:    "Microsoft.AzureArcData",
	},
	"SqlManagedInstances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureArcData/sqlManagedInstances/{{.sqlManagedInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "SqlManagedInstances_Get",
		Resource:    "Microsoft.AzureArcData",
	},
	"SqlServerInstances_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureArcData/sqlServerInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "SqlServerInstances_List",
		Resource:    "Microsoft.AzureArcData",
	},
	"SqlServerInstances_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureArcData/sqlServerInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "SqlServerInstances_ListByResourceGroup",
		Resource:    "Microsoft.AzureArcData",
	},
	"SqlServerInstances_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureArcData/sqlServerInstances/{{.sqlServerInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "SqlServerInstances_Get",
		Resource:    "Microsoft.AzureArcData",
	},
	"DataControllers_ListInSubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.AzureArcData/dataControllers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DataControllers_ListInSubscription",
		Resource:    "Microsoft.AzureArcData",
	},
	"DataControllers_ListInGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureArcData/dataControllers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DataControllers_ListInGroup",
		Resource:    "Microsoft.AzureArcData",
	},
	"DataControllers_GetDataController": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.AzureArcData/dataControllers/{{.dataControllerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-11-01",
		},
		OperationID: "DataControllers_GetDataController",
		Resource:    "Microsoft.AzureArcData",
	},
}
