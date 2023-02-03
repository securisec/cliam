package policy

// Microsoft_Network_networkManagerConnection policy
var Microsoft_Network_networkManagerConnection = map[string]Policy{
	"SubscriptionNetworkManagerConnections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkManagerConnections/{{.networkManagerConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SubscriptionNetworkManagerConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"SubscriptionNetworkManagerConnections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkManagerConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SubscriptionNetworkManagerConnections_List",
		Resource:    "Microsoft.Network",
	},
	"ManagementGroupNetworkManagerConnections_Get": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Network/networkManagerConnections/{{.networkManagerConnectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ManagementGroupNetworkManagerConnections_Get",
		Resource:    "Microsoft.Network",
	},
	"ManagementGroupNetworkManagerConnections_List": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Network/networkManagerConnections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "ManagementGroupNetworkManagerConnections_List",
		Resource:    "Microsoft.Network",
	},
}
