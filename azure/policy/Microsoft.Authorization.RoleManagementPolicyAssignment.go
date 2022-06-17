package policy

var Microsoft_Authorization_RoleManagementPolicyAssignment = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleManagementPolicyAssignments/{{.roleManagementPolicyAssignmentName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleManagementPolicyAssignments_Get",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleManagementPolicyAssignments",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleManagementPolicyAssignments_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
}
