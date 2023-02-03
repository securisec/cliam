package policy

// Microsoft_Synapse_azureADOnlyAuthentication policy
var Microsoft_Synapse_azureADOnlyAuthentication = map[string]Policy{
	"AzureADOnlyAuthentications_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/azureADOnlyAuthentications/{{.azureADOnlyAuthenticationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "AzureADOnlyAuthentications_Get",
		Resource:    "Microsoft.Synapse",
	},
	"AzureADOnlyAuthentications_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/azureADOnlyAuthentications",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "AzureADOnlyAuthentications_List",
		Resource:    "Microsoft.Synapse",
	},
}
