package policy

// Microsoft_Network_azureFirewall policy
var Microsoft_Network_azureFirewall = map[string]Policy{
	"AzureFirewalls_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/azureFirewalls/{{.azureFirewallName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AzureFirewalls_Get",
		Resource:    "Microsoft.Network",
	},
	"AzureFirewalls_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/azureFirewalls",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AzureFirewalls_List",
		Resource:    "Microsoft.Network",
	},
	"AzureFirewalls_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/azureFirewalls",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AzureFirewalls_ListAll",
		Resource:    "Microsoft.Network",
	},
	"AzureFirewalls_ListLearnedPrefixes": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/azureFirewalls/{{.azureFirewallName}}/learnedIPPrefixes",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AzureFirewalls_ListLearnedPrefixes",
		Resource:    "Microsoft.Network",
	},
}
