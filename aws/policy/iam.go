package policy

var IAMPolicies = []Service{
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListRoles&Version=2010-05-08",
		Permission:    "ListRoles",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountAuthorizationDetails&Version=2010-05-08",
		Permission:    "GetAccountAuthorizationDetails",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountPasswordPolicy&Version=2010-05-08",
		Permission:    "GetAccountPasswordPolicy",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetAccountSummary&Version=2010-05-08",
		Permission:    "GetAccountSummary",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetCredentialReport&Version=2010-05-08",
		Permission:    "GetCredentialReport",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=GetUser&Version=2010-05-08",
		Permission:    "GetUser",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListAccessKeys&Version=2010-05-08",
		Permission:    "ListAccessKeys",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListAccountAliases&Version=2010-05-08",
		Permission:    "ListAccountAliases",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListGroups&Version=2010-05-08",
		Permission:    "ListGroups",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListInstanceProfiles&Version=2010-05-08",
		Permission:    "ListInstanceProfiles",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListMFADevices&Version=2010-05-08",
		Permission:    "ListMFADevices",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListOpenIDConnectProviders&Version=2010-05-08",
		Permission:    "ListOpenIDConnectProviders",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListPolicies&Version=2010-05-08",
		Permission:    "ListPolicies",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListRoles&Version=2010-05-08",
		Permission:    "ListRoles",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListSamlProviders&Version=2010-05-08",
		Permission:    "ListSamlProviders",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListServerCertificates&Version=2010-05-08",
		Permission:    "ListServerCertificates",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListServiceSpecificCredentials&Version=2010-05-08",
		Permission:    "ListServiceSpecificCredentials",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListSigningCertificates&Version=2010-05-08",
		Permission:    "ListSigningCertificates",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListSshPublicKeys&Version=2010-05-08",
		Permission:    "ListSshPublicKeys",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListUsers&Version=2010-05-08",
		Permission:    "ListUsers",
	},
	{
		IgnoreRegion:  true,
		ServiceSuffix: "?Action=ListVirtualMfaDevices&Version=2010-05-08",
		Permission:    "ListVirtualMfaDevices",
	},
}
