package policy

// Microsoft_RecoveryServices_replicationusages policy
var Microsoft_RecoveryServices_replicationusages = map[string]Policy{
	"ReplicationUsages_List": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/replicationUsages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2023-01-01",
		},
		OperationID: "ReplicationUsages_List",
		Resource:    "Microsoft.RecoveryServices",
	},
}
