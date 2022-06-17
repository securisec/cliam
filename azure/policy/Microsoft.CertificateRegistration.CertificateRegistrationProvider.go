package policy

var Microsoft_CertificateRegistration_CertificateRegistrationProvider = map[string]Policy{
	"CertificateRegistrationProvider_ListOperations": {
		Path:   "/providers/Microsoft.CertificateRegistration/operations",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2021-03-01",
		},
		OperationID: "CertificateRegistrationProvider_ListOperations",
		Resource:    "Microsoft.CertificateRegistration",
	},
}
