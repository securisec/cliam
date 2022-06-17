package policy

var Microsoft_Authorization_policySetDefinitions = map[string]Policy{
	"PolicySetDefinitions_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policySetDefinitions/{{.policySetDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_Get",
		Resource:    "Microsoft.Authorization",
	},
	"PolicySetDefinitions_GetBuiltIn": {
		Path:   "/providers/Microsoft.Authorization/policySetDefinitions/{{.policySetDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_GetBuiltIn",
		Resource:    "Microsoft.Authorization",
	},
	"PolicySetDefinitions_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policySetDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_List",
		Resource:    "Microsoft.Authorization",
	},
	"PolicySetDefinitions_ListBuiltIn": {
		Path:   "/providers/Microsoft.Authorization/policySetDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_ListBuiltIn",
		Resource:    "Microsoft.Authorization",
	},
	"PolicySetDefinitions_GetAtManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policySetDefinitions/{{.policySetDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_GetAtManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
	"PolicySetDefinitions_ListByManagementGroup": {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policySetDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_ListByManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
}
