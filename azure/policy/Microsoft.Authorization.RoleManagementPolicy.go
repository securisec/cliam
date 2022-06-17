package policy

var Microsoft_Authorization_RoleManagementPolicy = map[string]Policy{
	"RoleManagementPolicies_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleManagementPolicies/{{.roleManagementPolicyName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleManagementPolicies_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleManagementPolicies_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleManagementPolicies",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleManagementPolicies_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
}
