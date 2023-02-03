package policy

// Microsoft_CertificateRegistration_AppServiceCertificateOrders policy
var Microsoft_CertificateRegistration_AppServiceCertificateOrders = map[string]Policy{
	"AppServiceCertificateOrders_List": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CertificateRegistration/certificateOrders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_List",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_ValidatePurchaseInformation": {
		Path:   "/subscriptions/{{.subscriptionId}}/providers/Microsoft.CertificateRegistration/validateCertificateRegistrationInformation",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_ValidatePurchaseInformation",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_ListByResourceGroup": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_ListByResourceGroup",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_Get": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_Get",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_ListCertificates": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/certificates",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_ListCertificates",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_GetCertificate": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/certificates/{{.name}}",
		Method: "GET",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_GetCertificate",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_Reissue": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/reissue",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_Reissue",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_Renew": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/renew",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_Renew",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_ResendEmail": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/resendEmail",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_ResendEmail",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_ResendRequestEmails": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/resendRequestEmails",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_ResendRequestEmails",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_RetrieveSiteSeal": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/retrieveSiteSeal",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_RetrieveSiteSeal",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_VerifyDomainOwnership": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.certificateOrderName}}/verifyDomainOwnership",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_VerifyDomainOwnership",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_RetrieveCertificateActions": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.name}}/retrieveCertificateActions",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_RetrieveCertificateActions",
		Resource:    "Microsoft.CertificateRegistration",
	},
	"AppServiceCertificateOrders_RetrieveCertificateEmailHistory": {
		Path:   "/subscriptions/{{.subscriptionId}}/resourceGroups/{{.resourceGroupName}}/providers/Microsoft.CertificateRegistration/certificateOrders/{{.name}}/retrieveEmailHistory",
		Method: "POST",
		QueryValues: map[string]string{
			"api-version": "2022-03-01",
		},
		OperationID: "AppServiceCertificateOrders_RetrieveCertificateEmailHistory",
		Resource:    "Microsoft.CertificateRegistration",
	},
}
