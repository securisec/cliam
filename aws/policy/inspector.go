package policy

import "github.com/securisec/cliam/shared"

var InspectorPolicies = []Service{
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeCrossAccountAccessRole",
		},
		Permission: "DescribeCrossAccountAccessRole",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentRuns",
		},
		Permission: "ListAssessmentRuns",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentTargets",
		},
		Permission: "ListAssessmentTargets",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentTemplates",
		},
		Permission: "ListAssessmentTemplates",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListEventSubscriptions",
		},
		Permission: "ListEventSubscriptions",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListFindings",
		},
		Permission: "ListFindings",
	},
	{
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListRulesPackages",
		},
		Permission: "ListRulesPackages",
	},

	// extra
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeAssessmentRuns",
		},
		Permission:             "DescribeAssessmentRuns",
		IsExtra:                true,
		ExtraComponentBodyKey:  "assessmentRunArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "assessment_run_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeAssessmentTargets",
		},
		Permission:             "DescribeAssessmentTargets",
		IsExtra:                true,
		ExtraComponentBodyKey:  "assessmentTargetArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "assessment_target_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeAssessmentTemplates",
		},
		Permission:             "DescribeAssessmentTemplates",
		IsExtra:                true,
		ExtraComponentBodyKey:  "assessmentTemplateArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "assessment_template_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeExclusions",
		},
		Permission:             "DescribeExclusions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "exclusionArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "exclusion_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeFindings",
		},
		Permission:             "DescribeFindings",
		IsExtra:                true,
		ExtraComponentBodyKey:  "findingArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "finding_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeResourceGroups",
		},
		Permission:             "DescribeResourceGroups",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceGroupArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_group_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeRulesPackages",
		},
		Permission:             "DescribeRulesPackages",
		IsExtra:                true,
		ExtraComponentBodyKey:  "rulesPackageArns",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "rules_package_arns",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.GetTelemetryMetadata",
		},
		Permission:             "GetTelemetryMetadata",
		IsExtra:                true,
		ExtraComponentBodyKey:  "assessmentRunArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "assessment_run_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentRunAgents",
		},
		Permission:             "ListAssessmentRunAgents",
		IsExtra:                true,
		ExtraComponentBodyKey:  "assessmentRunArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "assessment_run_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListExclusions",
		},
		Permission:             "ListExclusions",
		IsExtra:                true,
		ExtraComponentBodyKey:  "assessmentRunArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "assessment_run_arn",
	},
	{
		Method: "POST",
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListTagsForResource",
		},
		Permission:             "ListTagsForResource",
		IsExtra:                true,
		ExtraComponentBodyKey:  "resourceArn",
		ExtraComponentLocation: "json",
		ExtraCommandLineFlag:   "resource_arn",
	},
}
