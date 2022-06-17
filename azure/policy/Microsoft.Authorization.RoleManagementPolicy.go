package policy

var Microsoft_Authorization_RoleManagementPolicy = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleManagementPolicies/{{.roleManagementPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleManagementPolicies_Get",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleManagementPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleManagementPolicies_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
}
