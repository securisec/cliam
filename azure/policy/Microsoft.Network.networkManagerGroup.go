package policy

// Microsoft_Network_networkManagerGroup policy
var Microsoft_Network_networkManagerGroup = map[string]Policy{
	"NetworkGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/networkGroups/{{.networkGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/networkGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkGroups_List",
		Resource:    "Microsoft.Network",
	},
	"StaticMembers_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/networkGroups/{{.networkGroupName}}/staticMembers/{{.staticMemberName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "StaticMembers_Get",
		Resource:    "Microsoft.Network",
	},
	"StaticMembers_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/networkGroups/{{.networkGroupName}}/staticMembers",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "StaticMembers_List",
		Resource:    "Microsoft.Network",
	},
}
