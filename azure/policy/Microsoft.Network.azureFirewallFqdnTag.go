package policy

// Microsoft_Network_azureFirewallFqdnTag policy
var Microsoft_Network_azureFirewallFqdnTag = map[string]Policy{
	"AzureFirewallFqdnTags_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/azureFirewallFqdnTags",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AzureFirewallFqdnTags_ListAll",
		Resource:    "Microsoft.Network",
	},
}
