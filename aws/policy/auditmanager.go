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
}
