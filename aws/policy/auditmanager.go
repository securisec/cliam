package policy

import "github.com/securisec/cliam/shared"

// AuditManagerPolicies policy
var AuditManagerPolicies = map[string]Service{
	"DeregisterAccount": {
		Method:        "POST",
		ServiceSuffix: "account/deregisterAccount",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DeregisterAccount",
	},
	"DeregisterOrganizationAdminAccount": {
		Method:        "POST",
		ServiceSuffix: "account/deregisterOrganizationAdminAccount",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "DeregisterOrganizationAdminAccount",
	},
	"GetAccountStatus": {
		Method:        "GET",
		ServiceSuffix: "account/status",
		Permission:    "GetAccountStatus",
	},
	"GetDelegations": {
		Method:        "GET",
		ServiceSuffix: "delegations",
		Permission:    "GetDelegations",
	},
	"GetInsights": {
		Method:        "GET",
		ServiceSuffix: "insights",
		Permission:    "GetInsights",
	},
	"GetOrganizationAdminAccount": {
		Method:        "GET",
		ServiceSuffix: "account/organizationAdminAccount",
		Permission:    "GetOrganizationAdminAccount",
	},
	"GetServicesInScope": {
		Method:        "GET",
		ServiceSuffix: "services",
		Permission:    "GetServicesInScope",
	},
	"ListAssessmentReports": {
		Method:        "GET",
		ServiceSuffix: "assessmentReports",
		Permission:    "ListAssessmentReports",
	},
	"ListAssessments": {
		Method:        "GET",
		ServiceSuffix: "assessments",
		Permission:    "ListAssessments",
	},
	"ListControlDomainInsights": {
		Method:        "GET",
		ServiceSuffix: "insights/control-domains",
		Permission:    "ListControlDomainInsights",
	},
	"ListNotifications": {
		Method:        "GET",
		ServiceSuffix: "notifications",
		Permission:    "ListNotifications",
	},
	"RegisterAccount": {
		Method:        "POST",
		ServiceSuffix: "account/registerAccount",
		JsonData:      map[string]string{},
		Headers: map[string]string{
			shared.CONTENT_TYPE_HEADER: aws_JSON_1_1,
		},
		Permission: "RegisterAccount",
	},

	// extra
	"GetAssessment": {
		ServiceSuffix:          "/assessments/{{.assessment_id}}",
		Permission:             "GetAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	"GetAssessmentFramework": {
		ServiceSuffix:          "/assessmentFrameworks/{{.framework_id}}",
		Permission:             "GetAssessmentFramework",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "framework_id",
	},
	"GetChangeLogs": {
		ServiceSuffix:          "/assessments/{{.assessment_id}}/changelogs",
		Permission:             "GetChangeLogs",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	"GetControl": {
		ServiceSuffix:          "/controls/{{.control_id}}",
		Permission:             "GetControl",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "control_id",
	},
	"GetEvidenceFoldersByAssessment": {
		ServiceSuffix:          "/assessments/{{.assessment_id}}/evidenceFolders",
		Permission:             "GetEvidenceFoldersByAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	"GetInsightsByAssessment": {
		ServiceSuffix:          "/insights/assessments/{{.assessment_id}}",
		Permission:             "GetInsightsByAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	"GetSettings": {
		ServiceSuffix:          "/settings/{{.attribute}}",
		Permission:             "GetSettings",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "attribute",
	},
	"ListAssessmentFrameworkShareRequests": {
		ServiceSuffix:          "/assessmentFrameworkShareRequests",
		Permission:             "ListAssessmentFrameworkShareRequests",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "request_type",
	},
	"ListAssessmentFrameworks": {
		ServiceSuffix:          "/assessmentFrameworks",
		Permission:             "ListAssessmentFrameworks",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "framework_type",
	},
	"ListControlDomainInsightsByAssessment": {
		ServiceSuffix:          "/insights/control-domains-by-assessment",
		Permission:             "ListControlDomainInsightsByAssessment",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "assessment_id",
	},
	"ListControlInsightsByControlDomain": {
		ServiceSuffix:          "/insights/controls",
		Permission:             "ListControlInsightsByControlDomain",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "control_domain_id",
	},
	"ListControls": {
		ServiceSuffix:          "/controls",
		Permission:             "ListControls",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "control_type",
	},
	"ListKeywordsForDataSource": {
		ServiceSuffix:          "/dataSourceKeywords",
		Permission:             "ListKeywordsForDataSource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "source",
	},
	"ListTagsForResource": {
		ServiceSuffix:          "/tags/{{.resource_arn}}",
		Permission:             "ListTagsForResource",
		ExtraComponentLocation: "path",
		IsExtra:                true,
		ExtraCommandLineFlag:   "resource_arn",
	},
}
