package policy

var Microsoft_Authorization_policySetDefinitions = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policySetDefinitions/{{.policySetDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_Get",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Authorization/policySetDefinitions/{{.policySetDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_GetBuiltIn",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policySetDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_List",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Authorization/policySetDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_ListBuiltIn",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policySetDefinitions/{{.policySetDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_GetAtManagementGroup",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policySetDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicySetDefinitions_ListByManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
}
