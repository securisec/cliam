package policy

var Microsoft_Authorization_RoleAssignmentScheduleInstance = map[string]Policy{
	"RoleAssignmentScheduleInstances_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentScheduleInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentScheduleInstances_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignmentScheduleInstances_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentScheduleInstances/{{.roleAssignmentScheduleInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentScheduleInstances_Get",
		Resource:    "Microsoft.Authorization",
	},
}
