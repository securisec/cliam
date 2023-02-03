package policy

// Microsoft_DataLakeStore_account policy
var Microsoft_DataLakeStore_account = map[string]Policy{
	"Accounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeStore/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_List",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Accounts_EnableKeyVault": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/enableKeyVault",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_EnableKeyVault",
		Resource:    "Microsoft.DataLakeStore",
	},
	"FirewallRules_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/firewallRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "FirewallRules_ListByAccount",
		Resource:    "Microsoft.DataLakeStore",
	},
	"FirewallRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/firewallRules/{{.firewallRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "FirewallRules_Get",
		Resource:    "Microsoft.DataLakeStore",
	},
	"VirtualNetworkRules_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/virtualNetworkRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "VirtualNetworkRules_ListByAccount",
		Resource:    "Microsoft.DataLakeStore",
	},
	"VirtualNetworkRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/virtualNetworkRules/{{.virtualNetworkRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "VirtualNetworkRules_Get",
		Resource:    "Microsoft.DataLakeStore",
	},
	"TrustedIdProviders_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/trustedIdProviders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "TrustedIdProviders_ListByAccount",
		Resource:    "Microsoft.DataLakeStore",
	},
	"TrustedIdProviders_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.DataLakeStore/accounts/{{.accountName}}/trustedIdProviders/{{.trustedIdProviderName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "TrustedIdProviders_Get",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.DataLakeStore/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Locations_GetCapability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeStore/locations/{{.location}}/capability",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Locations_GetCapability",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Locations_GetUsage": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeStore/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Locations_GetUsage",
		Resource:    "Microsoft.DataLakeStore",
	},
	"Accounts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.DataLakeStore/locations/{{.location}}/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2016-11-01",
		},
		OperationID: "Accounts_CheckNameAvailability",
		Resource:    "Microsoft.DataLakeStore",
	},
}
