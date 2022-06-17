package policy

var Microsoft_Authorization_RoleEligibilityScheduleRequest = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{{.roleEligibilityScheduleRequestName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_Get",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_ListForScope",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{{.roleEligibilityScheduleRequestName}}/cancel",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_Cancel",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleRequests/{{.roleEligibilityScheduleRequestName}}/validate",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleRequests_Validate",
		Resource:    "Microsoft.Authorization",
	},
}
