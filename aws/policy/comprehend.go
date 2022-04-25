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

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeDocumentClassificationJob",
		},
		Permission:             "DescribeDocumentClassificationJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeDocumentClassifier",
		},
		Permission:             "DescribeDocumentClassifier",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DocumentClassifierArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "document_classifier_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeDominantLanguageDetectionJob",
		},
		Permission:             "DescribeDominantLanguageDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeEndpoint",
		},
		Permission:             "DescribeEndpoint",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EndpointArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "endpoint_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeEntitiesDetectionJob",
		},
		Permission:             "DescribeEntitiesDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeEntityRecognizer",
		},
		Permission:             "DescribeEntityRecognizer",
		IsExtra:                true,
		ExtraComponentBodyKey:  "EntityRecognizerArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "entity_recognizer_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeEventsDetectionJob",
		},
		Permission:             "DescribeEventsDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeKeyPhrasesDetectionJob",
		},
		Permission:             "DescribeKeyPhrasesDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribePiiEntitiesDetectionJob",
		},
		Permission:             "DescribePiiEntitiesDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeResourcePolicy",
		},
		Permission:             "DescribeResourcePolicy",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeSentimentDetectionJob",
		},
		Permission:             "DescribeSentimentDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeTargetedSentimentDetectionJob",
		},
		Permission:             "DescribeTargetedSentimentDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DescribeTopicsDetectionJob",
		},
		Permission:             "DescribeTopicsDetectionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "JobId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "job_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
