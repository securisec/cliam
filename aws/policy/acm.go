package policy

import "github.com/securisec/cliam/shared"

var ACMPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "CertificateManager.ListCertificates",
		},
		Permission: "ListCertificates",
	},

	// extra
	{
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
	{
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
	{
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
