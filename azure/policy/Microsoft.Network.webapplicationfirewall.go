package policy

var Microsoft_Network_webapplicationfirewall = map[string]Policy{
	"Policies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/frontDoorWebApplicationFirewallPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-01",
		},
		OperationID: "Policies_List",
		Resource:    "Microsoft.Network",
	},
	"Policies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallPolicies/{{.policyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-01",
		},
		OperationID: "Policies_Get",
		Resource:    "Microsoft.Network",
	},
	"ManagedRuleSets_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/FrontDoorWebApplicationFirewallManagedRuleSets",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-11-01",
		},
		OperationID: "ManagedRuleSets_List",
		Resource:    "Microsoft.Network",
	},
}
