package policy

// Microsoft_App_AvailableWorkloadProfiles policy
var Microsoft_App_AvailableWorkloadProfiles = map[string]Policy{
	"AvailableWorkloadProfiles_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.App/locations/{{.location}}/availableManagedEnvironmentsWorkloadProfileTypes",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-10-01",
		},
		OperationID: "AvailableWorkloadProfiles_Get",
		Resource:    "Microsoft.App",
	},
}
