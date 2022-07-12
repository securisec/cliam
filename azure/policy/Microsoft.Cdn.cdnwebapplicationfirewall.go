package policy

    var Microsoft_Cdn_cdnwebapplicationfirewall = map[string]Policy{
        "Policies_List": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/cdnWebApplicationFirewallPolicies",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Policies_List",
    Resource:       "Microsoft.Cdn",
},
"Policies_Get": {
    Path: "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Cdn/cdnWebApplicationFirewallPolicies/{{.policyName}}",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "Policies_Get",
    Resource:       "Microsoft.Cdn",
},
"ManagedRuleSets_List": {
    Path: "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Cdn/cdnWebApplicationFirewallManagedRuleSets",
	Method: "GET",
	QueryValues:   map[string]string{
        "api-version": "2021-06-01",
    },
	OperationID:    "ManagedRuleSets_List",
    Resource:       "Microsoft.Cdn",
},
    }
    