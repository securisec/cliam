package policy

import "github.com/securisec/cliam/shared"

var FraudDetectorPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.DescribeModelVersions",
		},
		Permission: "DescribeModelVersions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetBatchImportJobs",
		},
		Permission: "GetBatchImportJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetBatchPredictionJobs",
		},
		Permission: "GetBatchPredictionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetDetectors",
		},
		Permission: "GetDetectors",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetEntityTypes",
		},
		Permission: "GetEntityTypes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetEventTypes",
		},
		Permission: "GetEventTypes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetExternalModels",
		},
		Permission: "GetExternalModels",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetLabels",
		},
		Permission: "GetLabels",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetModels",
		},
		Permission: "GetModels",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetOutcomes",
		},
		Permission: "GetOutcomes",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetVariables",
		},
		Permission: "GetVariables",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.ListEventPredictions",
		},
		Permission: "ListEventPredictions",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.DescribeDetector",
		},
		Permission:             "DescribeDetector",
		IsExtra:                true,
		ExtraComponentBodyKey:  "detectorId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetDeleteEventsByEventTypeStatus",
		},
		Permission:             "GetDeleteEventsByEventTypeStatus",
		IsExtra:                true,
		ExtraComponentBodyKey:  "eventTypeName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "event_type_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.GetRules",
		},
		Permission:             "GetRules",
		IsExtra:                true,
		ExtraComponentBodyKey:  "detectorId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "detector_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "AWSHawksNestServiceFacade.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceARN",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
