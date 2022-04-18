package policy

import "github.com/securisec/cliam/shared"

var TranscribePolicies = []Service{
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListCallAnalyticsCategories",
		},
		Permission: "ListCallAnalyticsCategories",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListCallAnalyticsJobs",
		},
		Permission: "ListCallAnalyticsJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListLanguageModels",
		},
		Permission: "ListLanguageModels",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListMedicalTranscriptionJobs",
		},
		Permission: "ListMedicalTranscriptionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListMedicalVocabularies",
		},
		Permission: "ListMedicalVocabularies",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListTranscriptionJobs",
		},
		Permission: "ListTranscriptionJobs",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListVocabularies",
		},
		Permission: "ListVocabularies",
	},
	{
		Method:   "POST",
		JsonData: `{}`,
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_CONTENT_TYPE,
			aws_X_AMZ_TARGET:           "Transcribe.ListVocabularyFilters",
		},
		Permission: "ListVocabularyFilters",
	},
}
