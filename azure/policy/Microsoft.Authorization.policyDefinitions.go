package policy

var Microsoft_Authorization_policyDefinitions = map[string]Policy{
	"PolicyDefinitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policyDefinitions/{{.policyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_Get",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyDefinitions_GetBuiltIn": {
		Path:   "/providers/Microsoft.Authorization/policyDefinitions/{{.policyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_GetBuiltIn",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyDefinitions_GetAtManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policyDefinitions/{{.policyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_GetAtManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyDefinitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_List",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyDefinitions_ListBuiltIn": {
		Path:   "/providers/Microsoft.Authorization/policyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_ListBuiltIn",
		Resource:    "Microsoft.Authorization",
	},
	"PolicyDefinitions_ListByManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_ListByManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
}
