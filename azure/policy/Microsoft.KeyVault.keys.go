package policy

var Microsoft_KeyVault_keys = map[string]Policy{
	"Keys_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/keys/{{.keyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Keys_Get",
		Resource:    "Microsoft.KeyVault",
	},
	"Keys_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/keys",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Keys_List",
		Resource:    "Microsoft.KeyVault",
	},
	"Keys_GetVersion": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/keys/{{.keyName}}/versions/{{.keyVersion}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Keys_GetVersion",
		Resource:    "Microsoft.KeyVault",
	},
	"Keys_ListVersions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/keys/{{.keyName}}/versions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-10-01",
		},
		OperationID: "Keys_ListVersions",
		Resource:    "Microsoft.KeyVault",
	},
}
