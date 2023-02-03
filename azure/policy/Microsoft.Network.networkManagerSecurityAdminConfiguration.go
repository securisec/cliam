package policy

// Microsoft_Network_networkManagerSecurityAdminConfiguration policy
var Microsoft_Network_networkManagerSecurityAdminConfiguration = map[string]Policy{
	"SecurityAdminConfigurations_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/securityAdminConfigurations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityAdminConfigurations_List",
		Resource:    "Microsoft.Network",
	},
	"SecurityAdminConfigurations_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/securityAdminConfigurations/{{.configurationName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityAdminConfigurations_Get",
		Resource:    "Microsoft.Network",
	},
	"AdminRuleCollections_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/securityAdminConfigurations/{{.configurationName}}/ruleCollections",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AdminRuleCollections_List",
		Resource:    "Microsoft.Network",
	},
	"AdminRuleCollections_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/securityAdminConfigurations/{{.configurationName}}/ruleCollections/{{.ruleCollectionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AdminRuleCollections_Get",
		Resource:    "Microsoft.Network",
	},
	"AdminRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/securityAdminConfigurations/{{.configurationName}}/ruleCollections/{{.ruleCollectionName}}/rules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AdminRules_List",
		Resource:    "Microsoft.Network",
	},
	"AdminRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkManagers/{{.networkManagerName}}/securityAdminConfigurations/{{.configurationName}}/ruleCollections/{{.ruleCollectionName}}/rules/{{.ruleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "AdminRules_Get",
		Resource:    "Microsoft.Network",
	},
}
