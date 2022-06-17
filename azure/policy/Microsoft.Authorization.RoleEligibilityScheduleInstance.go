package policy

var Microsoft_Authorization_RoleEligibilityScheduleInstance = []Policy{
	{
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleInstances_ListForScope",
		Resource:    "Microsoft.Authorization",
	}, {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/{{.roleEligibilityScheduleInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleInstances_Get",
		Resource:    "Microsoft.Authorization",
	},
}
