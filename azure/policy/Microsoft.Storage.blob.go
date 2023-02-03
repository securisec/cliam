package policy

// Microsoft_Storage_blob policy
var Microsoft_Storage_blob = map[string]Policy{
	"BlobServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobServices_List",
		Resource:    "Microsoft.Storage",
	},
	"BlobServices_GetServiceProperties": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/{{.BlobServicesName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobServices_GetServiceProperties",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_List",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_Get",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_SetLegalHold": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/setLegalHold",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_SetLegalHold",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_ClearLegalHold": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/clearLegalHold",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_ClearLegalHold",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_GetImmutabilityPolicy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/immutabilityPolicies/{{.immutabilityPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_GetImmutabilityPolicy",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_LockImmutabilityPolicy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/immutabilityPolicies/default/lock",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_LockImmutabilityPolicy",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_ExtendImmutabilityPolicy": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/immutabilityPolicies/default/extend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_ExtendImmutabilityPolicy",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_Lease": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/lease",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_Lease",
		Resource:    "Microsoft.Storage",
	},
	"BlobContainers_ObjectLevelWorm": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/migrate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "BlobContainers_ObjectLevelWorm",
		Resource:    "Microsoft.Storage",
	},
}
