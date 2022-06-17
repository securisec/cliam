package policy

var Microsoft_Authorization_RoleAssignmentScheduleRequest = map[string]Policy{
	"RoleAssignmentScheduleRequests_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{{.roleAssignmentScheduleRequestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentScheduleRequests_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignmentScheduleRequests_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentScheduleRequests_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignmentScheduleRequests_Cancel": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{{.roleAssignmentScheduleRequestName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentScheduleRequests_Cancel",
		Resource:    "Microsoft.Authorization",
	},
	"RoleAssignmentScheduleRequests_Validate": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleAssignmentScheduleRequests/{{.roleAssignmentScheduleRequestName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleAssignmentScheduleRequests_Validate",
		Resource:    "Microsoft.Authorization",
	},
}
