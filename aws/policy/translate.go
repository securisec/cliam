package policy

import "github.com/securisec/cliam/shared"

var TranslatePolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.ListParallelData",
		},
		Permission: "ListParallelData",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.ListTerminologies",
		},
		Permission: "ListTerminologies",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.ListTextTranslationJobs",
		},
		Permission: "ListTextTranslationJobs",
	},
}
