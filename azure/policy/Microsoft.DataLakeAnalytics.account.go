package policy

// Microsoft_DataLakeAnalytics_account policy
var Microsoft_DataLakeAnalytics_account = map[string]Policy{
	"Accounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeAnalytics/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_List",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"DataLakeStoreAccounts_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/dataLakeStoreAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "DataLakeStoreAccounts_ListByAccount",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"DataLakeStoreAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/dataLakeStoreAccounts/{{.dataLakeStoreAccountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "DataLakeStoreAccounts_Get",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"StorageAccounts_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/storageAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "StorageAccounts_ListByAccount",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"StorageAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/storageAccounts/{{.storageAccountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "StorageAccounts_Get",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"StorageAccounts_ListStorageContainers": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/storageAccounts/{{.storageAccountName}}/containers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "StorageAccounts_ListStorageContainers",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"StorageAccounts_GetStorageContainer": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/storageAccounts/{{.storageAccountName}}/containers/{{.containerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "StorageAccounts_GetStorageContainer",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"StorageAccounts_ListSasTokens": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/storageAccounts/{{.storageAccountName}}/containers/{{.containerName}}/listSasTokens",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "StorageAccounts_ListSasTokens",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"ComputePolicies_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/computePolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "ComputePolicies_ListByAccount",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"ComputePolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/computePolicies/{{.computePolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "ComputePolicies_Get",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"FirewallRules_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/firewallRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "FirewallRules_ListByAccount",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"FirewallRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeAnalytics/accounts/{{.accountName}}/firewallRules/{{.firewallRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "FirewallRules_Get",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DataLakeAnalytics/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"Locations_GetCapability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeAnalytics/locations/{{.location}}/capability",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Locations_GetCapability",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
	"Accounts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeAnalytics/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_CheckNameAvailability",
		Resource:    "Microsoft.DataLakeAnalytics",
	},
}
