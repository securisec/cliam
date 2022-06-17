package policy

var Microsoft_Authorization_policyDefinitions = []Policy{
	{
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policyDefinitions/{{.policyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_Get",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Authorization/policyDefinitions/{{.policyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_GetBuiltIn",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policyDefinitions/{{.policyDefinitionName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_GetAtManagementGroup",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.Authorization/policyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_List",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Authorization/policyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_ListBuiltIn",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/providers/Microsoft.Management/managementGroups/{{.managementGroupId}}/providers/Microsoft.Authorization/policyDefinitions",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-06-01",
		},
		OperationID: "PolicyDefinitions_ListByManagementGroup",
		Resource:    "Microsoft.Authorization",
	},
}
