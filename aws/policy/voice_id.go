package policy

import "github.com/securisec/cliam/shared"

var VoiceIdPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "VoiceID.ListDomains",
		},
		Permission: "ListDomains",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "VoiceID.DescribeDomain",
		},
		Permission:             "DescribeDomain",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "VoiceID.ListFraudsterRegistrationJobs",
		},
		Permission:             "ListFraudsterRegistrationJobs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "VoiceID.ListSpeakerEnrollmentJobs",
		},
		Permission:             "ListSpeakerEnrollmentJobs",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "VoiceID.ListSpeakers",
		},
		Permission:             "ListSpeakers",
		IsExtra:                true,
		ExtraComponentBodyKey:  "DomainId",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "domain_id",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_0,
			aws_X_AMZ_TARGET:           "VoiceID.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "ResourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
