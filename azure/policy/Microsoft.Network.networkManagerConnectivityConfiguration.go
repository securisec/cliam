package policy

// Microsoft_Network_networkManagerConnectivityConfiguration policy
var Microsoft_Network_networkManagerConnectivityConfiguration = map[string]Policy{
	"ConnectivityConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/connectivityConfigurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectivityConfigurations_Get",
		Resource:    "Microsoft.Network",
	},
	"ConnectivityConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/connectivityConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ConnectivityConfigurations_List",
		Resource:    "Microsoft.Network",
	},
}
