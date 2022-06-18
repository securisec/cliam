package policy

import "github.com/securisec/cliam/shared"

var InspectorPolicies = map[string]Service{
	"DescribeCrossAccountAccessRole": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.DescribeCrossAccountAccessRole",
		},
		Permission: "DescribeCrossAccountAccessRole",
	},
	"ListAssessmentRuns": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentRuns",
		},
		Permission: "ListAssessmentRuns",
	},
	"ListAssessmentTargets": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentTargets",
		},
		Permission: "ListAssessmentTargets",
	},
	"ListAssessmentTemplates": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListAssessmentTemplates",
		},
		Permission: "ListAssessmentTemplates",
	},
	"ListEventSubscriptions": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListEventSubscriptions",
		},
		Permission: "ListEventSubscriptions",
	},
	"ListFindings": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListFindings",
		},
		Permission: "ListFindings",
	},
	"ListRulesPackages": {
		Method:   "POST",
		JsonData: map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
			aws_X_AMZ_TARGET:           "InspectorService.ListRulesPackages",
		},
		Permission: "ListRulesPackages",
	},

	// extra
	"DescribeAssessmentRuns": {
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
	"DescribeAssessmentTargets": {
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
	"DescribeAssessmentTemplates": {
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
	"DescribeExclusions": {
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
	"DescribeFindings": {
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
	"DescribeResourceGroups": {
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
	"DescribeRulesPackages": {
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
	"GetTelemetryMetadata": {
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
	"ListAssessmentRunAgents": {
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
	"ListExclusions": {
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
	"ListTagsForResource": {
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
