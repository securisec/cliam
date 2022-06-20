package policy

var Microsoft_RecoveryServices_replicationusages = map[string]Policy{
	"ReplicationUsages_List": {
		Path:   "/Subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.RecoveryServices/vaults/{{.vaultName}}/replicationUsages",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-04-01",
		},
		OperationID: "ReplicationUsages_List",
		Resource:    "Microsoft.RecoveryServices",
	},
}
