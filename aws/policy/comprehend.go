package policy

import "github.com/securisec/cliam/shared"

var ComprehendPolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassificationJobs",
		},
		Permission: "ListDocumentClassificationJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassifierSummaries",
		},
		Permission: "ListDocumentClassifierSummaries",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassifiers",
		},
		Permission: "ListDocumentClassifiers",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDominantLanguageDetectionJobs",
		},
		Permission: "ListDominantLanguageDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEndpoints",
		},
		Permission: "ListEndpoints",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntitiesDetectionJobs",
		},
		Permission: "ListEntitiesDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntityRecognizerSummaries",
		},
		Permission: "ListEntityRecognizerSummaries",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntityRecognizers",
		},
		Permission: "ListEntityRecognizers",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEventsDetectionJobs",
		},
		Permission: "ListEventsDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListKeyPhrasesDetectionJobs",
		},
		Permission: "ListKeyPhrasesDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListPiiEntitiesDetectionJobs",
		},
		Permission: "ListPiiEntitiesDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListSentimentDetectionJobs",
		},
		Permission: "ListSentimentDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTargetedSentimentDetectionJobs",
		},
		Permission: "ListTargetedSentimentDetectionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTopicsDetectionJobs",
		},
		Permission: "ListTopicsDetectionJobs",
	},
}
