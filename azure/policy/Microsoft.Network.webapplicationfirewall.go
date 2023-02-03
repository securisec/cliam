package policy

// Microsoft_Network_webapplicationfirewall policy
var Microsoft_Network_webapplicationfirewall = map[string]Policy{
	"WebApplicationFirewallPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "WebApplicationFirewallPolicies_List",
		Resource:    "Microsoft.Network",
	},
	"WebApplicationFirewallPolicies_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "WebApplicationFirewallPolicies_ListAll",
		Resource:    "Microsoft.Network",
	},
	"WebApplicationFirewallPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{{.policyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "WebApplicationFirewallPolicies_Get",
		Resource:    "Microsoft.Network",
	},
}
