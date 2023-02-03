package policy

// Microsoft_Network_networkManagerEffectiveConfiguration policy
var Microsoft_Network_networkManagerEffectiveConfiguration = map[string]Policy{
	"ListNetworkManagerEffectiveConnectivityConfigurations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/listNetworkManagerEffectiveConnectivityConfigurations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ListNetworkManagerEffectiveConnectivityConfigurations",
		Resource:    "Microsoft.Network",
	},
	"ListNetworkManagerEffectiveSecurityAdminRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/virtualNetworks/{{.virtualNetworkName}}/listNetworkManagerEffectiveSecurityAdminRules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ListNetworkManagerEffectiveSecurityAdminRules",
		Resource:    "Microsoft.Network",
	},
}
