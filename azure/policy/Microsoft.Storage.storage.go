package policy

var Microsoft_Storage_storage = map[string]Policy{
	"Operations_List": {
		Path:   "/providers/Microsoft.Storage/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Operations_List",
		Resource:    "Microsoft.Storage",
	},
	"Skus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Storage/skus",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Skus_List",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_CheckNameAvailability": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Storage/checkNameAvailability",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_CheckNameAvailability",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_GetProperties": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_GetProperties",
		Resource:    "Microsoft.Storage",
	},
	"DeletedAccounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Storage/deletedAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "DeletedAccounts_List",
		Resource:    "Microsoft.Storage",
	},
	"DeletedAccounts_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Storage/locations/{{.location}}/deletedAccounts/{{.deletedAccountName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "DeletedAccounts_Get",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Storage/storageAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_List",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_ListByResourceGroup",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_ListKeys",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_RegenerateKey": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/regenerateKey",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_RegenerateKey",
		Resource:    "Microsoft.Storage",
	},
	"Usages_ListByLocation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Storage/locations/{{.location}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "Usages_ListByLocation",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_ListAccountSAS": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/ListAccountSas",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_ListAccountSAS",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_ListServiceSAS": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/ListServiceSas",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_ListServiceSAS",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_Failover": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/failover",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_Failover",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_HierarchicalNamespaceMigration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/hnsonmigration",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_HierarchicalNamespaceMigration",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_AbortHierarchicalNamespaceMigration": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourcegroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/aborthnsonmigration",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_AbortHierarchicalNamespaceMigration",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_RestoreBlobRanges": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/restoreBlobRanges",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_RestoreBlobRanges",
		Resource:    "Microsoft.Storage",
	},
	"ManagementPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/managementPolicies/{{.managementPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "ManagementPolicies_Get",
		Resource:    "Microsoft.Storage",
	},
	"BlobInventoryPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/inventoryPolicies/{{.blobInventoryPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobInventoryPolicies_Get",
		Resource:    "Microsoft.Storage",
	},
	"BlobInventoryPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/inventoryPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobInventoryPolicies_List",
		Resource:    "Microsoft.Storage",
	},
	"PrivateEndpointConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "PrivateEndpointConnections_List",
		Resource:    "Microsoft.Storage",
	},
	"PrivateEndpointConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Storage",
	},
	"PrivateLinkResources_ListByStorageAccount": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/privateLinkResources",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "PrivateLinkResources_ListByStorageAccount",
		Resource:    "Microsoft.Storage",
	},
	"ObjectReplicationPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/objectReplicationPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "ObjectReplicationPolicies_List",
		Resource:    "Microsoft.Storage",
	},
	"ObjectReplicationPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/objectReplicationPolicies/{{.objectReplicationPolicyId}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "ObjectReplicationPolicies_Get",
		Resource:    "Microsoft.Storage",
	},
	"StorageAccounts_RevokeUserDelegationKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/revokeUserDelegationKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "StorageAccounts_RevokeUserDelegationKeys",
		Resource:    "Microsoft.Storage",
	},
	"LocalUsers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/localUsers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "LocalUsers_List",
		Resource:    "Microsoft.Storage",
	},
	"LocalUsers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/localUsers/{{.username}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "LocalUsers_Get",
		Resource:    "Microsoft.Storage",
	},
	"LocalUsers_ListKeys": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/localUsers/{{.username}}/listKeys",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "LocalUsers_ListKeys",
		Resource:    "Microsoft.Storage",
	},
	"LocalUsers_RegeneratePassword": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/localUsers/{{.username}}/regeneratePassword",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "LocalUsers_RegeneratePassword",
		Resource:    "Microsoft.Storage",
	},
	"EncryptionScopes_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/encryptionScopes/{{.encryptionScopeName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "EncryptionScopes_Get",
		Resource:    "Microsoft.Storage",
	},
	"EncryptionScopes_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/encryptionScopes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "EncryptionScopes_List",
		Resource:    "Microsoft.Storage",
	},
}
