package policy

// Microsoft_CertificateRegistration_CertificateOrdersDiagnostics policy
var Microsoft_CertificateRegistration_CertificateOrdersDiagnostics = map[string]Policy{
	"CertificateOrdersDiagnostics_ListAppServiceCertificateOrderDetectorResponse": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/detectors",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "CertificateOrdersDiagnostics_ListAppServiceCertificateOrderDetectorResponse",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"CertificateOrdersDiagnostics_GetAppServiceCertificateOrderDetectorResponse": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/detectors/{{.detectorName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "CertificateOrdersDiagnostics_GetAppServiceCertificateOrderDetectorResponse",
		Resource:    "Microsoft.CertificateRegistration",
	},
}
