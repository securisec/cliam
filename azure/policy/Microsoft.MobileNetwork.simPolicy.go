package policy

// Microsoft_MobileNetwork_simPolicy policy
var Microsoft_MobileNetwork_simPolicy = map[string]Policy{
	"SimPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/simPolicies/{{.simPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SimPolicies_Get",
		Resource:    "Microsoft.MobileNetwork",
	},
	"SimPolicies_ListByMobileNetwork": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.MobileNetwork/mobileNetworks/{{.mobileNetworkName}}/simPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-11-01",
		},
		OperationID: "SimPolicies_ListByMobileNetwork",
		Resource:    "Microsoft.MobileNetwork",
	},
}
