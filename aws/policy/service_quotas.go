package policy

import "github.com/securisec/cliam/shared"

// ServiceQuotasPolicies policy
var ServiceQuotasPolicies = map[string]Service{
	"AssociateServiceQuotaTemplate": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.AssociateServiceQuotaTemplate",
		},
		Permission: "AssociateServiceQuotaTemplate",
	},
	"DisassociateServiceQuotaTemplate": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.DisassociateServiceQuotaTemplate",
		},
		Permission: "DisassociateServiceQuotaTemplate",
	},
	"GetAssociationForServiceQuotaTemplate": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.GetAssociationForServiceQuotaTemplate",
		},
		Permission: "GetAssociationForServiceQuotaTemplate",
	},
	"ListRequestedServiceQuotaChangeHistory": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListRequestedServiceQuotaChangeHistory",
		},
		Permission: "ListRequestedServiceQuotaChangeHistory",
	},
	"ListServiceQuotaIncreaseRequestsInTemplate": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListServiceQuotaIncreaseRequestsInTemplate",
		},
		Permission: "ListServiceQuotaIncreaseRequestsInTemplate",
	},
	"ListServices": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListServices",
		},
		Permission: "ListServices",
	},

	// extra
	"GetRequestedServiceQuotaChange": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.GetRequestedServiceQuotaChange",
		},
		Permission:             "GetRequestedServiceQuotaChange",
		IsExtra:                true,
		ExtraComponentBodyKey:  "RequestId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "request_id",
	},
	"ListAWSDefaultServiceQuotas": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListAWSDefaultServiceQuotas",
		},
		Permission:             "ListAWSDefaultServiceQuotas",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceCode",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_code",
	},
	"ListServiceQuotas": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListServiceQuotas",
		},
		Permission:             "ListServiceQuotas",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ServiceCode",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "service_code",
	},
	"ListTagsForResource": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
