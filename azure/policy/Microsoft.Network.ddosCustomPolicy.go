package policy

// Microsoft_Network_ddosCustomPolicy policy
var Microsoft_Network_ddosCustomPolicy = map[string]Policy{
	"DdosCustomPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ddosCustomPolicies/{{.ddosCustomPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DdosCustomPolicies_Get",
		Resource:    "Microsoft.Network",
	},
}
