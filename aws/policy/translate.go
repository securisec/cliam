package policy

import "github.com/securisec/cliam/shared"

var TranslatePolicies = map[string]Service{
	"ListParallelData": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.ListParallelData",
		},
		Permission: "ListParallelData",
	},
	"ListTerminologies": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.ListTerminologies",
		},
		Permission: "ListTerminologies",
	},
	"ListTextTranslationJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.ListTextTranslationJobs",
		},
		Permission: "ListTextTranslationJobs",
	},

	// extra
	"DescribeTextTranslationJob": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.DescribeTextTranslationJob",
		},
		Permission:             "DescribeTextTranslationJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	"GetParallelData": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.GetParallelData",
		},
		Permission:             "GetParallelData",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
	"GetTerminology": {
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSShineFrontendService_20170701.GetTerminology",
		},
		Permission:             "GetTerminology",
		IsExtra:                true,
		ExtraComponentBodyKey:  "Name",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "name",
	},
}
