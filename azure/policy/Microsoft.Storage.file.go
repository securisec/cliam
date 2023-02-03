package policy

// Microsoft_Storage_file policy
var Microsoft_Storage_file = map[string]Policy{
	"FileServices_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "FileServices_List",
		Resource:    "Microsoft.Storage",
	},
	"FileServices_GetServiceProperties": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/{{.FileServicesName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "FileServices_GetServiceProperties",
		Resource:    "Microsoft.Storage",
	},
	"FileShares_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "FileShares_List",
		Resource:    "Microsoft.Storage",
	},
	"FileShares_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares/{{.shareName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "FileShares_Get",
		Resource:    "Microsoft.Storage",
	},
	"FileShares_Restore": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares/{{.shareName}}/restore",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "FileShares_Restore",
		Resource:    "Microsoft.Storage",
	},
	"FileShares_Lease": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Storage/storageAccounts/{{.accountName}}/fileServices/default/shares/{{.shareName}}/lease",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-09-01",
		},
		OperationID: "FileShares_Lease",
		Resource:    "Microsoft.Storage",
	},
}
