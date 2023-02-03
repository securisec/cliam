package policy

// Microsoft_Network_networkManager policy
var Microsoft_Network_networkManager = map[string]Policy{
	"NetworkManagers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkManagers_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkManagerCommits_Post": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/commit",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkManagerCommits_Post",
		Resource:    "Microsoft.Network",
	},
	"NetworkManagerDeploymentStatus_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/listDeploymentStatus",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkManagerDeploymentStatus_List",
		Resource:    "Microsoft.Network",
	},
	"NetworkManagers_ListBySubscription": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkManagers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkManagers_ListBySubscription",
		Resource:    "Microsoft.Network",
	},
	"NetworkManagers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkManagers_List",
		Resource:    "Microsoft.Network",
	},
}
