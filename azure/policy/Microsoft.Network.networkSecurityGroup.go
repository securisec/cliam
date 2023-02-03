package policy

// Microsoft_Network_networkSecurityGroup policy
var Microsoft_Network_networkSecurityGroup = map[string]Policy{
	"NetworkSecurityGroups_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkSecurityGroups/{{.networkSecurityGroupName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkSecurityGroups_Get",
		Resource:    "Microsoft.Network",
	},
	"NetworkSecurityGroups_ListAll": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Network/networkSecurityGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkSecurityGroups_ListAll",
		Resource:    "Microsoft.Network",
	},
	"NetworkSecurityGroups_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkSecurityGroups",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "NetworkSecurityGroups_List",
		Resource:    "Microsoft.Network",
	},
	"SecurityRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkSecurityGroups/{{.networkSecurityGroupName}}/securityRules/{{.securityRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityRules_Get",
		Resource:    "Microsoft.Network",
	},
	"SecurityRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkSecurityGroups/{{.networkSecurityGroupName}}/securityRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "SecurityRules_List",
		Resource:    "Microsoft.Network",
	},
	"DefaultSecurityRules_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkSecurityGroups/{{.networkSecurityGroupName}}/defaultSecurityRules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DefaultSecurityRules_List",
		Resource:    "Microsoft.Network",
	},
	"DefaultSecurityRules_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.Network/networkSecurityGroups/{{.networkSecurityGroupName}}/defaultSecurityRules/{{.defaultSecurityRuleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-07-01",
		},
		OperationID: "DefaultSecurityRules_Get",
		Resource:    "Microsoft.Network",
	},
}
