package policy

var Microsoft_Authorization_RoleAssignmentSchedule = map[string]Policy{
	"RoleAssignmentSchedules_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentSchedules/{{.roleAssignmentScheduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentSchedules_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignmentSchedules_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentSchedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentSchedules_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
}
