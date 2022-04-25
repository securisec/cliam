package policy

import "github.com/securisec/cliam/shared"

var ComprehendPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassificationJobs",
		},
		Permission: "ListDocumentClassificationJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassifierSummaries",
		},
		Permission: "ListDocumentClassifierSummaries",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassifiers",
		},
		Permission: "ListDocumentClassifiers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDominantLanguageDetectionJobs",
		},
		Permission: "ListDominantLanguageDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEndpoints",
		},
		Permission: "ListEndpoints",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntitiesDetectionJobs",
		},
		Permission: "ListEntitiesDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntityRecognizerSummaries",
		},
		Permission: "ListEntityRecognizerSummaries",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntityRecognizers",
		},
		Permission: "ListEntityRecognizers",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEventsDetectionJobs",
		},
		Permission: "ListEventsDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListKeyPhrasesDetectionJobs",
		},
		Permission: "ListKeyPhrasesDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListPiiEntitiesDetectionJobs",
		},
		Permission: "ListPiiEntitiesDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListSentimentDetectionJobs",
		},
		Permission: "ListSentimentDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTargetedSentimentDetectionJobs",
		},
		Permission: "ListTargetedSentimentDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTopicsDetectionJobs",
		},
		Permission: "ListTopicsDetectionJobs",
	},
}
