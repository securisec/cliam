package policy

var Microsoft_Storage_file = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "FileServices_List",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/{{.FileServicesName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "FileServices_GetServiceProperties",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "FileShares_List",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares/{{.shareName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "FileShares_Get",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares/{{.shareName}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "FileShares_Restore",
		Resource:    "Microsoft.Storage",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares/{{.shareName}}/lease",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-09-01",
		},
		OperationID: "FileShares_Lease",
		Resource:    "Microsoft.Storage",
	},
}
