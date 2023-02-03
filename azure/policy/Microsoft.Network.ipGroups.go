package policy

// Microsoft_Network_ipGroups policy
var Microsoft_Network_ipGroups = map[string]Policy{
	"IpGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ipGroups/{{.ipGroupsName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "IpGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"IpGroups_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ipGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "IpGroups_ListByResourceGroup",
		Resource:    "Microsoft.Network",
	},
	"IpGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ipGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "IpGroups_List",
		Resource:    "Microsoft.Network",
	},
}
