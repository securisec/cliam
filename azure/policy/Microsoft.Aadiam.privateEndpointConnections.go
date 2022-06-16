package policy

var Microsoft_Aadiam_privateEndpointConnections = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd/{{.policyName}}/privateEndpointConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "PrivateEndpointConnections_ListByPolicyName",
		Resource:    "Microsoft.Aadiam",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/microsoft.aadiam/privateLinkForAzureAd/{{.policyName}}/privateEndpointConnections/{{.privateEndpointConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-03-01",
		},
		OperationID: "PrivateEndpointConnections_Get",
		Resource:    "Microsoft.Aadiam",
	},
}
