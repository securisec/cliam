package policy

var Microsoft_Authorization_RoleEligibilitySchedule = map[string]Policy{
	"RoleEligibilitySchedules_Get": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilitySchedules/{{.roleEligibilityScheduleName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilitySchedules_Get",
		Resource:    "Microsoft.Authorization",
	},
	"RoleEligibilitySchedules_ListForScope": {
		Path:   "/{{.scope}}/providers/Microsoft.Authorization/roleEligibilitySchedules",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2020-10-01",
		},
		OperationID: "RoleEligibilitySchedules_ListForScope",
		Resource:    "Microsoft.Authorization",
	},
}
