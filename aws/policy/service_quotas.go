package policy

import "github.com/securisec/cliam/shared"

var ServiceQuotasPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.AssociateServiceQuotaTemplate",
		},
		Permission: "AssociateServiceQuotaTemplate",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.DisassociateServiceQuotaTemplate",
		},
		Permission: "DisassociateServiceQuotaTemplate",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.GetAssociationForServiceQuotaTemplate",
		},
		Permission: "GetAssociationForServiceQuotaTemplate",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListRequestedServiceQuotaChangeHistory",
		},
		Permission: "ListRequestedServiceQuotaChangeHistory",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListServiceQuotaIncreaseRequestsInTemplate",
		},
		Permission: "ListServiceQuotaIncreaseRequestsInTemplate",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "ServiceQuotasV20190624.ListServices",
		},
		Permission: "ListServices",
	},
}
