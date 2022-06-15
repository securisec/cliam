package policy

var Microsoft_Storage_blob = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobServices_List",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/{{.BlobServicesName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobServices_GetServiceProperties",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_List",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_Get",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/setLegalHold",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_SetLegalHold",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/clearLegalHold",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_ClearLegalHold",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/immutabilityPolicies/{{.immutabilityPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_GetImmutabilityPolicy",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/immutabilityPolicies/default/lock",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_LockImmutabilityPolicy",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/immutabilityPolicies/default/extend",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_ExtendImmutabilityPolicy",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/lease",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_Lease",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/blobServices/default/containers/{{.containerName}}/migrate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "BlobContainers_ObjectLevelWorm",
		Resource:    "Microsoft.Storage",
	},
}
