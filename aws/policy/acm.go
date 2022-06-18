package policy

import "github.com/securisec/cliam/shared"

var ACMPolicies = map[string]Service{
	"ListCertificates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CertificateManager.ListCertificates",
		},
		Permission: "ListCertificates",
	},

	// extra
	"DescribeCertificate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CertificateManager.DescribeCertificate",
		},
		Permission:             "DescribeCertificate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CertificateArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "certificate_arn",
	},
	"GetCertificate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CertificateManager.GetCertificate",
		},
		Permission:             "GetCertificate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CertificateArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "certificate_arn",
	},
	"ListTagsForCertificate": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CertificateManager.ListTagsForCertificate",
		},
		Permission:             "ListTagsForCertificate",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CertificateArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "certificate_arn",
	},
}
