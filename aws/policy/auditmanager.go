package policy

import "github.com/securisec/cliam/shared"

var AuditManagerPolicies = []Service{
	{
		Method:        "POST",
		ServiceSuffix: "account/deregisterAccount",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DeregisterAccount",
	},
	{
		Method:        "POST",
		ServiceSuffix: "account/deregisterOrganizationAdminAccount",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DeregisterOrganizationAdminAccount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "account/status",
		Permission:    "GetAccountStatus",
	},
	{
		Method:        "GET",
		ServiceSuffix: "delegations",
		Permission:    "GetDelegations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "insights",
		Permission:    "GetInsights",
	},
	{
		Method:        "GET",
		ServiceSuffix: "account/organizationAdminAccount",
		Permission:    "GetOrganizationAdminAccount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "services",
		Permission:    "GetServicesInScope",
	},
	{
		Method:        "GET",
		ServiceSuffix: "assessmentReports",
		Permission:    "ListAssessmentReports",
	},
	{
		Method:        "GET",
		ServiceSuffix: "assessments",
		Permission:    "ListAssessments",
	},
	{
		Method:        "GET",
		ServiceSuffix: "insights/control-domains",
		Permission:    "ListControlDomainInsights",
	},
	{
		Method:        "GET",
		ServiceSuffix: "notifications",
		Permission:    "ListNotifications",
	},
	{
		Method:        "POST",
		ServiceSuffix: "account/registerAccount",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "RegisterAccount",
	},

	// extra
	{
		ServiceSuffix:          "/assessments/{{.assessment_id}}",
		Permission:             "GetAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	{
		ServiceSuffix:          "/assessmentFrameworks/{{.framework_id}}",
		Permission:             "GetAssessmentFramework",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "framework_id",
	},
	{
		ServiceSuffix:          "/assessments/{{.assessment_id}}/changelogs",
		Permission:             "GetChangeLogs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	{
		ServiceSuffix:          "/controls/{{.control_id}}",
		Permission:             "GetControl",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "control_id",
	},
	{
		ServiceSuffix:          "/assessments/{{.assessment_id}}/evidenceFolders",
		Permission:             "GetEvidenceFoldersByAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	{
		ServiceSuffix:          "/insights/assessments/{{.assessment_id}}",
		Permission:             "GetInsightsByAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	{
		ServiceSuffix:          "/settings/{{.attribute}}",
		Permission:             "GetSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "attribute",
	},
	{
		ServiceSuffix:          "/assessmentFrameworkShareRequests",
		Permission:             "ListAssessmentFrameworkShareRequests",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "request_type",
	},
	{
		ServiceSuffix:          "/assessmentFrameworks",
		Permission:             "ListAssessmentFrameworks",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "framework_type",
	},
	{
		ServiceSuffix:          "/insights/control-domains-by-assessment",
		Permission:             "ListControlDomainInsightsByAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	{
		ServiceSuffix:          "/insights/controls",
		Permission:             "ListControlInsightsByControlDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "control_domain_id",
	},
	{
		ServiceSuffix:          "/controls",
		Permission:             "ListControls",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "control_type",
	},
	{
		ServiceSuffix:          "/dataSourceKeywords",
		Permission:             "ListKeywordsForDataSource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "source",
	},
	{
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
