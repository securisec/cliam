package policy

// Microsoft_RecoveryServices_vaultusages policy
var Microsoft_RecoveryServices_vaultusages = map[string]Policy{
	"Usages_ListByVaults": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/usages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "Usages_ListByVaults",
		Resource:    "Microsoft.RecoveryServices",
	},
}
