package policy

// Microsoft_Synapse_firewallRule policy
var Microsoft_Synapse_firewallRule = map[string]Policy{
	"IpFirewallRules_ListByWorkspace": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/firewallRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IpFirewallRules_ListByWorkspace",
		Resource:    "Microsoft.Synapse",
	},
	"IpFirewallRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/firewallRules/{{.ruleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IpFirewallRules_Get",
		Resource:    "Microsoft.Synapse",
	},
	"IpFirewallRules_ReplaceAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Synapse/workspaces/{{.workspaceName}}/replaceAllIpFirewallRules",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "IpFirewallRules_ReplaceAll",
		Resource:    "Microsoft.Synapse",
	},
}
