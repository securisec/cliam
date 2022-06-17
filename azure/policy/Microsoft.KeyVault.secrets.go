package policy

    var Microsoft_KeyVault_secrets = []Policy{
        {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/secrets/{{.secretName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Secrets_Get",
    Resource:       "Microsoft.KeyVault",
},{
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.KeyVault/vaults/{{.vaultName}}/secrets",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-10-01",
    },
	OperationID:    "Secrets_List",
    Resource:       "Microsoft.KeyVault",
},
    }
    