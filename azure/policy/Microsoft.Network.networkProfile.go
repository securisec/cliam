package policy

// Microsoft_Network_networkProfile policy
var Microsoft_Network_networkProfile = map[string]Policy{
	"NetworkProfiles_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkProfiles/{{.networkProfileName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkProfiles_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkProfiles_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkProfiles_ListAll",
		Resource:    "Microsoft.Network",
	},
	"NetworkProfiles_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkProfiles",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkProfiles_List",
		Resource:    "Microsoft.Network",
	},
}
