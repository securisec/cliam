package policy

// Microsoft_DomainRegistration_DomainRegistrationProvider policy
var Microsoft_DomainRegistration_DomainRegistrationProvider = map[string]Policy{
	"DomainRegistrationProvider_ListOperations": {
		Path:   "/providers/Microsoft.DomainRegistration/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "DomainRegistrationProvider_ListOperations",
		Resource:    "Microsoft.DomainRegistration",
	},
}
