package policy

var Microsoft_Authorization_RoleEligibilityScheduleInstance = map[string]Policy{
	"RoleEligibilityScheduleInstances_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleInstances_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
	"RoleEligibilityScheduleInstances_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/{{.roleEligibilityScheduleInstanceName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilityScheduleInstances_Get",
		Resource:    "Microsoft.Authorization",
	},
}
