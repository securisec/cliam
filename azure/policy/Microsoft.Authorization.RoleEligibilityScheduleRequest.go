package policy

var Microsoft_Authorization_RoleEligibilityScheduleRequest = map[string]Policy{
	"RoleEligibilityScheduleRequests_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{{.roleEligibilityScheduleRequestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleEligibilityScheduleRequests_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
	"RoleEligibilityScheduleRequests_Cancel": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{{.roleEligibilityScheduleRequestName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_Cancel",
		Resource:    "Microsoft.Authorization",
	},
	"RoleEligibilityScheduleRequests_Validate": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{{.roleEligibilityScheduleRequestName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_Validate",
		Resource:    "Microsoft.Authorization",
	},
}
