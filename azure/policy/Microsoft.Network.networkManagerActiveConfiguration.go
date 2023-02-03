package policy

// Microsoft_Network_networkManagerActiveConfiguration policy
var Microsoft_Network_networkManagerActiveConfiguration = map[string]Policy{
	"ListActiveConnectivityConfigurations": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/listActiveConnectivityConfigurations",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ListActiveConnectivityConfigurations",
		Resource:    "Microsoft.Network",
	},
	"ListActiveSecurityAdminRules": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/listActiveSecurityAdminRules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ListActiveSecurityAdminRules",
		Resource:    "Microsoft.Network",
	},
}
