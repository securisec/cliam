package policy

var GuardDutyPolicies = []Service{
	{
		Method:        "GET",
		ServiceSuffix: "invitation/count",
		Permission:    "GetInvitationsCount",
	},
	{
		Method:        "GET",
		ServiceSuffix: "detector",
		Permission:    "ListDetectors",
	},
	{
		Method:        "GET",
		ServiceSuffix: "invitation",
		Permission:    "ListInvitations",
	},
	{
		Method:        "GET",
		ServiceSuffix: "admin",
		Permission:    "ListOrganizationAdminAccounts",
	},
}
