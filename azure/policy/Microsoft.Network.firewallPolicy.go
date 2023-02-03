package policy

// Microsoft_Network_firewallPolicy policy
var Microsoft_Network_firewallPolicy = map[string]Policy{
	"FirewallPolicies_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicies_Get",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicies_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicies_List",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicies_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/firewallPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicies_ListAll",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicyRuleCollectionGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}/ruleCollectionGroups/{{.ruleCollectionGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicyRuleCollectionGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicyRuleCollectionGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}/ruleCollectionGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicyRuleCollectionGroups_List",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicyIdpsSignatures_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}/listIdpsSignatures",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicyIdpsSignatures_List",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicyIdpsSignaturesOverrides_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}/signatureOverrides/default",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicyIdpsSignaturesOverrides_Get",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicyIdpsSignaturesFilterValues_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}/listIdpsFilterOptions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicyIdpsSignaturesFilterValues_List",
		Resource:    "Microsoft.Network",
	},
	"FirewallPolicyIdpsSignaturesOverrides_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/firewallPolicies/{{.firewallPolicyName}}/signatureOverrides",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "FirewallPolicyIdpsSignaturesOverrides_List",
		Resource:    "Microsoft.Network",
	},
}
