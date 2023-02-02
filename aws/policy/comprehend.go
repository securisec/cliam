package policy

import "github.com/securisec/cliam/shared"

// ComprehendPolicies policy
var ComprehendPolicies = map[string]Service{
	"DetectEntities": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.DetectEntities",
		},
		Permission: "DetectEntities",
	},
	"ListDocumentClassificationJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassificationJobs",
		},
		Permission: "ListDocumentClassificationJobs",
	},
	"ListDocumentClassifierSummaries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassifierSummaries",
		},
		Permission: "ListDocumentClassifierSummaries",
	},
	"ListDocumentClassifiers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDocumentClassifiers",
		},
		Permission: "ListDocumentClassifiers",
	},
	"ListDominantLanguageDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListDominantLanguageDetectionJobs",
		},
		Permission: "ListDominantLanguageDetectionJobs",
	},
	"ListEndpoints": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEndpoints",
		},
		Permission: "ListEndpoints",
	},
	"ListEntitiesDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntitiesDetectionJobs",
		},
		Permission: "ListEntitiesDetectionJobs",
	},
	"ListEntityRecognizerSummaries": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntityRecognizerSummaries",
		},
		Permission: "ListEntityRecognizerSummaries",
	},
	"ListEntityRecognizers": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEntityRecognizers",
		},
		Permission: "ListEntityRecognizers",
	},
	"ListEventsDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListEventsDetectionJobs",
		},
		Permission: "ListEventsDetectionJobs",
	},
	"ListKeyPhrasesDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListKeyPhrasesDetectionJobs",
		},
		Permission: "ListKeyPhrasesDetectionJobs",
	},
	"ListPiiEntitiesDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListPiiEntitiesDetectionJobs",
		},
		Permission: "ListPiiEntitiesDetectionJobs",
	},
	"ListSentimentDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListSentimentDetectionJobs",
		},
		Permission: "ListSentimentDetectionJobs",
	},
	"ListTargetedSentimentDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTargetedSentimentDetectionJobs",
		},
		Permission: "ListTargetedSentimentDetectionJobs",
	},
	"ListTopicsDetectionJobs": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Comprehend_20171127.ListTopicsDetectionJobs",
		},
		Permission: "ListTopicsDetectionJobs",
	},

	// extra
	"DescribeDocumentClassificationJob": {
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
	"DescribeDocumentClassifier": {
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
	"DescribeDominantLanguageDetectionJob": {
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
	"DescribeEndpoint": {
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
	"DescribeEntitiesDetectionJob": {
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
	"DescribeEntityRecognizer": {
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
	"DescribeEventsDetectionJob": {
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
	"DescribeKeyPhrasesDetectionJob": {
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
	"DescribePiiEntitiesDetectionJob": {
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
	"DescribeResourcePolicy": {
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
	"DescribeSentimentDetectionJob": {
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
	"DescribeTargetedSentimentDetectionJob": {
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
	"DescribeTopicsDetectionJob": {
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
	"ListTagsForResource": {
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
