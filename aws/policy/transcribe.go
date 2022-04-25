package policy

import "github.com/securisec/cliam/shared"

var TranscribePolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListCallAnalyticsCategories",
		},
		Permission: "ListCallAnalyticsCategories",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListCallAnalyticsJobs",
		},
		Permission: "ListCallAnalyticsJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListLanguageModels",
		},
		Permission: "ListLanguageModels",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListMedicalTranscriptionJobs",
		},
		Permission: "ListMedicalTranscriptionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListMedicalVocabularies",
		},
		Permission: "ListMedicalVocabularies",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListTranscriptionJobs",
		},
		Permission: "ListTranscriptionJobs",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListVocabularies",
		},
		Permission: "ListVocabularies",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListVocabularyFilters",
		},
		Permission: "ListVocabularyFilters",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.DescribeLanguageModel",
		},
		Permission:             "DescribeLanguageModel",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ModelName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "model_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetCallAnalyticsCategory",
		},
		Permission:             "GetCallAnalyticsCategory",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CategoryName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "category_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetCallAnalyticsJob",
		},
		Permission:             "GetCallAnalyticsJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "CallAnalyticsJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "call_analytics_job_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetMedicalTranscriptionJob",
		},
		Permission:             "GetMedicalTranscriptionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "MedicalTranscriptionJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "medical_transcription_job_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetMedicalVocabulary",
		},
		Permission:             "GetMedicalVocabulary",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VocabularyName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vocabulary_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetTranscriptionJob",
		},
		Permission:             "GetTranscriptionJob",
		IsExtra:                true,
		ExtraComponentBodyKey:  "TranscriptionJobName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "transcription_job_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetVocabulary",
		},
		Permission:             "GetVocabulary",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VocabularyName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vocabulary_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.GetVocabularyFilter",
		},
		Permission:             "GetVocabularyFilter",
		IsExtra:                true,
		ExtraComponentBodyKey:  "VocabularyFilterName",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "vocabulary_filter_name",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "Transcribe.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
