package policy

// Microsoft_Purview_purview policy
var Microsoft_Purview_purview = map[string]Policy{
	"Accounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Accounts_ListByResourceGroup",
		Resource:    "Microsoft.Purview",
	},
	"Accounts_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Purview/accounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Accounts_ListBySubscription",
		Resource:    "Microsoft.Purview",
	},
	"Accounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Accounts_Get",
		Resource:    "Microsoft.Purview",
	},
	"Accounts_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}/listkeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Accounts_ListKeys",
		Resource:    "Microsoft.Purview",
	},
	"Accounts_AddRootCollectionAdmin": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}/addRootCollectionAdmin",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Accounts_AddRootCollectionAdmin",
		Resource:    "Microsoft.Purview",
	},
	"DefaultAccounts_Get": {
		Path:   "/providers/Microsoft.Purview/getDefaultAccount",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DefaultAccounts_Get",
		Resource:    "Microsoft.Purview",
	},
	"DefaultAccounts_Set": {
		Path:   "/providers/Microsoft.Purview/setDefaultAccount",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DefaultAccounts_Set",
		Resource:    "Microsoft.Purview",
	},
	"DefaultAccounts_Remove": {
		Path:   "/providers/Microsoft.Purview/removeDefaultAccount",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "DefaultAccounts_Remove",
		Resource:    "Microsoft.Purview",
	},
	"Operations_List": {
		Path:   "/providers/Microsoft.Purview/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Purview",
	},
	"PrivateEndpointConnections_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "PrivateEndpointConnections_ListByAccount",
		Resource:    "Microsoft.Purview",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Purview",
	},
	"PrivateLinkResources_ListByAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "PrivateLinkResources_ListByAccount",
		Resource:    "Microsoft.Purview",
	},
	"PrivateLinkResources_GetByGroupId": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Purview/accounts/{{.accountName}}/privateLinkResources/{{.groupId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "PrivateLinkResources_GetByGroupId",
		Resource:    "Microsoft.Purview",
	},
	"Accounts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Purview/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-07-01",
		},
		OperationID: "Accounts_CheckNameAvailability",
		Resource:    "Microsoft.Purview",
	},
}
