package policy

// Microsoft_HybridCompute_HybridCompute policy
var Microsoft_HybridCompute_HybridCompute = map[string]Policy{
	"Machines_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "Machines_Get",
		Resource:    "Microsoft.HybridCompute",
	},
	"Machines_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "Machines_ListByResourceGroup",
		Resource:    "Microsoft.HybridCompute",
	},
	"Machines_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridCompute/machines",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "Machines_ListBySubscription",
		Resource:    "Microsoft.HybridCompute",
	},
	"MachineExtensions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/extensions/{{.extensionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "MachineExtensions_Get",
		Resource:    "Microsoft.HybridCompute",
	},
	"MachineExtensions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/extensions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "MachineExtensions_List",
		Resource:    "Microsoft.HybridCompute",
	},
	"UpgradeExtensions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.HybridCompute/machines/{{.machineName}}/upgradeExtensions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "UpgradeExtensions",
		Resource:    "Microsoft.HybridCompute",
	},
	"ExtensionMetadata_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridCompute/locations/{{.location}}/publishers/{{.publisher}}/extensionTypes/{{.extensionType}}/versions/{{.version}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "ExtensionMetadata_Get",
		Resource:    "Microsoft.HybridCompute",
	},
	"ExtensionMetadata_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.HybridCompute/locations/{{.location}}/publishers/{{.publisher}}/extensionTypes/{{.extensionType}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "ExtensionMetadata_List",
		Resource:    "Microsoft.HybridCompute",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.HybridCompute/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-10",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.HybridCompute",
	},
}
